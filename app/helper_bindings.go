package uvs

import (
	"runtime"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

var TIMER int = 5
var POWER int = 80
var SPEED int = 30

var TIMER_MAX int = 30
var POWER_MAX int = 100
var SPEED_MAX int = 100

// Keep items for automatic refresh
// after language change
type Subitems struct {
	mainTab     *container.TabItem
	settingsTab *container.TabItem

	configButton *widget.Button
	configLabel  binding.String

	timerLabel          binding.String
	powerLabel          binding.String
	speedLabel          binding.String
	chooseThemeLabel    binding.String
	chooseLanguageLabel binding.String

	themeSelect *widget.RadioGroup
	langSelect  *widget.RadioGroup
}

func (uv *UV_Station) InitializeBindings() {
	// load values from storage
	timer := uv.config.IntWithFallback("TIMER", TIMER)
	power := uv.config.IntWithFallback("POWER", POWER)
	speed := uv.config.IntWithFallback("SPEED", SPEED)

	uv.IP = uv.config.StringWithFallback("IP", "127.0.0.1")
	uv.PORT = uv.config.StringWithFallback("PORT", "80")

	uv.dial.SetUri(uv.IP + ":" + uv.PORT)

	// define bindings for control values
	uv.timerBind = binding.NewFloat()
	uv.timerBind.Set(float64(timer))

	uv.powerBind = binding.NewFloat()
	uv.powerBind.Set(float64(power))

	uv.speedBind = binding.NewFloat()
	uv.speedBind.Set(float64(speed))

	// define control label bindings for automatic translate
	uv.sub.timerLabel = binding.NewString()
	uv.sub.timerLabel.Set(uv.T.Timer)

	uv.sub.powerLabel = binding.NewString()
	uv.sub.powerLabel.Set(uv.T.Power)

	uv.sub.speedLabel = binding.NewString()
	uv.sub.speedLabel.Set(uv.T.Speed)

	// define theme and language label bindings for automatic translate
	uv.sub.chooseThemeLabel = binding.NewString()
	uv.sub.chooseThemeLabel.Set(uv.T.ChooseTheme)

	uv.sub.chooseLanguageLabel = binding.NewString()
	uv.sub.chooseLanguageLabel.Set(uv.T.ChooseLanguage)

	// define IP & Port config label binding for automatic translate
	uv.sub.configLabel = binding.NewString()
	uv.sub.configLabel.Set(uv.T.Configuration)
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

func getOS() uint8 {
	switch runtime.GOOS {
	case "linux", "freebsd", "netbsd", "openbsd", "dragonfly":
		return 1
	case "android":
		return 2
	case "windows":
		return 3
	case "darwin":
		return 4
	case "ios":
		return 5
	}
	return 0
}

func (uv *UV_Station) isMobile() bool {
	return uv.APP.Driver().Device().IsMobile()
}
