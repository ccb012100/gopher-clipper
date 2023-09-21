package gui

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"golang.design/x/clipboard"
)

var content *fyne.Container

func BuildWindow(app fyne.App, windowTitle string) fyne.Window {
	window := app.NewWindow(windowTitle)

	// use window.SetIcon to set window & taskbar icon; app.SetIcon() doesn't work
	appIcon, err := fyne.LoadResourceFromPath("Icon.png")
	if err != nil {
		log.Println("error loading icon: " + err.Error())
	} else {
		appIcon = theme.ContentPasteIcon()
	}

	// Commented out because SystemTrayIcon is currently broken in fyne
	// buildSystemTray(app, window, appTitle, appIcon)

	window.SetIcon(appIcon)
	window.SetIcon(theme.ContentPasteIcon())

	return window
}

func InitializeStack(window *fyne.Window) {
	content = container.New(layout.NewVBoxLayout(), CreateCopyButton())

	(*window).SetContent(container.New(layout.NewCenterLayout(), content))
}

func CreateCopyButton() *widget.Button {
	return widget.NewButtonWithIcon("Add from Clipboard", theme.ContentPasteIcon(), func() {
		AddStackItem()
	})
}

func AddStackItem() {
	clip := GetClipboardText()
	btn := widget.NewButton(string(clip), func() {
		clip := clip
		CopyToClipboard(clip)
	})
	// TODO: add to top of list, not bottom
	content.Add(btn)
}

func CopyToClipboard(clip []byte) {
	clipboard.Write(clipboard.FmtText, clip)
}

func GetClipboardText() []byte {
	// TODO: is there a way to ignore images in clipboard?
	// TODO: ignore duplicates
	return clipboard.Read(clipboard.FmtText)
}

// ignore temporarily until https://github.com/fyne-io/fyne/issues/3968 is fixed
//
//lint:ignore U1000 Ignore unused function
func buildSystemTray(app fyne.App, window fyne.Window, appTitle string, appIcon fyne.Resource) {
	if desktopApp, ok := app.(desktop.App); ok {
		sysTrayMenu := fyne.NewMenu(appTitle,
			fyne.NewMenuItem("Show", func() {
				log.Println("Tapped show")
			}))

		desktopApp.SetSystemTrayIcon(appIcon)
		desktopApp.SetSystemTrayMenu(sysTrayMenu)
	} else {
		panic("This must run as a Desktop app!")
	}

	// Intercept the close event and minimize to system tray
	window.SetCloseIntercept(func() {
		window.Hide()
	})
}
