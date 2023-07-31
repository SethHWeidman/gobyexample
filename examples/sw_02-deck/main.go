package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	cards := newDeck()
	filename := "my_deck_file"
	cards.saveToFile(filename)

	cards2 := newDeckFromFile(filename)
	cards2.print()

	// tests!
	d := newDeck()

	if len(d) != 16 {
		fmt.Printf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		fmt.Printf("Expected first card of Ace of Spades, but got %v", d[0])
	}
}

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
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

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func (d deck) saveToFile(filename string) error {
	// https://pkg.go.dev/os#WriteFile
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	byte_slice, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(byte_slice), ",")
	return deck(s)
}
