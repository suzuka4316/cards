package main

import "fmt"

// Create new type of 'deck' which is a slice of strings
type deck []string

// receiver function: any vars of type deck have access to print() function
// d is a reference of actual variable that'll be used for this function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}