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
 * Versions:
 * v5.0 - add runtime statistics
 *         - # of strategies generated per second, output every 10 seconds
 *
 * v4.0 - minimize screen output, and give a final output to file option
 * v3.0 - reverse the string and output in correct order
 * (current) v2.0 - refining to output values as a matrix
 * v1.0 - initial concept
 *
 *
 * Next steps:
 *  output as a space separated string
 
 ****/
 
import(
	"os"
	"log"
	"fmt"
	"math/big"
	"strings"
	"time"
	bc "github.com/chtison/baseconverter"
)

var version = "2.0"




/* ----------- begin ripped code (do not make any changes) --------------- */


const NUM_COLS_SOFT = 10
const NUM_ROWS_SOFT = 9

var columnHeadingsSoft = [NUM_COLS_SOFT]string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "A"}
var rowHeadingsSoft = [NUM_ROWS_SOFT]string{"13", "14", "15", "16", "17", "18", "19", "20", "21"}

var softStrategyMatrix = [NUM_ROWS_SOFT][NUM_COLS_SOFT]string{
	{"S", "S", "S", "S", "S", "S", "S", "S", "S", "S"},
	{"S", "S", "S", "S", "S", "S", "S", "S", "S", "S"},
	{"S", "S", "S", "S", "S", "S", "S", "S", "S", "S"},
	{"S", "S", "S", "S", "S", "S", "S", "S", "S", "S"},
	{"S", "S", "S", "S", "S", "S", "S", "S", "S", "S"},
	{"S", "S", "S", "S", "S", "S", "S", "S", "S", "S"},
	{"S", "S", "S", "S", "S", "S", "S", "S", "S", "S"},
	{"S", "S", "S", "S", "S", "S", "S", "S", "S", "S"},
	{"S", "S", "S", "S", "S", "S", "S", "S", "S", "S"},
}

func appendSoftStrategyMatrixToFile(filename string) {
	
	fp, err := os.OpenFile(filename,  os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal("Cannot open: ", filename, " ", err)
	}
	
	defer fp.Close()
	
	// This is the hard hand strategy data
	_, err = fp.WriteString("\n\n[Soft]\n")
	if err != nil {
		log.Fatal("Cannot write to: ", filename, " ", err)
	}		
	
	// first column heading is 3-spaces over to right
	_, err = fp.WriteString("   ")
	if err != nil {
		log.Fatal("Cannot write to: ", filename, " ", err)
	}		
	
	// Spit out the column heading
	for col :=0; col <=(NUM_COLS_SOFT-1); col++ {
		text := fmt.Sprintf("%s ", columnHeadingsSoft[col])
		_, err = fp.WriteString(text)
			
		if err != nil {
			log.Fatal("Cannot write to: ", filename, " ", err)
		}		
	}
	
	// carriage return separates rows
	_, err = fp.WriteString("\n")
	if err != nil {
		log.Fatal("Cannot write to: ", filename, " ", err)
	}		
	
	
	// create the matrix	
	for row:=0; row<=(NUM_ROWS_SOFT-1); row++ {
		// Print the row heading
		// Make it look nice for headings with more than one character
		rh := rowHeadingsSoft[row]
		if len(rh) == 1 {
			//log.Printf("Single character row heading.")
			rh = fmt.Sprintf(" %s ", rh)
		} else {
			//log.Printf("Double character row heading.")
			rh = fmt.Sprintf("%s ", rh)
		}
		
		_, err = fp.WriteString(rh)
		if err != nil {
			log.Fatal("Cannot write to: ", filename, " ", err)
		}
	
		for col:=0; col<=(NUM_COLS_SOFT-1); col++ {
			// make it look nice for the one column with a 10-card in it.  add a leading space.
			var text string
			if len(columnHeadingsSoft[col]) > 1 {
				text = fmt.Sprintf(" %s ", softStrategyMatrix[row][col])
			} else {
				text = fmt.Sprintf("%s ", softStrategyMatrix[row][col])
			}
			
			// write the text to the file
			_, err = fp.WriteString(text)
				
			if err != nil {
				log.Fatal("Cannot write to: ", filename, " ", err)
			}
		}
		
		// carriage return separates rows
		_, err = fp.WriteString("\n")
		if err != nil {
			log.Fatal("Cannot write to: ", filename, " ", err)
		}		
	}
}



/* ----------- end ripped code (do not make any changes) --------------- */

