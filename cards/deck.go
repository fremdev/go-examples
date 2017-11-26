package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	suits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	ranks := []string{"Ace", "King", "Queen", "Jack", "Ten"}

	for _, suit := range suits {
		for _, rank := range ranks {
			cards = append(cards, rank+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0664)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error from newDeckFromFile:", err)
		os.Exit(1) // Exit from the program
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	// Random number generator depends on timestamp
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		// swap variables
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
