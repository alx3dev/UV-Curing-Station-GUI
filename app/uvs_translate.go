package uvs

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
