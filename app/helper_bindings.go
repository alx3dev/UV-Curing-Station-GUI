package uvs

import (
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func (uv *UV_Station) InitializeBindings() {
	// load values from storage
	timer := uv.config.IntWithFallback("TIMER", TIMER)
	power := uv.config.IntWithFallback("POWER", POWER)
	speed := uv.config.IntWithFallback("SPEED", SPEED)

	// define bindings for control values
	uv.timerBind = binding.NewFloat()
	uv.timerBind.Set(float64(timer))

	uv.powerBind = binding.NewFloat()
	uv.powerBind.Set(float64(power))

	uv.speedBind = binding.NewFloat()
	uv.speedBind.Set(float64(speed))

	// define label bindings for automatic translate
	uv.sub.timerLabel = binding.NewString()
	uv.sub.timerLabel.Set(uv.T.Timer)

	uv.sub.powerLabel = binding.NewString()
	uv.sub.powerLabel.Set(uv.T.Power)

	uv.sub.speedLabel = binding.NewString()
	uv.sub.speedLabel.Set(uv.T.Speed)
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
