package uvs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Main screen contain all controls for UV station
//
// Many things will be changed here during development
func mainScreen(uv *UV_Station) fyne.CanvasObject {

	// define our app, master window and translation
	a := uv.APP
	w := uv.WIN
	T := uv.T

	// load values from storage
	timer := uv.config.IntWithFallback("TIMER", TIMER)
	power := uv.config.IntWithFallback("POWER", POWER)
	speed := uv.config.IntWithFallback("SPEED", SPEED)

	// define binding and slide for each control
	timerBind := binding.NewFloat()
	timerBind.Set(float64(timer))

	powerBind := binding.NewFloat()
	powerBind.Set(float64(power))

	speedBind := binding.NewFloat()
	speedBind.Set(float64(speed))

	timerSlide := widget.NewSliderWithData(0, float64(TIMER_MAX), timerBind)
	powerSlide := widget.NewSliderWithData(0, float64(POWER_MAX), powerBind)
	speedSlide := widget.NewSliderWithData(0, float64(SPEED_MAX), speedBind)

	/*
		Format Values output
	*/

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
			func() { uv.decreaseValue(timerSlide, timerBind) }),

		widget.NewButtonWithIcon("",
			theme.ContentAddIcon(),
			func() { uv.increaseValue(timerSlide, timerBind) }))

	// led power buttons (- +)
	pbuttons := container.NewGridWithColumns(2,
		widget.NewButtonWithIcon("",
			theme.ContentRemoveIcon(),
			func() { uv.decreaseValue(powerSlide, powerBind) }),

		widget.NewButtonWithIcon("",
			theme.ContentAddIcon(),
			func() { uv.increaseValue(powerSlide, powerBind) }))

	// motor speed buttons (- +)
	sbuttons := container.NewGridWithColumns(2,
		widget.NewButtonWithIcon("",
			theme.ContentRemoveIcon(),
			func() { uv.decreaseValue(speedSlide, speedBind) }),

		widget.NewButtonWithIcon("",
			theme.ContentAddIcon(),
			func() { uv.increaseValue(speedSlide, speedBind) }))

	//	Format UV Station control output
	timerOpts := container.NewAdaptiveGrid(2, container.New(layout.NewFormLayout(), timerText, timerSlide), buttons)
	powerOpts := container.NewAdaptiveGrid(2, container.New(layout.NewFormLayout(), powerText, powerSlide), pbuttons)
	speedOpts := container.NewAdaptiveGrid(2, container.New(layout.NewFormLayout(), speedText, speedSlide), sbuttons)

	//	Control and sync slides/values
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

	/*
		Define buttons for main screen
	*/

	// to-do Send START command to ESP-32
	updateButton := widget.NewButtonWithIcon("", theme.ConfirmIcon(), func() {})
	updateButton.Importance = widget.HighImportance

	// load default values
	defaultsButton := widget.NewButtonWithIcon("", theme.ViewRefreshIcon(), func() {
		uv.loadDefaults()
	})

	// to-do Send STOP command to ESP-32
	quitButton := widget.NewButtonWithIcon("", theme.CancelIcon(), func() {
		a.Quit() // this blocks on Android
	})

	// only for development, app.quit blocks on android
	var controlButtons *fyne.Container
	if uv.APP.Driver().Device().IsMobile() || uv.APP.Driver().Device().IsBrowser() {
		controlButtons = container.New(layout.NewGridLayout(2), updateButton, defaultsButton)
	} else {
		controlButtons = container.New(layout.NewGridLayout(3), updateButton, defaultsButton, quitButton)
	}

	// put everything together (control and buttons)
	uvs_opts := container.NewVBox(timerOpts, powerOpts, speedOpts)
	bottom_buttons := container.NewVBox(controlButtons)

	screen := container.NewBorder(nil, bottom_buttons, nil, nil, uvs_opts)

	// add some shortcuts for easier work on PC
	w.Canvas().SetOnTypedKey(func(key *fyne.KeyEvent) {
		switch key.Name {

		case fyne.KeyRight:
			uv.increaseValue(timerSlide, timerBind)

		case fyne.KeyLeft:
			uv.decreaseValue(timerSlide, timerBind)

		case fyne.KeyUp:
			uv.increaseValue(powerSlide, powerBind)

		case fyne.KeyDown:
			uv.decreaseValue(powerSlide, powerBind)

		case fyne.KeyPageDown:
			uv.decreaseValue(speedSlide, speedBind)

		case fyne.KeyPageUp:
			uv.increaseValue(speedSlide, speedBind)

		case fyne.KeySpace:
			uv.loadDefaults()

		case fyne.KeyReturn: // to-do Send command to ESP32
			if w.FullScreen() {
				w.SetFullScreen(false)
			} else {
				w.SetFullScreen(true)
			}

		case fyne.KeyEscape:
			// to-do Minimize to taskbar?
		}
	})
	screen.Refresh()

	return screen
}

/*
Helper methods for main screen
*/

// allow user to change default values
func (uv *UV_Station) loadDefaults() {
	uv.timerBind.Set(float64(uv.config.IntWithFallback("TIMER_DEFAULT", TIMER)))
	uv.powerBind.Set(float64(uv.config.IntWithFallback("POWER_DEFAULT", POWER)))
	uv.speedBind.Set(float64(uv.config.IntWithFallback("SPEED_DEFAULT", SPEED)))
}

// increase slider bind value by 1
func (uv *UV_Station) increaseValue(slide *widget.Slider, bind binding.Float) {
	if slide.Value < slide.Max {
		bind.Set(slide.Value + 1)
	}
}

// decrease slider bind value by 1
func (uv *UV_Station) decreaseValue(slide *widget.Slider, bind binding.Float) {
	if slide.Value > slide.Min {
		bind.Set(slide.Value - 1)
	}
}
