package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades" // := only for new variables
	// cards := deck{newCard(), newCard()}
	// cards = append(cards, "Six of Spades") // Appends a value at the end of the slice, does not modify the used create a new slice
	cards := newDeck()
	fmt.Println("All deck")
	cards.print() //Any variable of type deck now have access to print method
	hand, remainingCards := deal(cards, 5)
	fmt.Println("hand 1")
	hand.print()
	fmt.Println("remaining cards")
	remainingCards.print()

}

// defining functions

// func newCard() string {
// 	return "Five of Diamonds"
// }

// Array is a fixed length list of things
// Slice is an array that can grow or shrink
// Both can support only one data type
