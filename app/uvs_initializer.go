package uvs

import (
	"os"
	theme2 "uvs/theme"
	uvs "uvs/translation"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
)

var TIMER int = 5
var POWER int = 80
var SPEED int = 30

var TIMER_MAX int = 30
var POWER_MAX int = 100
var SPEED_MAX int = 100

type UV_Station struct {
	WIN    fyne.Window
	APP    fyne.App
	T      uvs.Translation
	config fyne.Preferences
	sub    Subitems

	timerBind binding.Float
	powerBind binding.Float
	speedBind binding.Float
}

func (uv *UV_Station) Start() {

	uv.sub.mainTab = container.NewTabItem(uv.T.Home, container.NewPadded(mainScreen(uv)))
	uv.sub.consoleTab = container.NewTabItem(uv.T.Console, container.NewPadded(consoleScreen(uv)))
	uv.sub.settingsTab = container.NewTabItem(uv.T.Settings, container.NewPadded(settingsScreen(uv)))

	tabs := container.NewAppTabs(uv.sub.mainTab, uv.sub.consoleTab, uv.sub.settingsTab)

	tabs.OnSelected = func(t *container.TabItem) {
		t.Content.Refresh()
	}

	uv.WIN.SetContent(tabs)

	width := uv.WIN.Canvas().Size().Width
	height := uv.WIN.Canvas().Size().Height

	if !(uv.APP.Driver().Device().IsMobile() && uv.APP.Driver().Device().IsBrowser()) {
		os.Setenv("FYNE_SCALE", "1")

		width *= 2
		height *= 1.2
	}

	uv.WIN.Resize(fyne.NewSize(width, height))
	uv.WIN.CenterOnScreen()
	uv.WIN.SetFixedSize(true)
	uv.WIN.SetMaster()
	uv.WIN.Show()

	uv.APP.Run()
}

func Initialize(id string) *UV_Station {
	a := app.NewWithID(id)

	thm := a.Preferences().StringWithFallback("THEME", "Light")

	a.Settings().SetTheme(&theme2.MyTheme{Theme: thm})

	uv := &UV_Station{
		APP:    a,
		WIN:    a.NewWindow(""),
		config: a.Preferences(),
	}

	uv.InitializeTranslations()
	uv.WIN.SetTitle(uv.T.Title)

	return uv
}
