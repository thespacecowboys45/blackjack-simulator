package main
/**
 * @date Oct '23
 * @author dxb The Space Cowboy
 *
 ***
 *
 * DESCRIPTION:
 *
 * Generates all possible strategies in a matrix and outputs as a file
 *
 *****/
 
import(
	"fmt"
	"log"
	"os"
	)

func init() {
	log.Printf("[init][entry]")
}

var iterator int = 0

func exp(base int, multiplier int) int {

	var result int = 1
	
	log.Printf("[exp][entry]")
	log.Printf("base: %d multiplier: %d iterator: %d", base, multiplier, iterator)

	iterator++
	
	// safety valve while developing
	if iterator > 5 {
		log.Fatal("Too many iterations.")
	}
	
	if multiplier != 1 {
		// we are done
		result = base * exp(base, multiplier-1)
	}	
	
	return result
}


/**
 * DEV
 
const NUM_COLS = 3
//const NUM_COLS = 10

const NUM_ROWS = 2
//const NUM_ROWS = 3
//const NUM_ROWS = 4
const NUM_POTENTIAL_VALUES = 2

// THE MATRIX!
var strategyMatrix[NUM_ROWS][NUM_COLS] string

var columnHeadings = [NUM_COLS]string{"2", "13", "4"}
//var columnHeadings = [NUM_COLS]string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "A"}

var rowHeadings = [NUM_ROWS]string{"4", "15"}
//var rowHeadings = [NUM_ROWS]string{"4", "5", "6"}
//var rowHeadings = [NUM_ROWS]string{"4", "5", "6", "7"}

var valuesArray = [NUM_POTENTIAL_VALUES]string{"S", "H"}

const WRITEFILE_ITERATOR=10 

 DEV
 */
 

/**
 * PROD
 */
const NUM_COLS = 10
const NUM_ROWS = 18
const NUM_POTENTIAL_VALUES = 2

// THE MATRIX!
var strategyMatrix[NUM_ROWS][NUM_COLS] string

var columnHeadings = [NUM_COLS]string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "A"}
var rowHeadings = [NUM_ROWS]string{"4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21"}

var valuesArray = [NUM_POTENTIAL_VALUES]string{"S", "H"}

const WRITEFILE_ITERATOR=100000
 /*
 PROD
 */




func createStrategyMatrix(row_pos int, col_pos int, currVal string) string {
	log.Printf("[createStrategyMatrix][entry]")
	var result string
	//valuesArray := [3]string{"a", "b", "c"} 
	
	for _, v := range(valuesArray) {
		log.Printf("Add value: %s", v)
		if result == "" {
			result = fmt.Sprintf("%s", v)
		} else {
			result = fmt.Sprintf("%s,%s", result, v)
		}
	}
/*	
	valPos := findValuePosition(currVal)
	
	if valPos <= len(valuesArray) {
		nextVal := valuesArray[valPos]
		result = createStrategyMatrix(row_pos, col_pos, nextVal)
	}
*/

	return result
}



/**
 * increments the value at the specified position
 * returns if we are on the lastValue in the values array
 */
func incValAtPosition(row int, col int) bool {
	log.Printf("[incValAtPosition][entry]")
	log.Printf("request for row=%d col=%d", row, col)

	//valuesArray := [3]string{"a", "b", "c"} 
	log.Printf("valuesArray length: %d", len(valuesArray))
	
	currentValue := strategyMatrix[row][col]
	valPosition := 0
	for k, val := range valuesArray {
		if currentValue == val {
			log.Printf("Found matching value: %s at pos %d / %d", val, valPosition, k)
			break
		}
		valPosition++
	}
	
	// return value of len() is 1-based
	log.Printf("Compare: %d to %d", valPosition, (len(valuesArray)-1))
	if valPosition < (len(valuesArray)-1) {
		log.Printf("Not on last value")
		// set to the next value in the array of possible values
		strategyMatrix[row][col] = valuesArray[valPosition+1]

		// not on the last value, so return false
		return false
	}
	log.Printf("We are on the last value")
	
	// we are already using the last possible value
	// so, reset and start over at 0
	strategyMatrix[row][col] = valuesArray[0]
	
	// we are on the last value, so return true 
	return true
}




func initMatrix() [NUM_ROWS][NUM_COLS]string {
	log.Printf("[initMatrix][entry]")
	for i:=0; i<NUM_ROWS; i++ {
		for j:=0; j<NUM_COLS; j++ {
			log.Printf("%d,%d", i, j)
			strategyMatrix[i][j] = valuesArray[0]
		}
	}
	return strategyMatrix
}


