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

	// load defaults
	timer := uv.config.IntWithFallback("TIMER", TIMER)
	power := uv.config.IntWithFallback("POWER", POWER)
	speed := uv.config.IntWithFallback("SPEED", SPEED)

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
			func() { uv.decreaseValue(timerSlide, uv.timerBind) }),

		widget.NewButtonWithIcon("",
			theme.ContentAddIcon(),
			func() { uv.increaseValue(timerSlide, uv.timerBind) }))

	// led power buttons (- +)
	pbuttons := container.NewGridWithColumns(2,
		widget.NewButtonWithIcon("",
			theme.ContentRemoveIcon(),
			func() { uv.decreaseValue(powerSlide, uv.powerBind) }),

		widget.NewButtonWithIcon("",
			theme.ContentAddIcon(),
			func() { uv.increaseValue(powerSlide, uv.powerBind) }))

	// motor speed buttons (- +)
	sbuttons := container.NewGridWithColumns(2,
		widget.NewButtonWithIcon("",
			theme.ContentRemoveIcon(),
			func() { uv.decreaseValue(speedSlide, uv.speedBind) }),

		widget.NewButtonWithIcon("",
			theme.ContentAddIcon(),
			func() { uv.increaseValue(speedSlide, uv.speedBind) }))

	timerOpts := container.NewAdaptiveGrid(2, container.New(layout.NewFormLayout(), timerText, timerSlide), buttons)
	powerOpts := container.NewAdaptiveGrid(2, container.New(layout.NewFormLayout(), powerText, powerSlide), pbuttons)
	speedOpts := container.NewAdaptiveGrid(2, container.New(layout.NewFormLayout(), speedText, speedSlide), sbuttons)

	timerSlide.OnChanged = func(f float64) {
		timerBind.Set(f)
		uv.config.SetInt("TIMER", int(f))
	}

	powerSlide.OnChanged = func(f float64) {
		powerBind.Set(f)
		uv.config.SetInt("POWER", int(f))
	}

	speedSlide.OnChanged = func(f float64) {
		speedBind.Set(f)
		uv.config.SetInt("SPEED", int(f))
	}

	uv.timerBind = timerBind
	uv.powerBind = powerBind
	uv.speedBind = speedBind

	updateButton := widget.NewButtonWithIcon("", theme.ConfirmIcon(), func() {
		//to-do Send command to ESP-32
	})
	updateButton.Importance = widget.HighImportance

	defaultsButton := widget.NewButtonWithIcon("", theme.ViewRefreshIcon(), func() {
		uv.loadDefaults()
	})

	quitButton := widget.NewButtonWithIcon("", theme.CancelIcon(), func() {
		//to-do Send command to ESP-32
		a.Quit()
	})

	var controlButtons *fyne.Container

	if uv.APP.Driver().Device().IsMobile() || uv.APP.Driver().Device().IsBrowser() {
		controlButtons = container.New(layout.NewGridLayout(2), updateButton, defaultsButton)
	} else {
		controlButtons = container.New(layout.NewGridLayout(3), updateButton, defaultsButton, quitButton)
	}

	uvs_opts := container.NewVBox(timerOpts, powerOpts, speedOpts)
	bottom_buttons := container.NewVBox(controlButtons)

	screen := container.NewBorder(nil, bottom_buttons, nil, nil, uvs_opts)

	w.Canvas().SetOnTypedKey(func(key *fyne.KeyEvent) {
		switch key.Name {

		case fyne.KeyRight:
			uv.increaseValue(timerSlide, uv.timerBind)

		case fyne.KeyLeft:
			uv.decreaseValue(timerSlide, uv.timerBind)

		case fyne.KeyUp:
			uv.increaseValue(powerSlide, uv.powerBind)

		case fyne.KeyDown:
			uv.decreaseValue(powerSlide, uv.powerBind)

		case fyne.KeySpace:
			uv.loadDefaults()

		case fyne.KeyReturn:
			// to-do Send command to ESP32

		case fyne.KeyEscape:
			// to-do Minimize to taskbar?
		}
	})
	screen.Refresh()

	return screen
}

func (uv *UV_Station) loadDefaults() {
	uv.timerBind.Set(float64(uv.config.IntWithFallback("TIMER_DEFAULT", TIMER)))
	uv.powerBind.Set(float64(uv.config.IntWithFallback("POWER_DEFAULT", POWER)))
	uv.speedBind.Set(float64(uv.config.IntWithFallback("SPEED_DEFAULT", SPEED)))
}

func (uv *UV_Station) increaseValue(slide *widget.Slider, bind binding.Float) {
	if slide.Value < slide.Max {
		bind.Set(slide.Value + 1)
	}
}

func (uv *UV_Station) decreaseValue(slide *widget.Slider, bind binding.Float) {
	if slide.Value > slide.Min {
		bind.Set(slide.Value - 1)
	}
}
