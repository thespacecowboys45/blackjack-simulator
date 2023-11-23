package main

import(
	"fmt"
)

func main() {
	
	var input = "0123456789abcdefghijklmnopqrstuv"
	
	strLen := len(input)
	
	index1 := 0
	index2 := strLen
	
	fmt.Printf("String: %s len: %d\n", input, strLen)
	
	// slice out a part of the string from index1 to index2 (exclusive)
	// index2 (above) just happens to be 1-based, it's the string's length!
	// using array notation in golang it is going to slice out all the way
	// to (index2 - 1), and give back that portion of the string
	//
	// IF we happen to try and slice out something larger than the actual
	// string's length, then Golang will throw an exception with 
	// slice bounds out of range
	//
	fmt.Printf("substr: %s\n", string(input[index1:index2]))
}