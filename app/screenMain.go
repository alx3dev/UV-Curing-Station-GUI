package uvs

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Main screen contain all controls for UV station
// .
func mainScreen(uv *UV_Station) fyne.CanvasObject {

	// load default config and define bindings
	uv.InitializeBindings()

	timerSlide := widget.NewSliderWithData(0, float64(TIMER_MAX), uv.timerBind)
	powerSlide := widget.NewSliderWithData(0, float64(POWER_MAX), uv.powerBind)
	speedSlide := widget.NewSliderWithData(0, float64(SPEED_MAX), uv.speedBind)

	/*
		Format Values output
	*/

	//format timer output
	timerText := container.NewGridWithColumns(2,
		widget.NewLabelWithData(uv.sub.timerLabel),
		widget.NewLabelWithData(binding.FloatToStringWithFormat(
			uv.timerBind, "%0.0f m")))

	// format led power output
	powerText := container.NewGridWithColumns(2,
		widget.NewLabelWithData(uv.sub.powerLabel),
		widget.NewLabelWithData(binding.FloatToStringWithFormat(
			uv.powerBind, "%0.0f %%")))

	// format motor speed output
	speedText := container.NewGridWithColumns(2,
		widget.NewLabelWithData(uv.sub.speedLabel),
		widget.NewLabelWithData(binding.FloatToStringWithFormat(
			uv.speedBind, "%0.0f %%")))

	/*
		Add value buttons (- +)
	*/

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

	//	Format UV Station control output
	timerOpts := container.NewAdaptiveGrid(2, container.New(layout.NewFormLayout(), timerText, timerSlide), buttons)
	powerOpts := container.NewAdaptiveGrid(2, container.New(layout.NewFormLayout(), powerText, powerSlide), pbuttons)
	speedOpts := container.NewAdaptiveGrid(2, container.New(layout.NewFormLayout(), speedText, speedSlide), sbuttons)

	//	Control and sync slides/values
	timerSlide.OnChanged = func(f float64) {
		uv.timerBind.Set(f)
		uv.config.SetInt("TIMER", int(f))
	}

	powerSlide.OnChanged = func(f float64) {
		uv.powerBind.Set(f)
		uv.config.SetInt("POWER", int(f))
	}

	speedSlide.OnChanged = func(f float64) {
		uv.speedBind.Set(f)
		uv.config.SetInt("SPEED", int(f))
	}

	/*
		Define buttons for main screen
	*/

	// to-do Send START command to ESP-32
	startButton := widget.NewButtonWithIcon("", theme.ConfirmIcon(), func() {
		uv.dial.Params["timer"] = strconv.Itoa(int(timerSlide.Value))
		uv.dial.Params["power"] = strconv.Itoa(int(powerSlide.Value))
		uv.dial.Params["speed"] = strconv.Itoa(int(speedSlide.Value))

		uv.dial.GetRequest("uvs", uv.dial.Params)
	})
	startButton.Importance = widget.HighImportance

	stopButton := widget.NewButtonWithIcon("", theme.CancelIcon(), func() {
		uv.dial.Get("stop")
	})

	// load default values
	defaultsButton := widget.NewButtonWithIcon("", theme.ViewRefreshIcon(), func() {
		uv.loadDefaults()
	})

	controlButtons := container.NewGridWithColumns(3, startButton, defaultsButton, stopButton)

	// put everything together (control and buttons)
	uvs_opts := container.NewVBox(timerOpts, powerOpts, speedOpts)
	bottom_buttons := container.NewVBox(controlButtons)

	screen := container.NewBorder(nil, bottom_buttons, nil, nil, uvs_opts)

	// add some shortcuts for easier work on PC
	uv.WIN.Canvas().SetOnTypedKey(func(key *fyne.KeyEvent) {
		switch key.Name {

		case fyne.KeyRight:
			uv.increaseValue(timerSlide, uv.timerBind)

		case fyne.KeyLeft:
			uv.decreaseValue(timerSlide, uv.timerBind)

		case fyne.KeyUp:
			uv.increaseValue(powerSlide, uv.powerBind)

		case fyne.KeyDown:
			uv.decreaseValue(powerSlide, uv.powerBind)

		case fyne.KeyPageDown:
			uv.decreaseValue(speedSlide, uv.speedBind)

		case fyne.KeyPageUp:
			uv.increaseValue(speedSlide, uv.speedBind)

		case fyne.KeySpace:
			uv.loadDefaults()

		case fyne.KeyReturn: // to-do Send command to ESP32
			if uv.WIN.FullScreen() {
				uv.WIN.SetFullScreen(false)
			} else {
				uv.WIN.SetFullScreen(true)
			}

		case fyne.KeyEscape:
			uv.APP.Quit()
		}
	})
	screen.Refresh()

	return screen
}