// responsible *ultimately* for writing the [hard] strategy out
// to the file we want
func writeMetadataToFile(strategyNumberAsString string, filename string) {
	fmt.Printf("[singleStrategyGenerator.go][writeMetadataToFile][entry]\n")
	
	fp, err := os.OpenFile(filename,  os.O_APPEND|os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal("Cannot open: ", filename, " ", err)
	}
	
	defer fp.Close()


	_, err = fp.WriteString(fmt.Sprintf("#\n# SingleStrategy # %s\n", strategyNumberAsString))
	if err != nil {
		log.Fatal("Cannot write to: ", filename, " ", err)
	}
	
	currentTime := time.Now()
	/*
	nowStr := fmt.Sprintf("%d-%d-%d %d:%d:%d\n",
		currentTime.Year(),
		currentTime.Month(),
		currentTime.Day(),
		currentTime.Hour(),
		currentTime.Hour(),
		currentTime.Second())
	*/
				
	_, err = fp.WriteString(fmt.Sprintf("#\n# Copyright\n# @author dxb The Space Cowboy\n# @date %s\n#\n#\n\n", currentTime))
	if err != nil {
		log.Fatal("Cannot write to: ", filename, " ", err)
	}	

}



// responsible *ultimately* for writing the [hard] strategy out
// to the file we want
func writeMatrixToFile(matrix []string, filename string) {
	fmt.Printf("[singleStrategyGenerator.go][writeMatrixToFile][entry]\n")
	
	fp, err := os.OpenFile(filename,  os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal("Cannot open: ", filename, " ", err)
	}
	
	defer fp.Close()

	// loop through the matrix one character at a time
	// seems inefficient, to output a slice as a string to a file	
	for _, c := range matrix {
		fmt.Printf("%s",c)
		_, err = fp.WriteString(c)
		if err != nil {
			log.Fatal("Cannot write to: ", filename, " ", err)
		}	
	}
}







