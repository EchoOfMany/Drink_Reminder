package main

import (
	"log"
	"math/rand"
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
	myWindow := myApp.NewWindow("Drink Reminder")
	myWindow.Resize(fyne.NewSize(200, 100))

	DEFAULT_HOURS := 0
	DEFAULT_MINUTES := 20
	DEFAULT_SECONDS := 0
	running := true

	hourData := binding.NewInt()
	hourData.Set(DEFAULT_HOURS)
	minuteData := binding.NewInt()
	minuteData.Set(DEFAULT_MINUTES)
	secondData := binding.NewInt()
	secondData.Set(DEFAULT_SECONDS)
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
	progress.TextFormatter = func() string { return " " }
	progress.Min = 0.0
	progress.Max = float64(((DEFAULT_HOURS * 3600) + (DEFAULT_MINUTES * 60) + DEFAULT_SECONDS) + 1)
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
		for i := 0; i <= int(progress.Max); i += 1 {
			time.Sleep(time.Second * 1)
			progress.SetValue(float64(i + 1))
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

			if s == 0 {
				if m == 0 {
					if h == 0 {
						err := beeep.Notify("Drink Reminder", RandomMessage(), "drink.png")
						if err != nil {
							panic(err)
						}
					} else {
						hourData.Set(h - 1)
						minuteData.Set(59)
						secondData.Set(59)
					}
				} else {
					minuteData.Set(m - 1)
					secondData.Set(59)
				}
			} else {
				secondData.Set(s - 1)
			}
			if i == int(progress.Max) && running {
				i = 0
				hourData.Set(DEFAULT_HOURS)
				minuteData.Set(DEFAULT_MINUTES)
				secondData.Set(DEFAULT_SECONDS)
				progress.SetValue(progress.Min)
			}
		}
	}()

	content := container.NewBorder(toolbar, nil, nil, nil, container.NewVBox(timeLbl, progress))

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

func RandomMessage() string {
	messages := []string{
		"Make sure to have a sippy",
		"Sippy Break",
		"Hydration Required",
		"Drink some water you beautiful and capable but dehydrated bitch",
		"Drink up Buttercup!",
		"Remember: It is possible to be a multifaceted woman. Read books + twerk. Be Spiritual + a freak",
		"Note to self: you good, you poppin'",
		"I don't know take a drink or something",
	}

	index := rand.Int() % len(messages)
	return messages[index]
}
