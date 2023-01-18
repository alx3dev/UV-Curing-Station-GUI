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

	// define widgets for theme and language settings
	themeText := widget.NewLabelWithData(uv.sub.chooseThemeLabel)
	tdropdown := widget.NewSelect([]string{uv.T.Light, uv.T.Dark}, uv.parseTheme())

	langText := widget.NewLabelWithData(uv.sub.chooseLanguageLabel)
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

	ippName := uv.IP + " : " + uv.PORT

	if uv.IP == "" || uv.PORT == "" {
		ippName = "Set IP and Port"
	}

	// button to open IP and PORT settings window
	ip_port_settings := widget.NewButton(ippName, func() {

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

			ippName = uv.IP + " : " + uv.PORT
			uv.SET_WIN.Close()
		})

		// close settings window
		cancelButton := widget.NewButton("Cancel", func() {
			uv.SET_WIN.Close()
		})

		// put IP and PORT settings in a container
		ss := container.NewVBox(setIPtext, setIPentry)
		ss.Add(container.NewVBox(setPortText, setPortEntry))
		ss.Add(container.NewAdaptiveGrid(2, saveButton, cancelButton))

		uv.SET_WIN.SetContent(ss)

		// strech window for PC users
		if !uv.isMobile() {
			uv.SET_WIN.Resize(fyne.NewSize(
				uv.SET_WIN.Canvas().Size().Width*3,
				uv.SET_WIN.Canvas().Size().Height,
			))
		}
		uv.SET_WIN.Show()
	})

	// put all settings together on a screen
	txt := container.NewAdaptiveGrid(2, themeText, langText)
	ctrl := container.NewAdaptiveGrid(2, tdropdown, ldropdown)

	settings := container.NewVBox(ip_port_settings)
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
