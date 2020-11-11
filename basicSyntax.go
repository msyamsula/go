package main

import "fmt"

// function declaration
func newCard() string {
	return "Ace of Spades"
}

func main2() {
	// variable declaration
	// colon (:) is used for initialization
	// var card string = "Ace of Spades" // explicit variable declaration
	// card := "Ace of Spades" // delegate type definition to go language

	// mycard := newCard()
	// fmt.Println(mycard)

	// array declaration
	cards := []string{"Ace of Diamond", newCard()}
	cards = append(cards, "Six of Spades")

	// fancy looping
	for i, card := range cards {
		fmt.Println(i, card)
	}
	// index looping
	for i := 0; i < 3; i++ {
		fmt.Println(cards[i])
	}
	// while loop
	i := 0
	for i < 3 {
		fmt.Println(cards[i])
		i++
	}

}
