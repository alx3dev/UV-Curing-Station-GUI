package uvs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func mainScreen(uv *UV_Station) fyne.CanvasObject {
	/*
		fyne app & window with permanent
		storage and default configuration
	*/
	a := uv.APP
	w := uv.WIN
	T := uv.T

	config := a.Preferences()

	// load defaults
	timer := config.IntWithFallback("TIMER", TIMER)
	power := config.IntWithFallback("POWER", POWER)
	speed := config.IntWithFallback("SPEED", SPEED)

	/*
		define binding and slide for each parameter
	*/
	timerBind := binding.NewFloat()
	timerBind.Set(float64(timer))

	powerBind := binding.NewFloat()
	powerBind.Set(float64(power))

	speedBind := binding.NewFloat()
	speedBind.Set(float64(speed))

	timerSlide := widget.NewSliderWithData(0, float64(TIMER_MAX), timerBind)
	powerSlide := widget.NewSliderWithData(0, float64(POWER_MAX), powerBind)
	speedSlide := widget.NewSliderWithData(0, float64(SPEED_MAX), speedBind)

	//format timer output
	timerText := container.NewGridWithColumns(2,
		widget.NewLabel(T.Timer),
		widget.NewLabelWithData(binding.FloatToStringWithFormat(
			timerBind, "%0.0f m")))

	// format led power output
	powerText := container.NewGridWithColumns(2,
		widget.NewLabel(T.Power),
		widget.NewLabelWithData(binding.FloatToStringWithFormat(
			powerBind, "%0.0f %%")))

	// format motor speed output
	speedText := container.NewGridWithColumns(2,
		widget.NewLabel(T.Speed),
		widget.NewLabelWithData(binding.FloatToStringWithFormat(
			speedBind, "%0.0f %%")))

	// timer buttons (- +)
	buttons := container.NewGridWithColumns(2,
		widget.NewButtonWithIcon("",
			theme.ContentRemoveIcon(),
			func() {
				if timerSlide.Value > timerSlide.Min {
					timerBind.Set(timerSlide.Value - 1)
				}
			}),

		widget.NewButtonWithIcon("",
			theme.ContentAddIcon(),
			func() {
				if timerSlide.Value < timerSlide.Max {
					timerBind.Set(timerSlide.Value + 1)
				}
			}))

	// led power buttons (- +)
	pbuttons := container.NewGridWithColumns(2,
		widget.NewButtonWithIcon("",
			theme.ContentRemoveIcon(),
			func() {
				if powerSlide.Value > powerSlide.Min {
					powerBind.Set(powerSlide.Value - 1)
				}
			}),

		widget.NewButtonWithIcon("",
			theme.ContentAddIcon(),
			func() {
				if powerSlide.Value < powerSlide.Max {
					powerBind.Set(powerSlide.Value + 1)
				}
			}))

	// motor speed buttons (- +)
	sbuttons := container.NewGridWithColumns(2,
		widget.NewButtonWithIcon("",
			theme.ContentRemoveIcon(),
			func() {
				if speedSlide.Value > speedSlide.Min {
					speedBind.Set(speedSlide.Value - 2)
				}
			}),

		widget.NewButtonWithIcon("",
			theme.ContentAddIcon(),
			func() {
				if speedSlide.Value < speedSlide.Max {
					speedBind.Set(speedSlide.Value + 2)
				}
			}))

	timerOpts := container.NewAdaptiveGrid(2, container.New(layout.NewFormLayout(), timerText, timerSlide), buttons)
	powerOpts := container.NewAdaptiveGrid(2, container.New(layout.NewFormLayout(), powerText, powerSlide), pbuttons)
	speedOpts := container.NewAdaptiveGrid(2, container.New(layout.NewFormLayout(), speedText, speedSlide), sbuttons)

	timerSlide.OnChanged = func(f float64) {
		timerBind.Set(f)
		a.Preferences().SetInt("TIMER", int(f))
	}

	powerSlide.OnChanged = func(f float64) {
		powerBind.Set(f)
		a.Preferences().SetInt("POWER", int(f))
	}

	speedSlide.OnChanged = func(f float64) {
		speedBind.Set(f)
		a.Preferences().SetInt("SPEED", int(f))
	}

	uv.timerBind = timerBind
	uv.powerBind = powerBind
	uv.speedBind = speedBind

	updateButton := widget.NewButtonWithIcon("", theme.ConfirmIcon(), func() {
		//to-do Send command to ESP-32
	})
	updateButton.Importance = widget.HighImportance

	defaultsButton := widget.NewButtonWithIcon("", theme.ViewRefreshIcon(), func() {
		timerBind.Set(float64(TIMER))
		powerBind.Set(float64(POWER))
		speedBind.Set(float64(SPEED))
	})

	quitButton := widget.NewButtonWithIcon("", theme.CancelIcon(), func() {
		a.Quit()
	})

	var controlButtons *fyne.Container

	if fyne.CurrentDevice().IsMobile() || fyne.CurrentDevice().IsBrowser() {
		controlButtons = container.New(layout.NewGridLayout(2), updateButton, defaultsButton)
	} else {
		controlButtons = container.New(layout.NewGridLayout(3), updateButton, defaultsButton, quitButton)
	}

	uvs_opts := container.NewVBox(timerOpts, powerOpts, speedOpts)
	bottom_buttons := container.NewVBox(controlButtons)

	content := container.NewBorder(nil, bottom_buttons, nil, nil, uvs_opts)

	w.Canvas().SetOnTypedKey(func(key *fyne.KeyEvent) {
		switch key.Name {

		case fyne.KeyRight:
			if timerSlide.Value < timerSlide.Max {
				timerBind.Set(timerSlide.Value + 1)
			}
		case fyne.KeyLeft:
			if timerSlide.Value > timerSlide.Min {
				timerBind.Set(timerSlide.Value - 1)
			}
		case fyne.KeyUp:
			if powerSlide.Value < powerSlide.Max {
				powerBind.Set(powerSlide.Value + 1)
			}
		case fyne.KeyDown:
			if powerSlide.Value > powerSlide.Min {
				powerBind.Set(powerSlide.Value - 1)
			}
		case fyne.KeySpace:
			timerBind.Set(float64(TIMER))
			powerBind.Set(float64(POWER))
			speedBind.Set(float64(SPEED))

		case fyne.KeyReturn:

		case fyne.KeyEscape:
		}
	})
	content.Refresh()

	return content
}
