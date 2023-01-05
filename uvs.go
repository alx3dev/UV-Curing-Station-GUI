package main

import (
	uvs "uvs/app"
)

func main() {
	uvsApp := uvs.Initialize("curing.station.uv")

	uvsApp.CheckOnStartNotification()

	uvsApp.Start()
}
