package main

import "fmt"

var firstName, lastName string = "Timon", "Soom"
var dayOfBirth, monthOfBirth, yearOfBirth, numberOfSiblings int = 21, 03, 2008, 3
var heightInMeters float32 = 1.85
var zodiacSign rune = '\u2648'

func main() {
	// TODO: Declare and initialize the variables being used in the output!
	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
