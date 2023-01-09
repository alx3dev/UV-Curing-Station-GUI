package uvs

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
)

func (uv *UV_Station) InitializeTranslations() {
	language :=
		uv.config.StringWithFallback("LANGUAGE", "English")

	// so we don't need to translate everything,
	// english string will be used if translation not found
	uv.SetLanguage("English")

	uv.SetLanguage(language)
}

func (uv *UV_Station) SetLanguage(lang string) {
	uv.config.SetString("LANGUAGE", lang)
	uv.T.ImplementTranslation(lang)
}

// Keep pointers for automatic refresh
// after language change
type Subitems struct {
	mainTab     *container.TabItem
	consoleTab  *container.TabItem
	settingsTab *container.TabItem

	timerLabel          binding.String
	powerLabel          binding.String
	speedLabel          binding.String
	chooseThemeLabel    binding.String
	chooseLanguageLabel binding.String
}
