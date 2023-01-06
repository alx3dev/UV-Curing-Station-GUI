package uvs

import (
	"fyne.io/fyne/v2"
)

func (uv *UV_Station) Notify(message string) {
	uv.APP.SendNotification(fyne.NewNotification(uv.T.Title, message))
}

func (uv *UV_Station) NotifyOnNextStart(message string) {
	msg := uv.APP.Preferences().StringWithFallback("ON_NEXT_START_MESSAGE", "")

	if msg != "" {
		msg += "\n"
	}
	msg += message
	uv.APP.Preferences().SetString("ON_NEXT_START_TITLE", "UV Station")
	uv.APP.Preferences().SetString("ON_NEXT_START_MESSAGE", msg)
}

func (uv *UV_Station) CheckOnStartNotification() {
	n := uv.APP.Preferences().StringWithFallback("ON_NEXT_START_TITLE", "")
	msg := uv.APP.Preferences().StringWithFallback("ON_NEXT_START_MESSAGE", "")

	if n != "" {
		uv.Notify(msg)
		uv.ClearNotifications()
	}
}

func (uvs *UV_Station) ClearNotifications() {
	uvs.APP.Preferences().SetString("ON_NEXT_START_TITLE", "")
	uvs.APP.Preferences().SetString("ON_NEXT_START_MESSAGE", "")
}

