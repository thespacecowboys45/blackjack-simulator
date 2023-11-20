package main
/*****
 *
 * @author dxb The Space Cowboy
 * @date November '23
 *
 * DESCRIPTION:
 *
 * Links:
 *   https://pkg.go.dev/math/big
 *   https://github.com/ctison/baseconverter
 *
 ****/
 
import(
	"fmt"

	bc "github.com/chtison/baseconverter"
)

/**
 * potential conversion functions:
 *
 * func BaseToDecimal(number string, inBase string) (result *big.Int, err error) {
 * func DecimalToBase(number *big.Int, toBase string) (result string, err error) {
 * func UInt64ToBase(number uint64, toBase string) (result string, err error) {
 * func BaseToBase(number string, inBase string, toBase string) (result string, e1, e2 error) {
 *
 */


func example1() {
	fmt.Println("Example 1")
	
	/* For example, you can convert a decimal number to base 16: */
	nbrInBase16, _ := bc.UInt64ToBase(51966, "0123456789ABCDEF")
	fmt.Println(nbrInBase16)
	
	
	fmt.Println("Example 2")
	
	/* Or convert back a number in base "01" (base 2) to base 10: */
	// 32 + 8 + 2 = 42
	nbr, _ := bc.BaseToDecimal("101010", "01")
	fmt.Println(nbr)	
	
	fmt.Println("Example 3")
	
	/* Or convert a number from any base to any other: */
	var number string = "🌴🐭🌞🌝🍀💎💎🌝🐱🍀💜🍀🐵🐱🐭🌴🐼🌵🍀🐱💎🐼"
	var inBase string = "🌵🐱🚗🌍🌞🍀💎💰🐼🍋🐵🌴💜🐭🌝"
	var toBase string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ !"
	converted, _, _ := bc.BaseToBase(number, inBase, toBase)
	fmt.Println(converted)
}

func tryIt() {
	fmt.Printf("[tryIt()][entry]\n")
	
	// convert from binary to decimal
	var number string = "1010"
	var inBase string = "01"
	var toBase string = "0123456789"
	converted, _, _ := bc.BaseToBase(number, inBase, toBase)
	fmt.Printf("Converted 1: %s\n", converted)
	

	// CONVERT TO DECIMAL IN THE FOLLOWING EXAMPLES:
	toBase = "0123456789"
	
	// convert from binary to decimal
	number = "yxyx" // should be 10
	inBase = "xy"
	converted, _, _ = bc.BaseToBase(number, inBase, toBase)
	fmt.Printf("Converted 2: %s\n", converted)
	
	number = "xyxy" // should be 5 -> base == 01, number == 0101, representative of the runes for inBase
	inBase = "xy"
	converted, _, _ = bc.BaseToBase(number, inBase, toBase)
	fmt.Printf("Converted 3: %s\n", converted)

	// what happens	
	number = "xyxy" // should be 5 -> base == 01, number == 0101, representative of the runes for inBase
	inBase = "xyz"
	converted, _, _ = bc.BaseToBase(number, inBase, toBase)
	fmt.Printf("Converted 4: %s\n", converted)
	
	
	// what happens	
	number = "xxxy" // should be 1
	inBase = "xyz"
	converted, _, _ = bc.BaseToBase(number, inBase, toBase)
	fmt.Printf("Converted 5: %s\n", converted)
	

	// what happens	
	number = "xxxz" // should be 2
	inBase = "xyz"
	converted, _, _ = bc.BaseToBase(number, inBase, toBase)
	fmt.Printf("Converted 6: %s\n", converted)
	
	// ----------- other way around now


	// CONVERT TO my custome base IN THE FOLLOWING EXAMPLES:
	number = "10"
	toBase = "01"
	inBase = "01"
	converted, _, _ = bc.BaseToBase(number, inBase, toBase)
	fmt.Printf("Converted 7: %s\n", converted) // should be '10'
	
	
	number = "10"
	toBase = "xy"
	inBase = "01"
	converted, _, _ = bc.BaseToBase(number, inBase, toBase)
	fmt.Printf("Converted 8: %s\n", converted) // should be 'yx'


	number = "20"
	toBase = "xyz"
	inBase = "012"
	converted, _, _ = bc.BaseToBase(number, inBase, toBase)
	fmt.Printf("Converted 9: %s\n", converted) // should be 'zx'


	// This is beginning to hit the idea - for blackjack
	number = "20"
	toBase = "hsd"
	inBase = "012"
	converted, _, _ = bc.BaseToBase(number, inBase, toBase)
	fmt.Printf("Converted 10: %s\n", converted) // should be 'dh'


	//
}

func convertToHex() {
	fmt.Printf("[convert_base.go][convertToHex()][entry]\n")
	
	// This is beginning to hit the idea - for blackjack
	var number string = "1"
	var toBase string= "0123456789ABCDEF"
	var inBase string= "0123456789"
	converted, _, _ := bc.BaseToBase(number, inBase, toBase)
	fmt.Printf("convertToHex 1: %s\n", converted) // should be 'dh'	
}


func convertToHand() {
	fmt.Printf("[convert_base.go][convertToHand()][entry]\n")
	
	// This is beginning to hit the idea - for blackjack
	var number string = "20"
	var toBase string= "hsd"
	var inBase string= "012"
	converted, _, _ := bc.BaseToBase(number, inBase, toBase)
	fmt.Printf("convertToHand 1: %s\n", converted) // should be 'dh'
	
	// This is beginning to hit the idea - for blackjack
	number = "1111111111"
	toBase = "hsd"
	inBase = "012"
	converted, _, _ = bc.BaseToBase(number, inBase, toBase)
	fmt.Printf("convertToHand 2: %s\n", converted) // should be 'dh'
	
	
}

func main() {
	fmt.Printf("[convert_base.go][main][entry]")
	//example1()
	//tryIt()
	
	convertToHex()

	// not ready for this yet
	//convertToHand()
	
}