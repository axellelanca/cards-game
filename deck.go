package main

import "fmt"

// Create a new type of 'deck'
// wich is a slice of strings
type deck []string

// function that creates and returns a list of playing cards
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

// Loop through the deck of card and print the cards
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// The function returns two values 
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}