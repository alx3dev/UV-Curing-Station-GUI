package theme

import (
	_ "embed"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

//go:embed OpenSans-Regular.ttf
var font []byte //(looking for better fonts...)
var my_font = &fyne.StaticResource{
	StaticName:    "OpenSans-Regular.ttf",
	StaticContent: font,
}

//go:embed Icon.ico
var icon []byte
var Ico = &fyne.StaticResource{
	StaticName:    "Icon.ico",
	StaticContent: icon,
}

type MyTheme struct {
	Theme string
}

var _ fyne.Theme = (*MyTheme)(nil)

func (m *MyTheme) Font(_ fyne.TextStyle) fyne.Resource {
	return my_font
}

func (m *MyTheme) Size(n fyne.ThemeSizeName) float32 {
	switch n {

	case theme.SizeNamePadding:
		return 2
	case theme.SizeNameScrollBar:
		return 0
	case theme.SizeNameScrollBarSmall:
		return 0
	case theme.SizeNameText:
		return 14
	case theme.SizeNameInputBorder:
		return 1
	}

	return theme.DefaultTheme().Size(n)
}

func (m *MyTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	switch m.Theme {
	case "Dark", "Тамна":
		v = theme.VariantDark

	case "Light", "Светла":
		v = theme.VariantLight
	}

	switch n {
	case theme.ColorNameSeparator:
		return color.Transparent

	case theme.ColorNameButton:
		if v == theme.VariantDark {
			return color.RGBA{R: 33, G: 38, B: 33, A: 255}
		}
	case theme.ColorNameMenuBackground:
		if v == theme.VariantDark {
			return color.RGBA{R: 30, G: 30, B: 30, A: 255}
		}

	case theme.ColorNamePrimary:
		if v == theme.VariantDark {
			return theme.WarningColor()
		}

	case theme.ColorNameFocus:
		return color.Transparent

	case theme.ColorNameShadow:
		return color.Transparent

	case theme.ColorNameHover:
		return color.Transparent

	case theme.ColorNameInputBorder:
		return color.Transparent

	case theme.ColorNameInputBackground:
		return color.Transparent
	}

	return theme.DefaultTheme().Color(n, v)
}

func (m *MyTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}
