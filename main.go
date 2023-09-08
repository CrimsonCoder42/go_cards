package main

import (
	"fmt"
)


func main() {
	
	cards := []string{
		newCard(),
		"Two of Diamonds",
	}

	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}

}

func newCard() string {
	return "Five of Diamonds"
}