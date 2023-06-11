package uvs

import (
	"uvs/theme"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Let user to choose theme and language
// Configure Esp32 IP and Port
func settingsScreen(uv *UV_Station) fyne.CanvasObject {

	//define hostname entry settings
	hostLabel := widget.NewLabelWithData(uv.sub.setHostLabel)
	hostEntry := widget.NewEntry()
	hostEntry.Text = uv.HOSTNAME

	hostEntry.OnChanged = func(s string) {
		uv.config.SetString("HOSTNAME", s)
		uv.dial.SetUri(s)
	}

	// define widgets for theme and language settings
	themeLabel := widget.NewLabelWithData(uv.sub.chooseThemeLabel)
	themeSelect := widget.NewRadioGroup([]string{uv.T.Light, uv.T.Dark}, uv.parseTheme())
	themeSelect.Required = true

	if uv.config.StringWithFallback("THEME", "Light") == "Light" {
		themeSelect.Selected = uv.T.Light
	} else {
		themeSelect.Selected = uv.T.Dark
	}

	langLabel := widget.NewLabelWithData(uv.sub.chooseLanguageLabel)
	langSelect := widget.NewRadioGroup([]string{uv.T.EN, uv.T.SR}, uv.parseLanguage())
	langSelect.Required = true

	switch uv.config.StringWithFallback("LANGUAGE", "English") {
	case "Serbian":
		langSelect.Selected = uv.T.SR
	default:
		langSelect.Selected = uv.T.EN
	}

	uv.sub.themeSelect = themeSelect
	uv.sub.langSelect = langSelect

	host_settings := container.NewVBox(hostLabel, hostEntry)

	// theme and language settings
	thm := container.NewVBox(themeLabel, themeSelect)
	lng := container.NewVBox(langLabel, langSelect)

	ui_settings := container.NewVBox(lng, thm)

	// put all settings together on a screen
	settings := container.NewVBox(host_settings, ui_settings)

	return settings
}

func (uv *UV_Station) parseTheme() func(string) {
	return func(t string) {
		if t == uv.T.Light {
			uv.config.SetString("THEME", "Light")
			uv.sub.themeSelect.SetSelected(uv.T.Light)
			uv.APP.Settings().SetTheme(&theme.MyTheme{Theme: "Light"})
		} else {
			uv.config.SetString("THEME", "Dark")
			uv.sub.themeSelect.SetSelected(uv.T.Dark)
			uv.APP.Settings().SetTheme(&theme.MyTheme{Theme: "Dark"})
		}
	}
}

func (uv *UV_Station) parseLanguage() func(string) {
	return func(l string) {
		go func() {
			switch l {
			case uv.T.EN:
				uv.SetLanguage("English")
				uv.sub.langSelect.SetSelected(uv.T.EN)
			case uv.T.SR:
				uv.SetLanguage("Serbian")
				uv.sub.langSelect.SetSelected(uv.T.SR)
			}
			uv.refreshTitles()
		}()
	}
}

// change language without restart
func (uv *UV_Station) refreshTitles() {
	uv.sub.mainTab.Text = uv.T.Home
	uv.sub.settingsTab.Text = uv.T.Settings

	uv.sub.timerLabel.Set(uv.T.Timer)
	uv.sub.powerLabel.Set(uv.T.Power)
	uv.sub.speedLabel.Set(uv.T.Speed)

	uv.sub.setHostLabel.Set(uv.T.Hostname)

	uv.sub.chooseThemeLabel.Set(uv.T.ChooseTheme)
	uv.sub.themeSelect.Options = []string{uv.T.Light, uv.T.Dark}
	if uv.config.StringWithFallback("THEME", "Light") == "Light" {
		uv.sub.themeSelect.SetSelected(uv.T.Light)
	} else {
		uv.sub.themeSelect.SetSelected(uv.T.Dark)
	}

	uv.sub.chooseLanguageLabel.Set(uv.T.ChooseLanguage)
	uv.sub.langSelect.Options = []string{uv.T.EN, uv.T.SR}
	uv.sub.langSelect.Refresh()

	uv.WIN.SetTitle(uv.T.Title)
}
