package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"

	"github.com/mattn/go-gtk/gtk"
	"golang.org/x/text/unicode/norm"
)

func getTextViewContent(textView *gtk.TextView) string {
	buffer := textView.GetBuffer()

	var startIter, endIter gtk.TextIter
	buffer.GetStartIter(&startIter)
	buffer.GetEndIter(&endIter)
	text := buffer.GetText(&startIter, &endIter, true)
	return text
}

func updateTextViewContent(textView *gtk.TextView, text string) {
	buffer := textView.GetBuffer()
	var startIter, endIter gtk.TextIter
	buffer.GetStartIter(&startIter)
	buffer.GetEndIter(&endIter)
	buffer.Delete(&startIter, &endIter)
	buffer.Insert(&startIter, text)
}

func openTlsBtn(text string) *gtk.Button {
	openGgtlstBtn := gtk.NewButtonWithLabel("Open Google Translate")
	openGgtlstBtn.SetSizeRequest(20, 10)

	openGgtlstBtn.Connect("clicked", func() {
		url := fmt.Sprintf("https://translate.google.com/?sl=en&tl=vi&text=%s&op=translate", text)
		cmd := exec.Command("google-chrome", url)
		err := cmd.Start()
		if err != nil {
			fmt.Println("Failed to execute commandeclaredd:", err)
		}
	})
	return openGgtlstBtn
}

func getSpeakBtn(text string) *gtk.Button {
	btn := gtk.NewButtonWithLabel("Speak")
	btn.SetSizeRequest(20, 10)

	btn.Connect("clicked", func() {
		cmd := exec.Command("google_speech", text)
		err := cmd.Start()

		if err != nil {
			fmt.Println("Failed to execute command:", err)
		}
	})
	return btn
}

func getWindow(orgText, transText string) *gtk.Window {

	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetPosition(gtk.WIN_POS_CENTER)
	window.SetTitle("Translate GUI")
	window.SetIconName("textview")

	window.Connect("destroy", gtk.MainQuit)
	window.Connect("focus-out-event", gtk.MainQuit)

	orgTextView := gtk.NewTextView()
	orgTextView.SetWrapMode(gtk.WRAP_WORD)
	orgTextView.SetEditable(true)
	orgTextView.SetCursorVisible(true)
	orgTextView.SetBorderWidth(2)

	orgBuffer := orgTextView.GetBuffer()
	var orgTextIter gtk.TextIter
	orgBuffer.GetStartIter(&orgTextIter)
	orgBuffer.Insert(&orgTextIter, orgText)

	translatedTextView := gtk.NewTextView()
	translatedTextView.SetWrapMode(gtk.WRAP_WORD)
	translatedTextView.SetCursorVisible(false)

	var (
		inputTextcancel context.CancelFunc
		inputTextCtx    context.Context
	)
	orgBuffer.Connect("changed", func() {
		text := getTextViewContent(orgTextView)
		if inputTextcancel != nil {
			inputTextcancel()
		}
		inputTextCtx, inputTextcancel = context.WithCancel(context.Background())
		go func() {
			for {
				select {
				case <-inputTextCtx.Done():
					return
				case <-time.After(1 * time.Second):
					transText := trans(text)
					// reset cancel function when user typed new input
					updateTextViewContent(translatedTextView, transText)
				}
			}
		}()
	})

	var iter gtk.TextIter
	buffer := translatedTextView.GetBuffer()

	buffer.GetStartIter(&iter)
	buffer.Insert(&iter, transText)
	topVbox := gtk.NewVBox(true, 0)
	topVbox.PackStart(orgTextView, true, true, 0)
	topVbox.SetSpacing(5)
	topVbox.PackStart(translatedTextView, true, true, 0)
	topVbox.SetSizeRequest(400, 200)

	// Add the button to a vertical box
	vbox := gtk.NewVBox(false, 1)
	vbox.PackStart(topVbox, true, true, 0)

	// button group
	bottomHbox := gtk.NewHBox(true, 1)
	bottomHbox.SetSizeRequest(400, 50)

	openGgtlstBtn := openTlsBtn(orgText)
	speakBtn := getSpeakBtn(orgText)

	bottomHbox.PackStart(openGgtlstBtn, true, true, 0)
	bottomHbox.PackStart(speakBtn, true, true, 0)

	vbox.PackStart(bottomHbox, false, false, 0)
	// vbox.PackEnd(openGgtlstBtn, true, true, 0)
	window.Add(vbox)

	window.SetSizeRequest(400, 300)
	window.ShowAll()

	return window
}

func trans(text string) string {
	cmd := exec.Command("trans", ":vi", "-b", text)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Failed to execute command:", err)
		return ""
	} else {
		res := norm.NFC.String(string(out))
		return string(res)
	}
}

func getSelectedText() string {
	cmd := exec.Command("xclip", "-out", "-selection", "primary")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Failed to execute command:", err)
		panic(err)
	} else {
		return string(out)
	}
}

func showWindow() {
	selectedText := getSelectedText()
	transText := trans(selectedText)

	gtk.Init(nil)
	_ = getWindow(selectedText, transText)
	gtk.Main()
}

func main() {
	showWindow()
}
