package uvs

import (
	"fyne.io/fyne/v2"
)

func (uv *UV_Station) Notify(message string) {
	uv.APP.SendNotification(fyne.NewNotification(uv.T.Title, message))
}

func (uv *UV_Station) NotifyOnNextStart(message string) {
	msg := uv.config.StringWithFallback("ON_NEXT_START_MESSAGE", "")

	if msg != "" {
		msg += "\n"
	}
	msg += message
	uv.config.SetString("ON_NEXT_START_TITLE", "UV Station")
	uv.config.SetString("ON_NEXT_START_MESSAGE", msg)
}

func (uv *UV_Station) CheckOnStartNotification() {
	n := uv.config.StringWithFallback("ON_NEXT_START_TITLE", "")
	msg := uv.config.StringWithFallback("ON_NEXT_START_MESSAGE", "")

	if n != "" {
		uv.Notify(msg)
		uv.ClearNotifications()
	}
}

func (uv *UV_Station) ClearNotifications() {
	uv.config.SetString("ON_NEXT_START_TITLE", "")
	uv.config.SetString("ON_NEXT_START_MESSAGE", "")
}
