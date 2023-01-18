package uvs

type Translation struct {
	Title          string
	Home           string
	Settings       string
	Console        string
	Timer          string
	Power          string
	Speed          string
	ChooseTheme    string
	Light          string
	Dark           string
	ChooseLanguage string
	EN             string
	SR             string
	Send           string
	IP             string
	Port           string
}

func (t *Translation) ImplementTranslation(l string) {
	switch l {

	case "English":

		t.Title = "UV Curing Station"
		t.Home = "Home"
		t.Settings = "Settings"
		t.Console = "Console"
		t.Timer = "Time  "
		t.Power = "Power"
		t.Speed = "Speed"
		t.ChooseTheme = "Choose Theme"
		t.Light = "Light"
		t.Dark = "Dark"
		t.ChooseLanguage = "Choose Language"
		t.EN = "English"
		t.SR = "Serbian"
		t.Send = "Send"
		t.IP = "IP Address"
		t.Port = "Port"

	case "Serbian":

		t.Title = "УВ Станица"
		t.Home = "Насловна"
		t.Settings = "Подешавања"
		t.Console = "Терминал"
		t.Timer = "Време"
		t.Power = "Снага "
		t.Speed = "Брзина"
		t.ChooseTheme = "Изаберите Тему"
		t.Light = "Светла"
		t.Dark = "Тамна"
		t.ChooseLanguage = "Изаберите Језик"
		t.EN = "Енглески"
		t.SR = "Српски"
		t.Send = "Пошаљи"
		t.IP = "Адреса"
		t.Port = "Порт"

	}
}
