package main

func main() {
	// cards := []string{"Ace of Diamonds", newCard()}
	// cards := deck{"Ace of Diamonds", newCard()}

	//// newDeck
	cards := newDeck()

	// cards = append(cards, "Six of Spades")

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	// cards.print()

	//// deal
	// hand, remainingCard := deal(cards, 5)
	// hand.print()
	// remainingCard.print()

	// fmt.Println(cards.toString())

	//// saveToFile
	// cards.saveToFile("my_cards.txt")
	
	//// newDeckFromFile
	// cards := newDeckFromFile("my_cards.txt")
	// cards.print()


	//// shuffle
	cards.shuffle()
	cards.print()
}