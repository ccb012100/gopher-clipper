package main

import (
	"log"

	"fyne.io/fyne/v2/app"
	"github.com/ccb012100/gopher-clip/gui"
	"golang.design/x/clipboard"
)

func main() {
	err := clipboard.Init()
	if err != nil {
		log.Println("Exiting because the system clipboard can't be accessed")
		panic(err)
	}

	windowTitle := "GopherClip"

	app := app.New()

	window := gui.BuildWindow(app, windowTitle)

	gui.InitializeStack(&window)

	window.ShowAndRun()
}
