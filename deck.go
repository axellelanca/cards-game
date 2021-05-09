package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

// To convert a deck to a string
func (d deck) toString() string {
	//We convert the deck in a slice of string
	// To join the slices all together in one string, we use another package called strings
	return strings.Join([]string(d), ",")
}

// To save a deck in a local file
// 0666 permission : anyone can read and modify a file
func (d deck) saveToFile(filename string ) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}


func newDeckFromFile(filename string) deck {
	// ReadFile function can return two values
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option #1 - Log the error and return a call to newDeck()
		// Option #2 - Log the error and entirely quit the program:
		// we're going to use another package called "os"
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	//We convert the bytes slice into a string, and then we split it to have an array
	s := strings.Split(string(bs), ",")

	return deck(s)
}


func (d deck) shuffle() {

	//To generate a random number 
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) -1)
		// Swap the elements inside the Slice
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

 