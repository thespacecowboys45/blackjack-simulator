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
	"strconv"
	"math"
	"math/big"
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

func convertDecimalToHex(number string) {
	fmt.Printf("[convert_base.go][convertToHex()][entry]\n")
	
	// This is beginning to hit the idea - for blackjack
	//var number string = "15"
	var toBase string= "0123456789ABCDEF"
	var inBase string= "0123456789"
	converted, _, _ := bc.BaseToBase(number, inBase, toBase)
	fmt.Printf("convertToHex 1: %s\n", converted) // should be 'dh'	
}


func convertToHand(number string) {
	fmt.Printf("[convert_base.go][convertToHand()][entry]\n")
	
	// This is beginning to hit the idea - for blackjack
	//var number string = "20"
	var toBase string= "hsd"
	var inBase string= "012"
	converted, _, _ := bc.BaseToBase(number, inBase, toBase)
	fmt.Printf("convertToHand 1: %s\n", converted) // should be 'dh'
	
	// @TODO - add a 'precision' value - in order to indicate how large the number could be
	// trick out the system - bitmask a '1' value in one bit higher than the highest value,
	// causing the outcome result to have a 1's place in the highest value, and trickling
	// down zero's (0) to the rest of the chain
}


func convertDecimalToHand(number string) {
	fmt.Printf("[convert_base.go][convertDecimalToHand()][entry] number=%s\n", number)
	
	// This is beginning to hit the idea - for blackjack
	//var number string = "20"
	var toBase string= "hsd"
	var inBase string= "0123456789"
	converted, _, _ := bc.BaseToBase(number, inBase, toBase)
	fmt.Printf("convertDecimalToHand 1: %s\n", converted) // should be 'dh'
	
	// @TODO - add a 'precision' value - in order to indicate how large the number could be
	// trick out the system - bitmask a '1' value in one bit higher than the highest value,
	// causing the outcome result to have a 1's place in the highest value, and trickling
	// down zero's (0) to the rest of the chain
}


func writeBigIntToString() {
	fmt.Printf("[convert_base.go][writeBigIntToString()][entry]\n")
	bigInt := big.NewFloat(math.Pow(3, 150))
	//bigInt := big.NewInt(123456789)
	bigStr := bigInt.String()

	fmt.Println("String value is " , bigStr)

}

// ripped from: https://stackoverflow.com/questions/30182129/calculating-large-exponentiation-in-golang
func powBig(a, n int) *big.Int{
    tmp := big.NewInt(int64(a))
    res := big.NewInt(1)
    for n > 0 {
        temp := new(big.Int)
        if n % 2 == 1 {
            temp.Mul(res, tmp)
            res = temp
        }
        temp = new(big.Int)
        temp.Mul(tmp, tmp)
        tmp = temp
        n /= 2
    }
    return res
}


func main() {
	fmt.Printf("[convert_base.go][main][entry]\n")
	//example1()
	//tryIt()
/*
 works, commenting out	
	
	convertDecimalToHex("15")

	// not ready for this yet
	convertToHand("20")
	convertToHand("1")
	convertToHand("101")
	convertToHand("0101") // produces same output result as above
	convertToHand("1111111111")
	convertToHand("1000000000")
	
	convertDecimalToHand("0")
	convertDecimalToHand("1")
	convertDecimalToHand("2")
	convertDecimalToHand("3")
	convertDecimalToHand("4")
	convertDecimalToHand("5")
	convertDecimalToHand("6")
	convertDecimalToHand("7")
	convertDecimalToHand("8")
	convertDecimalToHand("9")
	convertDecimalToHand("10")
	
	// start somewhere and print out 10 values
	//startAt := 20
	var num_values float64 = 0
	startAt := math.Pow(2,4) // 2^4 = 16
	//startAt = math.Pow(3,100) // 3^10 - does not work with this large number to strconv.FormatInt()
	
	fmt.Printf("startAt: %v", startAt)
	
	for i:=startAt; i< (startAt + num_values); i++ {
		s1 := strconv.FormatInt(int64(i), 10)
  		//s2 := strconv.Itoa(i)
  		//fmt.Printf("[%d] Looper: %v, %v\n", i, s1, s2)
  		fmt.Printf("[%v] Looper: %v\n", i, s1)
  		convertDecimalToHand(s1)
	}	
*/
	
	
	/*****
	 *
	 * Stopping point - at this point I have a methodology to do what I want to do with this idea
	 *
	 ****/
	
	
	//writeBigIntToString()	
	res := powBig(3, 181)
	fmt.Printf("result: %v\n", res)
	str := res.String()
	fmt.Printf("result string: %s\n", str)
	convertDecimalToHand(str)
	
}