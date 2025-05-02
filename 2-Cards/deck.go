package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	// cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) shuffle() {
	// seed variable to randomize the shuffle
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// Loop through the deck and swap each card with a random card
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) saveToFile(filename string) error {
	// Convert the deck to a byte slice
	data := []byte(d.toString())
	// Write the byte slice to a file
	return os.WriteFile(filename, data, 0666)
}

func newDeckFromFile(filename string) (deck) {
	// Read the byte slice from the file
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Print("Error: ",err)
		os.Exit(1)
	}
	// Convert the byte slice to a string and split it into a slice of strings
	return deck(strings.Split(string(data), ","))
}

func (d deck) toString() string {
	// Convert the deck to a string and join with a comma
	return strings.Join([]string(d), ",")
}