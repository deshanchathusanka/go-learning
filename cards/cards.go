package main

import "fmt"

func main() {
	cards := buildDeck()
	fmt.Println("Original Deck Print")
	cards.print()
	fmt.Println()

	d1, d2 := deal(cards, 4)
	fmt.Println("first Deal Print")
	d1.print()
	fmt.Println()

	fmt.Println("Last Deal Print")
	d2.print()

}