func incVal(row int, col int) bool {
	//log.Printf("[incVal][entry]")
	//valuesArray := [3]string{"a", "b", "c"} 

	currentValue := strategyMatrix[row][col]
	valPosition := 0
	for _, val := range valuesArray {
		if currentValue == val {
			//log.Printf("[%d][%d]Found matching value: %s at pos %d / %d", row, col, val, valPosition, k)
			break
		}
		valPosition++
	}

	//log.Printf("[%d][%d] Compare: %d to %d", row,col, valPosition, (len(valuesArray)-1))
	if valPosition >= (len(valuesArray)-1) {
		//log.Printf("cannot increment")
		// cannot increment
		
		// First, see if we have any more columns left after this one
		if col >= (NUM_COLS-1) {
			//log.Printf("DONE with this column - try to move to next row")
			
			// reset current item to the first possible value
			strategyMatrix[row][col] = valuesArray[0]
			
			// start over in first column
			col = 0
			
			//log.Printf("Compare row values: %d to %d", row, (NUM_ROWS-1))
			if row >= (NUM_ROWS-1) {
				log.Printf("DONE with all rows - final inner loop point reached.")
				// DONE!!!!! -> this is the innermost loop
				return false
			}
			
			return incVal(row+1, col)			
		}
		
		// Else, we're not done, we have more columns to affect.  Continue.
		
		// reset current item to the first possible value
		strategyMatrix[row][col] = valuesArray[0]
		
		// increment the value of the next column over instead
		return incVal(row, col+1)
	}
	//log.Printf("we can increment")
	
	nextVal := valuesArray[valPosition+1]
	strategyMatrix[row][col] = nextVal
	// we incremented successfully
	return true
}

func writeMatrixToFile(filename string) {
	//fp, err := os.OpenFile(filename,  os.O_APPEND|os.O_WRONLY, 0600)
	fp, err := os.Create(filename)
	if err != nil {
		log.Fatal("Cannot open: ", filename, " ", err)
	}
	
	defer fp.Close()

	// This is the hard hand strategy data
	_, err = fp.WriteString("[Hard]\n")
	if err != nil {
		log.Fatal("Cannot write to: ", filename, " ", err)
	}		
	
	// first column heading is 3-spaces over to right
	_, err = fp.WriteString("   ")
	if err != nil {
		log.Fatal("Cannot write to: ", filename, " ", err)
	}		
	
	// Spit out the column heading
	for col :=0; col <=(NUM_COLS-1); col++ {
		text := fmt.Sprintf("%s ", columnHeadings[col])
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
	for row:=0; row<=(NUM_ROWS-1); row++ {
		// Print the row heading
		// Make it look nice for headings with more than one character
		rh := rowHeadings[row]
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
	
		for col:=0; col<=(NUM_COLS-1); col++ {
			// make it look nice for the one column with a 10-card in it.  add a leading space.
			var text string
			if len(columnHeadings[col]) > 1 {
				text = fmt.Sprintf(" %s ", strategyMatrix[row][col])
			} else {
				text = fmt.Sprintf("%s ", strategyMatrix[row][col])
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


func main2() {
	log.Printf("[main2][entry]")
	count := 0
	
	var strategyFileDir = "output"
	var strategyFileBasename = "autostrat"
	
	for incVal(0, 0) != false {
		log.Printf("[%d] Matrix NOW: %v", count, strategyMatrix)
		//log.Printf("iterate count: %d", count)
		
		// every so often write a file
		if count % WRITEFILE_ITERATOR == 0 {
			log.Printf("[%d] Qualifies for strategy test.", count)
	
			filename := fmt.Sprintf("%s/%s_%d.txt", strategyFileDir, strategyFileBasename, count)
			log.Printf("[%d] Qualifies for strategy test. Write: %s", count, filename)

			writeMatrixToFile(filename)
		}
		count++
	}
	log.Printf("Count of iterations: %d", count)
	
	// Check if any more rows to go through
	
	// At this point - we are done
	log.Printf("Matrix FINAL: %v", strategyMatrix)
	writeMatrixToFile("autostrat1.txt")
}


func main() {
	log.Printf("[main][entry]")
	
	matrixToPrint := initMatrix()
	log.Printf("matrixToPrint INITIALIZED: %v", matrixToPrint) 
		
	main2()
	log.Fatal("quit")
	
	base := 2
	exponent := 4
	
	total := exp(base, exponent)
	log.Printf("Total of %d to the %d power is %d", 2, 4, total)
	
	matrixToPrint = initMatrix()
	log.Printf("matrixToPrint 1: %v", matrixToPrint) 


	row_incrementor := 0
	col_incrementor := 0

	log.Printf("strategyMatrix INITIAL: %v", strategyMatrix)
	rc := incValAtPosition(row_incrementor, col_incrementor)
	log.Printf("strategyMatrix FIRST: %v", strategyMatrix)
	
	
	for col_incrementor < (NUM_COLS-1) {
		log.Printf("Top of loop: row=%d, col=%d", row_incrementor, col_incrementor)
		for rc != true {
			log.Printf("Request to increment")
			rc = incValAtPosition(row_incrementor, col_incrementor)
			log.Printf("strategyMatrix AFTER: %v", strategyMatrix)
		}
		col_incrementor++
		
		rc = incValAtPosition(row_incrementor, col_incrementor)
		log.Printf("strategyMatrix at bottom of loop: %v", strategyMatrix)
	}
	



	
//	matrixToPrint = createStrategyMatrix()
//	log.Printf("matrixToPrint 2: %v", matrixToPrint) 

}