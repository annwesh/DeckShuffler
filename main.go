package main

//gives package main all the access and functionality of the package fmt
import (
	"fmt"
)

func main() {
	fmt.Println("Hellow Worle")
    
   
	cards := newDeck()
	cards.print()
	cards, _ = deal(cards, 5)
	cards.print()
	cards.saveToFile("my_cards")
    //fmt.Println([]byte(testCheck[1]))
    cards = readDeckFromFile("my_cards")


    cards.print()
    cards.shuffleDeck()
    cards.print()

}

func newSlice (length int) (string, int) {
	return "new Slice", 5
}