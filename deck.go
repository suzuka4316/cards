package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

func (d deck) shuffle() {
	// create a new seed value to make the deck return value to be truly random (by default this seed value will be always the same which makes not so random)
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1) // random number between 0 to (5-1)

		d[i], d[newPosition] = d[newPosition], d[i] // swapping elements
	}
}