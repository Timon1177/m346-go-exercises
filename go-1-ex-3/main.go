package main

import (
	"fmt"
	"math/rand"
	"time"
	"os"
)

func main() {
	var eyes = rand.Intn(5) + 1
	var when = time.Now()

	// TODO: use fmt.Fprintln instead!
	fmt.Fprintln(os.Stdout, "the dice shows", eyes, "eyes")
	
	// TODO: use fmt.Fprintln instead!
	fmt.Fprintln(os.Stdout, "the dice was rolled at", when)
	
	// TODO: how to write the output into eyes.txt and dice.log?
	// go run ex3/main.go TODO

}
