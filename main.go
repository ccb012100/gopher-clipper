package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/ccb012100/gopher-clip/gui"
)

func main() {
	windowTitle := "GopherClip"

	app := app.New()

	window := gui.BuildWindow(app, windowTitle)

	gui.InitializeStack(&window)

	window.ShowAndRun()
}
