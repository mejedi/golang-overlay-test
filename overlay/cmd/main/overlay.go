package main

import _ "embed"

//go:embed .overlay-greeting
var overlay_greeting string

func init() {
	greeting = overlay_greeting
}
