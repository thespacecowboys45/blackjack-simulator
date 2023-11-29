package main

import (
	dlog "bitbucket.org/thespacecowboys45/dlogger"
	"fmt"
	)

func configure_dlogger() {
	fmt.Printf("[main.go][configure_dlogger()][entry]\n")
	
	dlog.SetLevel("debug")
	set_dlogger_events()
	dlog.LogEvent("[main.go][configure_dlogger()][exit]", "trace")
}

func set_dlogger_events() {
	dlog.Always("[main.go][set_dlogger_events()][entry]")
	dlog.EnableEvent("trace")
	dlog.EnableEvent("web")
	dlog.Always("[main.go][set_dlogger_events()][exit]")
}

func splitSet(method string, s Set) (Set, Set) {
	dlog.LogEvent("[dev_splitter.go][splitSet][entry]", "trace")
	
	var d1, d2 Set
	
	dlog.Info("[splitSet][method %s s %v", method, s)
	switch method {
		case "mid":
			dlog.Info("[splitSet][chose midpoint method]")
			// actually find the midpoint and return
			//
			// NOTE: For now this does not matter if points are reversed numerically
			mid := (s.p1 + s.p2) / 2
			d1 = Set{s.p1,mid}
			d2 = Set{mid,s.p2}

			dlog.Debug("[splitSet][midpoint: %v]", mid)
			break
		default:
			dlog.Error("[splitSet][invalid method: %s]", method)
			break
	}
	
	dlog.LogEvent("[dev_splitter.go][splitSet][exit]", "trace")
	
	
	return d1, d2
}

func main() {
	configure_dlogger()
	fmt.Printf("Hello world\n")
	
	//s := make(Set,1)
	s := Set{}
	dlog.Info("s: %v ", s)
	
	s.p1 = 5
	s.p2 = 10
	
	dlog.Info("s now: %v ", s)
	
	d1, d2 := splitSet("mid", s)
	dlog.Info("S: %v split into %v and %v", s, d1, d2)
	
}