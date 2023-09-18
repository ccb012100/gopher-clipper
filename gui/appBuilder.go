package gui

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func buildApp() {
	appTitle := "GopherClip"

	app := app.New()
	window := app.NewWindow(appTitle)

	// use window.SetIcon to set window & taskbar icon; app.SetIcon() doesn't work
	appIcon, err := fyne.LoadResourceFromPath("Icon.png")
	if err != nil {
		log.Println("error loading icon: " + err.Error())
	} else {
		appIcon = theme.ContentPasteIcon()
	}

	// Commented out because SystemTrayIcon is currently broken in fyne
	// buildSystemTray(app, window, appTitle, appIcon)

	hello := widget.NewLabel(appTitle)

	window.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Press me!", func() {
			hello.SetText("Clicked!")
		}),
	))

	window.SetIcon(appIcon)
	window.SetIcon(theme.ContentPasteIcon())

	window.ShowAndRun()
}

//lint:ignore U1000 Ignore unused function temporarily until https://github.com/fyne-io/fyne/issues/3968 is fixed
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
