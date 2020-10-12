package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// new type of 'deck', it is equivalent with object in c++
type deck []string

// receiver function (object method in c++)
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// multiple return
func (d deck) deal(num int) (deck, deck) {
	dealed := d[:num]
	left := d[num:len(d)]

	return dealed, left
}

func (d deck) toString() string {
	d = []string(d)
	ans := ""
	sep := ", "

	// manual join
	// for i, s := range d {
	// 	if i != 0 {
	// 		ans += sep
	// 	}
	// 	ans += s
	// }

	// fancy (STL) from "strings" package
	ans = strings.Join(d, sep)

	return ans
}

func (d deck) save(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func (d deck) shuffle() deck {
	// random number with seed in go, boilerplate
	// copas it, no need to think
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	r := rand.New(source)

	for i := 0; i < len(d); i++ {
		randomNumber := r.Intn(len(d)) // random number in go

		// old school swap
		// temp := d[i]
		// d[i] = d[randomNumber]
		// d[randomNumber] = temp

		// fancy swap
		d[i], d[randomNumber] = d[randomNumber], d[i]
	}

	return d
}

// generate a deck, be aware it is not receiver (not deck method)
func newDeck() deck {
	cardSuit := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValue := []string{"ace", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "jack", "queen", "king"}

	cards := deck{}
	for i := 0; i < len(cardSuit); i++ {
		for j := 0; j < len(cardValue); j++ {
			currentCard := cardValue[j] + " of " + cardSuit[i]
			cards = append(cards, currentCard)
		}
	}

	return cards
}

// read deck from file
func readDeck(filename string) deck {
	// if something wrong while reading a file err will not null
	// success mean err = nil
	byteSlice, err := ioutil.ReadFile(filename)
	if err != nil {
		// log the error and exit
		fmt.Println("error when reading deck from file")
		os.Exit(1)
	}

	stringDeck := string(byteSlice)
	mydeck := strings.Split(stringDeck, ", ")

	return mydeck
}
