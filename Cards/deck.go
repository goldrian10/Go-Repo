package main

import "fmt"

// Create a new type of deck
// which is a slice of strings
type deck []string //deck borrow all the behavior of the slice strings

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardNumber := []string{"Ace", "Two", "Three", "Four"} // , "Five", "Six", "Seven", "Eight", "Nine", "Ten", "J", "Q", "K"}
	for _, suit := range cardSuits {                      // _ is a variable that we will not gonna use
		for _, number := range cardNumber {
			cards = append(cards, suit+" of "+number)
		}
	}
	return cards
}

func (d deck) print() { // (d deck) is called receiver
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// slice[startIndexIncluding : upToNotIncluded]
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
