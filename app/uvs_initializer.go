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

type UV_Station struct {
	WIN     fyne.Window // main window
	SET_WIN fyne.Window // settings window (for ip and port)
	APP     fyne.App
	T       uvs.Translation  // translateable params
	config  fyne.Preferences // shortuct to fyne preferences
	sub     Subitems         // items that need refresh on language change

	IP   string // esp32 access-point IP address
	PORT string // esp32 server port

	system uint8

	timerBind binding.Float
	powerBind binding.Float
	speedBind binding.Float
}

func Initialize(id string) *UV_Station {
	a := app.NewWithID(id)

	thm := a.Preferences().StringWithFallback("THEME", "Тамна")

	a.Settings().SetTheme(&theme2.MyTheme{Theme: thm})
	a.SetIcon(theme2.Ico)

	uv := &UV_Station{
		APP:    a,
		WIN:    a.NewWindow(""),
		config: a.Preferences(),
		system: getOS(),
	}

	uv.onInitialize()

	return uv
}

func (uv *UV_Station) Start() {

	uv.sub.mainTab = container.NewTabItem(uv.T.Home, container.NewPadded(mainScreen(uv)))
	uv.sub.settingsTab = container.NewTabItem(uv.T.Settings, container.NewPadded(settingsScreen(uv)))

	tabs := container.NewAppTabs(uv.sub.mainTab, uv.sub.settingsTab)

	tabs.OnSelected = func(t *container.TabItem) {
		t.Content.Refresh()
	}

	uv.WIN.SetContent(tabs)

	width := uv.WIN.Canvas().Size().Width
	height := uv.WIN.Canvas().Size().Height

	if !uv.isMobile() {
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

// TO-DO: add "how-to" window on first start
func (uv *UV_Station) onInitialize() {
	if uv.system == 0 {
		println("Operating System is not supported.")
		uv.APP.Quit()
	}
	uv.InitializeTranslations()
	uv.WIN.SetTitle(uv.T.Title)

	first := uv.config.StringWithFallback("FIRST_INIT", "")
	if first == "" {
		//go uv.Notify("Welcome Message")
		uv.config.SetString("FIRST_INIT", "false")
	}
}
