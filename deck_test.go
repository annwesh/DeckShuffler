package main 

import "testing"
import "log"


func TestNewDeck (t *testing.T) {
	deck := newDeck()

	if len(deck) != 16 {
		log.Fatal( "Numner not equa to the expected deck size")
		t.Errorf("Exit as size expected 16 : got %v", len(deck))
	}
}