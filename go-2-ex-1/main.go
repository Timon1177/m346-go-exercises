package main

import "fmt"

type FullName struct {
	FirstName string
	LastName string
}

type BirthDate struct {
	DayOfBirth byte
	MonthOfBirth byte
	YearOfBirth int16
}

type Profile struct {
	Name FullName
	Born BirthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}
var myName FullName= FullName{
	FirstName: "Timon",
	LastName: "Soom",
}
var born BirthDate= BirthDate{
	DayOfBirth: 21,
	MonthOfBirth: 03,
	YearOfBirth: 2008,
}

func main() {
	var me = Profile{
		Name: myName,
		Born: born,
		NumberOfSiblings: 3,   // TODO: adjust
		ZodiacSign: '\u2648', // TODO: adjust
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	me.NumberOfSiblings++;
	// TODO: imagine, you get a little brother or sister
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
