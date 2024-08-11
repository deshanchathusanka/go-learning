package main

import "fmt"

// define deck type
type deck []string

func buildDeck() deck {
	cards := deck{}
	shapes := []string{"Diamonds", "Spades", "Hearts", "Clubs"}
	values := []string{"Ace", "Two", "Three", "Four"}

	for _, shape := range shapes {
		for _, value := range values {
			cards = append(cards, value+" of "+shape)
		}
	}

	return cards
}

// define receiver to print cards in the deck instance
func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}

func deal(d deck, n int) (deck, deck) {
	return d[:n], d[n:]
}
