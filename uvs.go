package main

import (
	uvs "uvs/app"
)

func main() {
	uvsApp := uvs.Initialize()

	uvsApp.CheckOnStartNotification()

	uvsApp.Start()
}

