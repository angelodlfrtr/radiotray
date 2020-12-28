// Package settings
package settings

import (
	"github.com/webview/webview"
)

var view webview.WebView

func Init() {
	view = webview.New(true)
	// defer view.Destroy()
}

func Open() {
	// debug := true
	// w := webview.New(debug)
	// defer w.Destroy()
	// w.SetTitle("Minimal webview example")
	// w.SetSize(800, 600, webview.HintNone)
	// w.Navigate("https://en.m.wikipedia.org/wiki/Main_Page")
	// w.Run()
}
