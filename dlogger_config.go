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

func configureDlogger() {
	dlog.SetLevel("debug")
}