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

	// define entry settings for IP and PORT
	setIPentry := widget.NewEntry()
	setPortEntry := widget.NewEntry()

	ippsLabel := widget.NewLabelWithData(uv.sub.configLabel)
	ibutTitle := uv.IP + " : " + uv.PORT

	// button to open IP and PORT settings window
	ip_port_settings := widget.NewButton(ibutTitle, func() {

		setIPentry.Text = uv.IP
		setPortEntry.Text = uv.PORT

		setIPtext := widget.NewLabel(uv.T.IP)
		setPortText := widget.NewLabel(uv.T.Port)

		setWin := uv.APP.NewWindow(uv.T.Settings)

		// save IP settings and close window
		saveButton := widget.NewButton("Save", func() {
			uv.IP = setIPentry.Text
			uv.PORT = setPortEntry.Text

			uv.dial.SetUri(uv.IP + ":" + uv.PORT)

			uv.config.SetString("IP", uv.IP)
			uv.config.SetString("PORT", uv.PORT)

			go func() {
				uv.sub.configButton.Text = uv.IP + ":" + uv.PORT
				uv.sub.configButton.Refresh()
			}()
			setWin.Close()
		})

		// close settings window
		cancelButton := widget.NewButton("Cancel", func() {
			setWin.Close()
		})

		// put IP and PORT settings in a container
		ipps := container.NewVBox(setIPtext, setIPentry)
		ipps.Add(container.NewVBox(setPortText, setPortEntry))
		ipps.Add(container.NewAdaptiveGrid(2, saveButton, cancelButton))

		setWin.SetContent(ipps)

		// strech window for PC users
		if !uv.isMobile() {
			setWin.Resize(fyne.NewSize(
				setWin.Canvas().Size().Width*3,
				setWin.Canvas().Size().Height,
			))
		}
		setWin.Show()
	})

	uv.sub.configButton = ip_port_settings

	ipp := container.NewGridWithRows(2, ippsLabel, ip_port_settings)

	// theme and language settings
	thm := container.NewVBox(themeLabel, themeSelect)
	lng := container.NewVBox(langLabel, langSelect)

	ui_settings := container.NewVBox(lng, thm)

	// put all settings together on a screen
	settings := container.NewVBox(ipp, ui_settings)

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

	uv.sub.configLabel.Set(uv.T.Configuration)

	uv.WIN.SetTitle(uv.T.Title)
}
