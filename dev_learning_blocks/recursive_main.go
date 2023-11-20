package main
/**
 * @author dxb The Space Cowboy David Boardman
 * @date Nov '23
 *
 * An example of a recursive function
 */
import(
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

func main() {
	log.Printf("[main][entry]")
	
	base := 2
	exponent := 4
	
	total := exp(base, exponent)
	log.Printf("Total of %d to the %d power is %d", 2, 4, total)
}