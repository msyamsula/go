package main

func newCard() string {
	return "jack of hearts"
}

func main() {

	// cards := deck{"ace of diamonds", "king of hearts", "queen of club", newCard()}
	playingDeck := newDeck()
	// hand := deck{}

	// shuffle
	playingDeck = playingDeck.shuffle()
	playingDeck.print()

	// deal function
	// hand, playingDeck = playingDeck.deal(10)

	// playingDeck.print()
	// hand.print()

	// // save function
	// hand.save("hand.txt")

	// // read from file
	// fileDeck := readDeck("hand.txt")
	// fmt.Println("from file: ")
	// fileDeck.print()

}
