package main

import "fmt"

func main() {
	// TODO: create a map called "modules"
	modules := map[int]string{
		104:"Datenmodell implementieren",
		117:"Netzinfrastruktur",
		346:"Cloud Lösungen konzipieren und realisieren",
	}
	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	delete(modules, 346)
	
	// TODO: delete one
	// TODO: add one
	// TODO: replace one
	fmt.Println(modules)
}
