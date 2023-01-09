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
	w := uv.WIN
	T := uv.T

	mainTab := container.NewTabItem(T.Home, container.NewPadded(mainScreen(uv)))
	consoleTab := container.NewTabItem(T.Console, container.NewPadded(consoleScreen(uv)))
	settingsTab := container.NewTabItem(T.Settings, container.NewPadded(settingsScreen(uv)))

	uv.sub.mainTab = mainTab
	uv.sub.consoleTab = consoleTab
	uv.sub.settingsTab = settingsTab

	tabs := container.NewAppTabs(mainTab, consoleTab, settingsTab)

	tabs.OnSelected = func(t *container.TabItem) {
		t.Content.Refresh()
	}

	w.SetContent(tabs)

	width := w.Canvas().Size().Width
	height := w.Canvas().Size().Height

	if !(uv.APP.Driver().Device().IsMobile() && uv.APP.Driver().Device().IsBrowser()) {
		os.Setenv("FYNE_SCALE", "1")

		width *= 2
		height *= 1.1
	}

	w.Resize(fyne.NewSize(width, height))
	w.CenterOnScreen()
	w.SetFixedSize(true)
	w.SetMaster()
	w.Show()

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
