package player

import "github.com/linesmerrill/blackjackSolver/deck"

// Action of a player or dealer
type Action interface {
	Hit(c deck.Card) bool
	Stand()
}
