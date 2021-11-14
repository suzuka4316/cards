package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// receiver function => deck.print()
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// pass deck as parameter => deal(deck, 3)
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// deck.toString()
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// deck.saveToFile(filename)
func (d deck) saveToFile(filename string) error { // type error
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666) // anyone can read and write file
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename) //byte slice, error
	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - log the error and entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}