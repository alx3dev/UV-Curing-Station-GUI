package uvs

import (
	"uvs/theme"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

// Let user to choose theme and language
// to-do ESP32 configuration
func settingsScreen(uv *UV_Station) fyne.CanvasObject {
	T := uv.T

	uv.sub.chooseThemeLabel = binding.NewString()
	uv.sub.chooseThemeLabel.Set(T.ChooseTheme)

	uv.sub.chooseLanguageLabel = binding.NewString()
	uv.sub.chooseLanguageLabel.Set(T.ChooseLanguage)

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

	settings := container.NewVBox(themeText, tdropdown)
	settings.Add(container.NewVBox(langText, ldropdown))

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
			uv.Notify("Please restart app for changes to take effect")
		case T.SR:
			uv.SetLanguage("Serbian")
			uv.config.SetString("LANGUAGE", "Serbian")
			uv.Notify("Рестартујте апликацију након промене језика")
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
