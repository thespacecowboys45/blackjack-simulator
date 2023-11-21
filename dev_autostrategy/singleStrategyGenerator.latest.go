package main
/*****
 *
 * @date Nov '23
 * @author dxb The Space Cowboy
 *
 * DESCRIPTION:
 *
 *   Concepts brought from dev_learning_blocks/convert_base.go
 *
 * Generate a single strategy based on a large number.  The large
 * number is somewhere between 0 and the maximum number of possibilities
 * for a playing strategy, given our matrix-based idea of running a
 * strategy.
 *
 * This may also output it to a file (tbd)
 *
 * Links:
 *   https://pkg.go.dev/math/big
 *   https://github.com/ctison/baseconverter
 *
 ****/
 
import(
	"os"
	"fmt"
	"math/big"
	bc "github.com/chtison/baseconverter"
)

func convertDecimalToHand(number string) string {
	fmt.Printf("[convert_base.go][convertDecimalToHand()][entry]\n number=%s\n", number)
	
	// This is beginning to hit the idea - for blackjack
	//var number string = "20"
	var toBase string= "hsd"
	var inBase string= "0123456789"
	converted, _, _ := bc.BaseToBase(number, inBase, toBase)
	fmt.Printf("convertDecimalToHand 1: %s\n", converted) // should be 'dh'
	
	return converted
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


// Will slice out a portion of the string based on the "row" we're dealing
// with in the strategy generator.  Should auto-calculate where in the 
// large string we need to slice out and return that info
func sliceOutRowInfo(n int, input string) string {
	fmt.Printf("[singleStrategyGenerator.go][sliceOutRowInfo][entry]\n")
	fmt.Printf("Input: %s\n", input)
	fmt.Printf("Row: %d\n", n)
	
	return "testing"
}
 
 
/**
 * Displays the first strategy, starting at one
 */
func displayFirstStrategy() {
	// let's try a couple things.  First, what does 'one' look like ? 
	res := powBig(1, 1) // 
	fmt.Printf("result: %v\n", res)
	
	// Convert to string
	str := res.String()
	fmt.Printf("result string: %s\n", str)
	
	// Use our base conversion thingie to convert into the 'pseudo string we want'
	s := convertDecimalToHand(str)
	fmt.Printf("\nFirst Strategy looks like: %s\n", s)	
} 

/**
 * Displays the first strategy, starting at one
 */
func displayFinalStrategy() {
	// let's try a couple things.  First, what does 'one' look like ? 
	res := powBig(3, 180) // 
	fmt.Printf("result: %v\n", res)
	
	// Convert to string
	str := res.String()
	fmt.Printf("result string: %s\n", str)

	// subtract one from res
	resMinusOne := res.Sub(res, big.NewInt(1))
	
	// Convert to string
	str = resMinusOne.String()
	fmt.Printf("result minus one string: %s\n", str)
	
	// Use our base conversion thingie to convert into the 'pseudo string we want'
	s := convertDecimalToHand(str)
	fmt.Printf("\nFinal Strategy looks like: %s\n", s)	
} 

/**
 * Displays a specified strategy
 */
func displayStrategyNumber(bigInt *big.Int) {
	fmt.Printf("[singleStrategyGenerator.go][displayStrategyNumber][entry]\n")
	fmt.Printf("Display for strategy #%v\n", bigInt)
	// let's try a couple things.  First, what does 'one' look like ? 
//	res := powBig(3, n) // 
//	fmt.Printf("result: %v\n", res)
	

	fmt.Printf("bigInt string # bits: %d\n", bigInt.Bits())

	
	// Convert to string
	str := bigInt.String()
	fmt.Printf("bigInt string length: %d\n", len(str))
	fmt.Printf("bigInt string: %s\n", str)
	
	// Use our base conversion thingie to convert into the 'pseudo string we want'
	s := convertDecimalToHand(str)
	fmt.Printf("\nStrategy %v\nlooks like: %s\n", bigInt, s)	
	fmt.Printf("--------------------------\n")
} 

/**
 * Displays a specified strategy, however this is shifted over
 * so that we print out all possible values for the entire matrix
 */
func displayShiftedStrategyNumber(bigInt *big.Int, nShift uint) {
	fmt.Printf("[singleStrategyGenerator.go][displayShiftedStrategyNumber][entry]\n")
	fmt.Printf("Display for strategy #%v\n", bigInt)
	fmt.Printf("Shift by %d\n", nShift)
	// let's try a couple things.  First, what does 'one' look like ? 
//	res := powBig(3, n) // 
//	fmt.Printf("result: %v\n", res)
	
	shifted := bigInt.Lsh(bigInt, nShift)
	fmt.Printf("shifted: %v\n", shifted)
	
	// Convert to string
	str := bigInt.String()
	fmt.Printf("result string length: %d\n", len(str))
	fmt.Printf("result string: %s\n", str)
	
	// Use our base conversion thingie to convert into the 'pseudo string we want'
	s := convertDecimalToHand(str)
	fmt.Printf("\nStrategy %v\nlooks like: %s\n", bigInt, s)	
	fmt.Printf("--------------------------\n")
} 



func testingStill2() {
	// let's try a couple things.  First, what does one look like ? 
	//res := powBig(3, 181) // what we may have to deal with as maximum # of possibilities
	res := powBig(1, 1) // start halfway
	fmt.Printf("result: %v\n", res)
	
	// Convert to string
	str := res.String()
	fmt.Printf("result string: %s\n", str)
	
	shifted := res.Lsh(res, 180)
	fmt.Printf("shifted: %v\n", shifted)	
	
	// Use our base conversion thingie to convert into the 'pseudo string we want'
	s := convertDecimalToHand(str)
	fmt.Printf("Final: %s\n", s)	
} 

func printMaxPossibilities() {
	fmt.Printf("[singleStrategyGenerator.go][printMaxPossibilities()][entry]\n")
	
	res := powBig(3, 180) // start halfway
	fmt.Printf(" max # of possibilities result: %v\n", res)
	
	// Convert to string
	str := res.String()
	fmt.Printf(" max # possibilities as a string: %s\n", str)
	
	fmt.Printf("\n---\nThis is what we are shooting for.  To run all these possibilities.  Yes!  DWAV quantum computing, here we come.\n---\n\n")
	
}
 
func main() {
 	fmt.Printf("[singleStrategyGenerator.go][main][entry]\n")
 	
 	printMaxPossibilities()
 	
 	// in-dev
/*
 //fun
  	
	for i:=0; i<11; i++ {
		displayStrategyNumber(big.NewInt(int64(i)))	
	}
*/	

	// 1st strategy possible
 	displayStrategyNumber(big.NewInt(int64(1)))	
	displayFirstStrategy() // should match above
	
	
	// last strategy possible
	f := powBig(3, 180)
	g := f.Sub(f, big.NewInt(1))
	displayStrategyNumber(g)	
 	displayFinalStrategy() // should match above
 	
 	fmt.Printf("=============================\n")
 	
 	// Info link: https://stackoverflow.com/questions/54758130/how-to-obtain-the-amount-of-bits-of-a-bigint
 	//
 	
/*** 	
 	displayShiftedStrategyNumber(big.NewInt(int64(1)), 1)
 	displayShiftedStrategyNumber(big.NewInt(int64(1)), 2)
 	displayShiftedStrategyNumber(big.NewInt(int64(1)), 3)
 	displayShiftedStrategyNumber(big.NewInt(int64(1)), 4)
 	displayShiftedStrategyNumber(big.NewInt(int64(1)), 85)
 	displayShiftedStrategyNumber(big.NewInt(int64(1)), 86)
 	displayShiftedStrategyNumber(big.NewInt(int64(1)), 170)
 	displayShiftedStrategyNumber(big.NewInt(int64(1)), 180)
 	displayShiftedStrategyNumber(big.NewInt(int64(1)), 190)
 	
 	
 	// ^^^ the above is not going to work, I do not think.
 	//
***/ 	
 	
 	// why not take the largest possible number, and then subtract
 	// which number we want, or, no, add the largest
 	// possible number to the number we want, and then that will set the
 	// final bit on the flag.
 	//
 	// so , this is 3^180
 	//
 	// take 3^180, which is 7.617e+85, and add it to the number.
 	//
 	// for #1 we'll get whatever 7.617e+85 is.  
 	//
 	// For the last strategy we'll get it, plus all the bits underneath it.
 	
 	fmt.Printf("HERE WE GO ------------->\n")
 	
 	addTo := powBig(3, 180)
 	for i:=0; i<4; i++ {
	 	stratNum := big.NewInt(int64(i))
	 	
	 	h := stratNum.Add(stratNum, addTo)
	 	displayStrategyNumber(h)
 		
 	}
 	
 	
 	/*
 	displayStrategyNumber(big.NewInt(1))
 	displayStrategyNumber(big.NewInt(2))
 	displayStrategyNumber(big.NewInt(3))
 	displayStrategyNumber(big.NewInt(4))
 	displayStrategyNumber(big.NewInt(4))
 	*/

// 	testingStill2()
 	
 	os.Exit(0)
 	
		
	//res := powBig(3, 181) // what we may have to deal with as maximum # of possibilities
	res := powBig(3, 20) // start halfway
	fmt.Printf("result: %v\n", res)
	
	// Link: https://pkg.go.dev/math/big#Int.Lsh
	// and
	// https://stackoverflow.com/questions/30182129/calculating-large-exponentiation-in-golang
	//
	// bit shift 

/**

// okay, deal with this later - the idea is that we want to left shift this thing all the
// way over to max_value + 1, so that the resultant "string" will have leading values
// in order to fill in the entire matrix.
//
// We are going to use this resulting string to output the final strategy to a file
	
	shifted := res.Lsh(res, 2)
	fmt.Printf("shifted: %v\n", shifted)
**/
	
	// Convert to string
	str := res.String()
	fmt.Printf("result string: %s\n", str)
	
	// Use our base conversion thingie to convert into the 'pseudo string we want'
	s := convertDecimalToHand(str)
	fmt.Printf("Final: %s\n", s)
	
	rowInfo := sliceOutRowInfo(0, s)
	fmt.Printf("RowInfo: %s\n", rowInfo)
	
	
	
	
 }