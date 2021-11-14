package main

import "fmt"

func main() {
	cards := []string{"Ace of Diamonds", newCard()}
	// append creates a new slice and returns it
	cards = append(cards, "Six of Spades")

	// using := because we are redeclaring variable for each item inside slice
	for i, card := range cards {
		fmt.Println(i, card)
	}

	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}