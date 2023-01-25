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
	if n == theme.SizeNameScrollBar {
		return 0
	}
	if n == theme.SizeNameScrollBarSmall {
		return 0
	}
	return theme.DefaultTheme().Size(n)
}

func (m *MyTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {

	switch m.Theme {
	case "Тамна", "Dark":
		v = theme.VariantDark

	case "Светла", "Light":
		v = theme.VariantLight
	}
	return theme.DefaultTheme().Color(n, v)
}

func (m *MyTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}
