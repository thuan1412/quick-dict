package main

import (
	"fmt"
	"os/exec"

	"github.com/mattn/go-gtk/gtk"
	"golang.org/x/text/unicode/norm"
)

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

func getWindow(originText, transText string) *gtk.Window {

	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetPosition(gtk.WIN_POS_CENTER)
	window.SetTitle("Translate GUI")
	window.SetIconName("textview")

	window.Connect("destroy", gtk.MainQuit)
	window.Connect("focus-out-event", gtk.MainQuit)

	textview := gtk.NewTextView()
	textview.SetWrapMode(gtk.WRAP_WORD)
	textview.SetEditable(false)
	textview.SetCursorVisible(false)
	// textview.SetSizeRequest(20, 10)
	var iter gtk.TextIter
	buffer := textview.GetBuffer()

	buffer.GetStartIter(&iter)
	buffer.Insert(&iter, transText)

	topHbox := gtk.NewHBox(true, 0)
	topHbox.PackStart(textview, true, true, 0)
	topHbox.SetSizeRequest(400, 200)

	// Add the button to a vertical box
	vbox := gtk.NewVBox(false, 1)
	vbox.PackStart(topHbox, true, true, 0)

	// button group
	bottomHbox := gtk.NewHBox(true, 1)
	bottomHbox.SetSizeRequest(400, 50)

	openGgtlstBtn := openTlsBtn(originText)
	speakBtn := getSpeakBtn(originText)

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

func showWindow() {
	cmd := exec.Command("xclip", "-out", "-selection", "primary")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Failed to execute command:", err)
	} else {
		selectedText := string(out)
		transText := trans(selectedText)

		gtk.Init(nil)
		_ = getWindow(selectedText, transText)
		gtk.Main()
	}

}

func main() {
	showWindow()
}
