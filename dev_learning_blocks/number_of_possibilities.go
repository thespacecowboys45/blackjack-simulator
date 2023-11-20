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
	"math/big"
	"reflect"
	"strconv"
)

func calculate() {
	fmt.Printf("[calculate()][entry]\n")
    res := math.Pow(3, 180) // 3^180
    fmt.Printf("Type: %T\n", res)
    fmt.Println(res)
    
    
	bigNum := "1267650600228229401496703205376"
	b, ok := big.NewInt(0).SetString(bigNum, 10)
	fmt.Println(ok, b)
	// true 1267650600228229401496703205376    

	var f float64 = 3.1415926535
	fmt.Println(reflect.TypeOf(f))
	fmt.Println(f)

	// https://www.golangprograms.com/how-to-convert-float-to-string-type-in-go.html
	var s string = strconv.FormatFloat(f, 'E', -1, 32)
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(s)
	
	c := 12.454
	fmt.Println(reflect.TypeOf(c))

	s = fmt.Sprintf("%v", c)
	fmt.Println(s)
	fmt.Println(reflect.TypeOf(s))	
}	
	

func main () {
	fmt.Printf("[number_of_possibilities.go][entry]\n")
	calculate()
	
}