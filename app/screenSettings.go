package uvs

import (
	"uvs/theme"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Let user to choose theme and language
// to-do ESP32 configuration
func settingsScreen(uv *UV_Station) fyne.CanvasObject {
	T := uv.T

	themeText := widget.NewLabelWithData(uv.sub.chooseThemeLabel)
	tdropdown := widget.NewSelect([]string{T.Light, T.Dark}, uv.parseTheme())

	langText := widget.NewLabelWithData(uv.sub.chooseLanguageLabel)
	ldropdown := widget.NewSelect([]string{T.EN, T.SR}, uv.parseLanguage())

	t := uv.config.StringWithFallback("THEME", T.Light)
	switch t {
	case "Светла", "Light":
		tdropdown.PlaceHolder = T.Light
	case "Тамна", "Dark":
		tdropdown.PlaceHolder = T.Dark
	}
	tdropdown.Refresh()

	l := uv.config.StringWithFallback("LANGUAGE", "English")
	switch l {
	case "English":
		ldropdown.PlaceHolder = T.EN
	case "Serbian":
		ldropdown.PlaceHolder = T.SR
	}
	ldropdown.Refresh()

	setIPentry := widget.NewEntry()
	setPortEntry := widget.NewEntry()

	uvSettings := widget.NewButton("Set IP and Port", func() {
		setIPtext := widget.NewLabel(uv.T.IP)
		setPortText := widget.NewLabel(uv.T.Port)

		setIPentry.Text = uv.IP
		setPortEntry.Text = uv.PORT

		uv.SET_WIN = uv.APP.NewWindow(uv.T.Settings)

		saveButton := widget.NewButton("Save", func() {
			uv.IP = setIPentry.Text
			uv.config.SetString("IP", uv.IP)

			uv.PORT = setPortEntry.Text
			uv.config.SetString("PORT", uv.PORT)
			uv.SET_WIN.Close()
		})

		cancelButton := widget.NewButton("Cancel", func() {
			uv.SET_WIN.Close()
		})

		ss := container.NewVBox(setIPtext, setIPentry)
		ss.Add(container.NewVBox(setPortText, setPortEntry))
		ss.Add(container.NewAdaptiveGrid(2, saveButton, cancelButton))

		uv.SET_WIN.SetContent(ss)

		width := uv.SET_WIN.Canvas().Size().Width * 3
		height := uv.SET_WIN.Canvas().Size().Height

		uv.SET_WIN.Resize(fyne.NewSize(width, height))
		uv.SET_WIN.Show()
	})

	settings := container.NewVBox(uvSettings)

	txt := container.NewAdaptiveGrid(2, themeText, langText)
	ctrl := container.NewAdaptiveGrid(2, tdropdown, ldropdown)

	settings.Add(container.NewVBox(txt, ctrl))

	return settings
}

func (uv *UV_Station) parseTheme() func(string) {
	return func(t string) {
		switch t {
		case "Светла", "Light":
			uv.config.SetString("THEME", "Светла")
			uv.APP.Settings().SetTheme(&theme.MyTheme{Theme: "Светла"})
		case "Тамна", "Dark":
			uv.config.SetString("THEME", "Тамна")
			uv.APP.Settings().SetTheme(&theme.MyTheme{Theme: "Тамна"})
		}
	}
}

func (uv *UV_Station) parseLanguage() func(string) {
	T := uv.T
	return func(l string) {
		switch l {
		case T.EN:
			uv.SetLanguage("English")
			uv.config.SetString("LANGUAGE", "English")
		case T.SR:
			uv.SetLanguage("Serbian")
			uv.config.SetString("LANGUAGE", "Serbian")
		}
		uv.refreshTitles()
	}
}

// change language without restart
func (uv *UV_Station) refreshTitles() {
	uv.sub.mainTab.Text = uv.T.Home
	uv.sub.consoleTab.Text = uv.T.Console
	uv.sub.settingsTab.Text = uv.T.Settings

	uv.sub.timerLabel.Set(uv.T.Timer)
	uv.sub.powerLabel.Set(uv.T.Power)
	uv.sub.speedLabel.Set(uv.T.Speed)

	uv.sub.chooseThemeLabel.Set(uv.T.ChooseTheme)
	uv.sub.chooseLanguageLabel.Set(uv.T.ChooseLanguage)

	uv.WIN.SetTitle(uv.T.Title)
}