func convertDecimalToHand(number string) string {
	fmt.Printf("[convert_base.go][convertDecimalToHand()][entry]\n number=%s\n", number)
	
	// This is beginning to hit the idea - for blackjack
	//var number string = "20"
	var toBase string= "HSD"
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
func displayStrategyNumber(bigInt *big.Int) string {
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
	return s
} 



// outputs:
// '    2 3 4 5 6 7 8 9 10 A'
func outputColumnHeader(outputFile string) string {
	header := "[hard]\n   2 3 4 5 6 7 8 9 10 A\n"
	fmt.Printf("%s", header)
	
	if outputFile != "" {
		fmt.Printf("Write header to file: %s\n", outputFile)
		/*
		fp = file.open(outputFile, "w")
		fp.write(header)
		
		defer fp.close()
		*/
	}
	return header
}

// performs the finalization of the input string into a matrix
//
// pull from concepts: dev_learning_blocks/slice_of_string.go
//
//
// I know that I'm going to have a second revision of this, because
// in reality I need to reverse the string before outputting the
// final version of the matrix
//
//////
func formatStringAsMatrix(inputString string, nColumns int) {
	fmt.Printf("[singleStrategyGenerator.go][formatAsMatrix][entry]\n")
	fmt.Printf("inputString: %s\n", inputString)
	fmt.Printf("nColumns: %d\n", nColumns)
	
	// start at the end
	// chop off nColumns at a time from the end
	strLen := len(inputString)
	for j:=strLen; j>0; j-=nColumns {
		//fmt.Printf("j: %d\n", j)
		
		// first iteration i want [171:181]
		// second iteration i want [161:171]
		//
		// save first index and use in next index ? 
		// no.  
		
		// for us, if we see that 'j' == 1, then we know we are
		// done outputting the matrix, because there will always
		// be an extra character at the end due to the way we
		// are computing the string to output (by always adding
		// one larger than the maximum number of possibilities
		// 
		if j == 1 {
			break
		}
		
		// chop off the next n characters from the string
		// do some sanity checking
		firstIndex := j-nColumns
		if firstIndex < 0 {
			//fmt.Printf("What is happening? string too short! truncate.\n")
			firstIndex = 0
		}

		subStr := inputString[firstIndex:j]
		fmt.Printf("[%d]\tsubStr: %s\n", j, subStr)
	}
}


// performs the finalization of the input string into a matrix
//
// pull from concepts: dev_learning_blocks/slice_of_string.go
//
//
// I know that I'm going to have a second revision of this, because
// in reality I need to reverse the string before outputting the
// final version of the matrix
//
//////
func printStrategyAsMatrix(inputString string, nColumns int) []string {
	fmt.Printf("[singleStrategyGenerator.go][printStrategyAsMatrix][entry]\n")
	fmt.Printf("inputString: %s\n", inputString)
	fmt.Printf("nColumns: %d\n", nColumns)
	
	header := outputColumnHeader("")
	
	
	// store the response, so we can return it.  The caller can then
	// do what they want (ehm, write to file?)
	//
	response := make([]string, 0)
	response = append(response, header)

	// start at '4'
	// '3' is covered by a soft-hand, ace+2 is only waay to get a '3'
	// idk what covers a hand of '2' , since the program does not (yet)
	// handle properly splitting a hand.
	//
	// @TODO - will have to revisit (nov '23)
	//	
	rowValue := 4
	
	// start at the end
	// chop off nColumns at a time from the end
	strLen := len(inputString)
	for j:=strLen; j>0; j-=nColumns {
		//fmt.Printf("j: %d\n", j)
		
		// first iteration i want [171:181]
		// second iteration i want [161:171]
		//
		// save first index and use in next index ? 
		// no.  
		
		// for us, if we see that 'j' == 1, then we know we are
		// done outputting the matrix, because there will always
		// be an extra character at the end due to the way we
		// are computing the string to output (by always adding
		// one larger than the maximum number of possibilities
		// 
		if j == 1 {
			break
		}
		
		// chop off the next n characters from the string
		// do some sanity checking
		firstIndex := j-nColumns
		if firstIndex < 0 {
			//fmt.Printf("What is happening? string too short! truncate.\n")
			firstIndex = 0
		}

		subStr := inputString[firstIndex:j]
		// dev, take out
		//fmt.Printf("[%d]\tsubStr: %s\n", j, subStr)
		
		// output the column heading (what row # we are on)
		if rowValue < 10 {
			// values less than ten are only one space
			//fmt.Printf("%d  %s\n", rowValue, subStr)
			fmt.Printf(" %d ", rowValue)
			response = append(response, fmt.Sprintf(" %d ", rowValue))
			
			/*
			strSlice := strings.Split(subStr, "")
			//fmt.Printf("strSlice: %s\n", strSlice)
			

			response = append(response, fmt.Sprintf("%d   ", rowValue))
			for k, c := range strSlice {
				// to explain - the column with a header of '10' needs to be lined up visually
				// It's all about the visual!
				if k == 8 {
					fmt.Printf(" %s ", c)
					response = append(response, fmt.Sprintf(" %s ", c))
				} else {
					fmt.Printf("%s ", c)	
					response = append(response, fmt.Sprintf("%s ", c))
				}
				
			}
			
			fmt.Printf("\n")
			response = append(response, fmt.Sprintf("\n"))
			//fmt.Printf("\n- row done -\n")
			*/
		} else {
			// 10 and larger take up two spaces
			//fmt.Printf("%d %s\n", rowValue, subStr)

			fmt.Printf("%d ", rowValue)
			response = append(response, fmt.Sprintf("%d ", rowValue))

			/*
			strSlice := strings.Split(subStr, "")
			//fmt.Printf("strSlice: %s\n", strSlice)
			
			for k, c := range strSlice {
				// to explain - the column with a header of '10' needs to be lined up visually
				// It's all about the visual!
				if k == 8 {
					fmt.Printf(" %s ", c)
				} else {
					fmt.Printf("%s ", c)	
				}
				
			}
			
			fmt.Printf("\n")
			//fmt.Printf("\n- row done -\n")
			*/
		}
		
		strSlice := strings.Split(subStr, "")
		for k, c := range strSlice {
			// to explain - the column with a header of '10' needs to be lined up visually
			// It's all about the visual!
			if k == 8 {
				fmt.Printf(" %s ", c)
				response = append(response, fmt.Sprintf(" %s ", c))
			} else {
				fmt.Printf("%s ", c)	
				response = append(response, fmt.Sprintf("%s ", c))
			}
			
		}
		
		fmt.Printf("\n")
		response = append(response, fmt.Sprintf("\n"))
		//fmt.Printf("\n- row done -\n")
		
				
		rowValue++
	}
	
	return response
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
 	
 	basename := "bi_singlestrat"
 	 	
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

 		filename := fmt.Sprintf("%s_%s.txt", basename, stratNum.String())
	 	
		writeMetadataToFile(stratNum.String(), filename)

		fmt.Printf("Stratnum before: %v\n", stratNum)	 	
	 	h := stratNum.Add(stratNum, addTo)
	 	
	 	fmt.Printf("h after: %v\n", h)
	 	fmt.Printf("Stratnum after: %v\n", stratNum)
	 	
	 	asString := displayStrategyNumber(h)
	 	
	 	// 10-columns
 		//formatStringAsMatrix(asString, 10)
 		matrix := printStrategyAsMatrix(asString, 10)
 		fmt.Printf("OUTPUT MATRIX:\n%v\n", matrix)
 		fmt.Printf("--- reformatted:\n")
 		for _, c := range matrix {
 			fmt.Printf("%s",c)
 		}

 		writeMatrixToFile(matrix, filename)
 		
 		
 		
 		// ripped from strategyGenerator_main program
 		appendSoftStrategyMatrixToFile(filename)
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