package main

import (
	"fmt"
	"os/exec"

	"github.com/mattn/go-gtk/gtk"
	"golang.org/x/text/unicode/norm"
)

func getWindow(text string) *gtk.Window {

	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetPosition(gtk.WIN_POS_CENTER)
	window.SetTitle("GTK Go!")
	window.SetIconName("textview")
	window.Connect("destroy", gtk.MainQuit)

	textview := gtk.NewTextView()
	textview.SetEditable(false)
	textview.SetCursorVisible(false)
	var iter gtk.TextIter
	buffer := textview.GetBuffer()

	buffer.GetStartIter(&iter)
	buffer.Insert(&iter, text)

	window.Add(textview)
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
		fmt.Println(res)
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
		_ = getWindow(transText)
		gtk.Main()
	}

}

func main() {
	showWindow()
}