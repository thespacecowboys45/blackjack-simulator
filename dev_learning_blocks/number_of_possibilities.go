package main
/*****
 *
 * @date November '23
 * @author dxb The Space Cowboy
 *
 * DESCRIPTION:
 *    A program to spit out the number represented by 3e180
 *
 *****/
import (
	"fmt"
	"math"
)

func calculate() {
	fmt.Printf("[calculate()][entry]\n")
    res := math.Pow(3, 20) // 3^180
    fmt.Println(res)
}	
	

func main () {
	fmt.Printf("[number_of_possibilities.go][entry]\n")
	calculate()
	
}