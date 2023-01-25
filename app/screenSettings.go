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
	tdropdown := widget.NewSelect([]string{uv.T.Light, uv.T.Dark}, uv.parseTheme())

	langLabel := widget.NewLabelWithData(uv.sub.chooseLanguageLabel)
	ldropdown := widget.NewSelect([]string{uv.T.EN, uv.T.SR}, uv.parseLanguage())

	// define theme dropdown menu
	t := uv.config.StringWithFallback("THEME", uv.T.Light)
	switch t {
	case "Светла", "Light":
		tdropdown.PlaceHolder = uv.T.Light
	case "Тамна", "Dark":
		tdropdown.PlaceHolder = uv.T.Dark
	}
	tdropdown.Refresh()

	// define language dropdown menu
	l := uv.config.StringWithFallback("LANGUAGE", "English")
	switch l {
	case "English":
		ldropdown.PlaceHolder = uv.T.EN
	case "Serbian":
		ldropdown.PlaceHolder = uv.T.SR
	}
	ldropdown.Refresh()

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

		uv.SET_WIN = uv.APP.NewWindow(uv.T.Settings)

		// save IP settings and close window
		saveButton := widget.NewButton("Save", func() {
			uv.IP = setIPentry.Text
			uv.PORT = setPortEntry.Text

			uv.config.SetString("IP", uv.IP)
			uv.config.SetString("PORT", uv.PORT)

			uv.sub.configButton.Text = uv.IP + ":" + uv.PORT
			uv.sub.configButton.Refresh()
			uv.SET_WIN.Close()
		})

		// close settings window
		cancelButton := widget.NewButton("Cancel", func() {
			uv.SET_WIN.Close()
		})

		// put IP and PORT settings in a container
		ipps := container.NewVBox(setIPtext, setIPentry)
		ipps.Add(container.NewVBox(setPortText, setPortEntry))
		ipps.Add(container.NewAdaptiveGrid(2, saveButton, cancelButton))

		uv.SET_WIN.SetContent(ipps)

		// strech window for PC users
		if !uv.isMobile() {
			uv.SET_WIN.Resize(fyne.NewSize(
				uv.SET_WIN.Canvas().Size().Width*3,
				uv.SET_WIN.Canvas().Size().Height,
			))
		}
		uv.SET_WIN.Show()
	})

	uv.sub.configButton = ip_port_settings

	ipp := container.NewGridWithRows(2, ippsLabel, ip_port_settings)

	// theme and language settings
	thm := container.NewGridWithRows(2, themeLabel, tdropdown)
	lng := container.NewGridWithRows(2, langLabel, ldropdown)

	ui_settings := container.NewAdaptiveGrid(2, thm, lng)

	// put all settings together on a screen
	settings := container.NewVBox(ui_settings, ipp)

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
	uv.sub.settingsTab.Text = uv.T.Settings

	uv.sub.timerLabel.Set(uv.T.Timer)
	uv.sub.powerLabel.Set(uv.T.Power)
	uv.sub.speedLabel.Set(uv.T.Speed)

	uv.sub.chooseThemeLabel.Set(uv.T.ChooseTheme)
	uv.sub.chooseLanguageLabel.Set(uv.T.ChooseLanguage)

	uv.sub.configLabel.Set(uv.T.Configuration)

	uv.WIN.SetTitle(uv.T.Title)
}
