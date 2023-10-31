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


const NUM_COLS = 3
const NUM_ROWS = 2

var strategyMatrix[NUM_ROWS][NUM_COLS] string




func createStrategyMatrix(row_pos int, col_pos int, currVal string) string {
	log.Printf("[createStrategyMatrix][entry]")
	var result string
	valuesArray := [3]string{"a", "b", "c"} 
	
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

	valuesArray := [3]string{"a", "b", "c"} 
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
			strategyMatrix[i][j] = "a"
		}
	}
	return strategyMatrix
}


func incVal(row int, col int) bool {
	log.Printf("[incVal][entry]")
	valuesArray := [3]string{"a", "b", "c"} 

	currentValue := strategyMatrix[row][col]
	valPosition := 0
	for k, val := range valuesArray {
		if currentValue == val {
			log.Printf("[%d][%d]Found matching value: %s at pos %d / %d", row, col, val, valPosition, k)
			break
		}
		valPosition++
	}

	log.Printf("[%d][%d] Compare: %d to %d", row,col, valPosition, (len(valuesArray)-1))
	if valPosition >= (len(valuesArray)-1) {
		log.Printf("cannot increment")
		// cannot increment
		
		// First, see if we have any more columns left after this one
		if col >= (NUM_COLS-1) {
			log.Printf("DONE with this column - try to move to next row")
			
			// reset current item to the first possible value
			strategyMatrix[row][col] = valuesArray[0]
			
			// start over in first column
			col = 0
			
			log.Printf("Compare row values: %d to %d", row, (NUM_ROWS-1))
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
	log.Printf("we can increment")
	
	nextVal := valuesArray[valPosition+1]
	strategyMatrix[row][col] = nextVal
	// we incremented successfully
	return true
}


func main2() {
	log.Printf("[main2][entry]")
	count := 0
	for incVal(0, 0) != false {
		log.Printf("Matrix NOW: %v", strategyMatrix)
		log.Printf("iterate count: %d", count)
		count++
	}
	
	// Check if any more rows to go through
	
	// At this point - we are done
	log.Printf("Matrix FINAL: %v", strategyMatrix)
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