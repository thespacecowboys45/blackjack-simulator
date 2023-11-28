package main
// Reference origin: 
//
// https://bitbucket.org/thespacecowboys45/dlogger/src/master/



import(
	//"fmt"
	dlog "bitbucket.org/thespacecowboys45/dlogger"

)
/*

func dlog(msg string) {
	fmt.Printf(msg)
}
*/

func enableEvents() {
	
}

func configureDlogger() {
	dlog.SetLevel("debug")

	// event types available so far	
	dlog.EnableEvent("basic")
	//dlog.EnableEvent("bets")
	//dlog.EnableEvent("bettingstrategy")
	//dlog.EnableEvent("bankroll")
	//dlog.EnableEvent("foo")
	//dlog.EnableEvent("metrics")
	//dlog.EnableEvent("trace")
}