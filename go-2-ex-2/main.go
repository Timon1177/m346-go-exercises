package main

import "fmt"

func main() {
	var fibs = []int{1, 1}
	
	for i:=0; i<10;i++{
	fibs = append(fibs, 0)
	fibs[i+2] = fibs[i] + fibs[i+1]
	// TODO: correct up to index 4 using direct element access
	
	 // TODO: replace 0 with the next Fibonacci number
	// TODO: compute three more Fibonacci numbers and append them
	
	}
	fmt.Println(fibs) // expected output: [1 1 2 3 5 8 13 21 34]
}
