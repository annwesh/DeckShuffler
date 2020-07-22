package main

import (
	"fmt"
	"strconv"
	"strings"
	"io/ioutil"
	"log"
	"os"
	"math/rand"
	"time"
)

type deck []string

//return a new deck of cards
func newDeck () deck {
	deckOfCards :=deck{}

	cardHouses := []string{"Spade", "Diamond", "Heart", "Club"}
	cardNumbers := []int{1,2,3,4}

	for _,cardHouse := range cardHouses {
		for _, cardNumber := range cardNumbers {
			deckOfCards = append(deckOfCards, cardHouse+strconv.Itoa(cardNumber))
		}
	}

	return deckOfCards
}

//receiver function to print the complete deck
func (d deck) print() {
	for i, data := range d {
		fmt.Println(i,data)
	}
}

// receiver function to cut a deal of cards , and returns the remaining cards
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//comma spearated 
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}


//save the deck to file
func (d deck) saveToFile(fileName string) {
   ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

//read deck from file
func readDeckFromFile(fileName string) deck {
	content, err := ioutil.ReadFile(fileName)
	if err!=nil {
		log.Fatal("Error occured while reading from the file :", err)
		os.Exit(0)
	}
	return deck(strings.Split(string(content), ",")) 
}

//shuffle logic of the current deck
func (d deck) shuffleDeck()  {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i,_ := range d {
		index := r.Intn(len(d))
		d[i], d[index] = d[index], d[i]
	}
}