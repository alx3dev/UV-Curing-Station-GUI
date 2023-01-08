package uvs

import (
	"fyne.io/fyne/v2"
)

// send system notification
func (uv *UV_Station) Notify(message string) {
	uv.APP.SendNotification(fyne.NewNotification(uv.T.Title, message))
}

// show notification on next start
func (uv *UV_Station) NotifyOnNextStart(message string) {
	msg := uv.config.StringWithFallback("ON_NEXT_START_MESSAGE", "")

	if msg != "" {
		msg += "\n"
	}

	msg += message
	uv.config.SetString("ON_NEXT_START_TITLE", uv.T.Title)
	uv.config.SetString("ON_NEXT_START_MESSAGE", msg)
}

// check for "on-app-start" notifications
func (uv *UV_Station) CheckOnStartNotification() {
	n := uv.config.StringWithFallback("ON_NEXT_START_TITLE", "")
	msg := uv.config.StringWithFallback("ON_NEXT_START_MESSAGE", "")

	if n != "" {
		uv.Notify(msg)
		uv.ClearNotifications()
	}
}

func (uv *UV_Station) ClearNotifications() {
	uv.config.RemoveValue("ON_NEXT_START_TITLE")
	uv.config.RemoveValue("ON_NEXT_START_MESSAGE")
}
