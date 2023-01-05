package uvs

import (
	"uvs/theme"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func settingsScreen(uv *UV_Station) fyne.CanvasObject {
	T := uv.T

	themeText := canvas.NewText(T.ChooseTheme, nil)
	tdropdown := widget.NewSelect([]string{T.Light, T.Dark}, uv.parseTheme())

	langText := canvas.NewText(T.ChooseLanguage, nil)
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
			uv.config.SetString("LANGUAGE", "English")
			uv.Notify("Please restart app for changes to take effect")
		case T.SR:
			uv.config.SetString("LANGUAGE", "Serbian")
			uv.Notify("Рестартујте апликацију након промене језика")
		}
	}
}
