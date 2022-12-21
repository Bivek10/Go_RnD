package main

import (
	"fmt"
)

type Country struct {
	name        string
	poupulation int
}

func main() {
	countryData := Country{"Africa", 100000}
	fmt.Println(countryData)
	fmt.Printf("Countryname %v \n", countryData.name)
	fmt.Printf("Population %v \n", countryData.poupulation)

}
