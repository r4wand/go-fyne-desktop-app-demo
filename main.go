package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Main window")

	// this window will close all other windows when it is closed
	w.SetMaster()

	// resize the window to a specific size
	w.Resize(fyne.NewSize(800, 600))

	hello := widget.NewLabel("Hello Fyne")

	w.SetContent(container.NewVBox(
		widget.NewLabel("Hello Fyne!"),
		hello,
		widget.NewButton("New window", func() {
			w2 := a.NewWindow("New window")
			w2.SetContent(widget.NewLabel("Hello Fyne!"))
			w2.Show()
		}),
		widget.NewButton("Quit", func() {
			a.Quit()
		}),
	))
	go func() {
		for range time.Tick(time.Second) {
			formatted := time.Now().Format("Time: 03:04:05:00")
			hello.SetText(formatted)

		}
	}()

	w.ShowAndRun()
}
