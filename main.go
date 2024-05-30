package main

import (
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/gen2brain/beeep"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("ProgressBar and Toolbar Widget")
	myWindow.Resize(fyne.NewSize(200, 100))

	hourData := binding.NewInt()
	hourData.Set(0)
	minuteData := binding.NewInt()
	minuteData.Set(0)
	secondData := binding.NewInt()
	secondData.Set(0)
	hourLbl := binding.IntToStringWithFormat(hourData, "%02d")
	minuteLbl := binding.IntToStringWithFormat(minuteData, "%02d")
	secondLbl := binding.IntToStringWithFormat(secondData, "%02d")
	timeLbl := container.New(layout.NewGridLayoutWithColumns(7),
		widget.NewLabel(" "),
		widget.NewLabelWithData(hourLbl),
		widget.NewLabel(" : "),
		widget.NewLabelWithData(minuteLbl),
		widget.NewLabel(" : "),
		widget.NewLabelWithData(secondLbl),
		widget.NewLabel(" "),
	)
	progress := widget.NewProgressBar()
	progress.Min = 0.0
	progress.Max = 1200.0
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			log.Println("Edit Time: needs implimented")
		}),
		widget.NewToolbarAction(theme.MediaPlayIcon(), func() {
			log.Println("Start Timer: needs implimented")
		}),
		widget.NewToolbarAction(theme.MediaStopIcon(), func() {
			log.Println("Stop Timer: needs implimented")
		}),
	)

	go func() {
		for i := 0.0; i <= progress.Max; i += 1.0 {
			time.Sleep(time.Second * 1)
			progress.SetValue(i)
			h, err := hourData.Get()
			if err != nil {
				panic(err)
			}
			m, err := minuteData.Get()
			if err != nil {
				panic(err)
			}
			s, err := secondData.Get()
			if err != nil {
				panic(err)
			}

			if s >= 59 {
				secondData.Set(0)
				err := beeep.Notify("Drink Reminder", "Make sure to have a sippy", "")
				if err != nil {
					panic(err)
				}
				if m >= 59 {
					minuteData.Set(0)
					hourData.Set(h + 1)
				} else {
					minuteData.Set(m + 1)
				}
			} else {
				secondData.Set(s + 1)
			}

		}
	}()

	content := container.NewBorder(toolbar, nil, nil, nil, container.NewVBox(timeLbl, progress))

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
