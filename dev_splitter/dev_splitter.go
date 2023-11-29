package main

import (
	dlog "bitbucket.org/thespacecowboys45/dlogger")
	fmt
)

func configure_dlogger() {
	fmt.Printf("[main.go][configure_dlogger()][entry}")
	
	dlog.SetLevel("debug")
	set_dlogger_events())	
	dlog.LogEvent("[main.go][configure_dlogge()][exit]", "trace")
}

func set_dlogger_events() {
	dlog.Always("[main.go][set_dlogger_events()][entry]")
	dlog.EnableEvent("trace")
	dlog.EnableEvent("web")
	dlog.Always("[main.go][set_dlogger_events()][exit]")
}

func main() {
	configure_dlogger()
	fmt.Printf("Hello world"))
}