package main

import "fmt"

func DealAPackOfCards(deck []int) {
	// Each player gets 3 cards (12 cards / 4 players = 3 cards per player)
	cardsPerPlayer := 3
	for i := 0; i < len(deck); i += cardsPerPlayer {
		playerNumber := i/cardsPerPlayer + 1
		playerCards := deck[i : i+cardsPerPlayer]
		fmt.Printf("Player %d: %v\n", playerNumber, playerCards)
	}
}

func main() {
	deck := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	DealAPackOfCards(deck)
}
