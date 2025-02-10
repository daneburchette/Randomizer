package game

import (
	"fmt"
)

type Game struct {

	// Structure for a retro game
	Title   string
	Console string
}

func NewGame(input []string) Game {
	// Parse set of strings into game
	return Game{
		Title:   input[0],
		Console: input[1],
	}
}

func (g *Game) GameString() string {
	return fmt.Sprintf("%s (%s)", g.Title, g.Console)
}

func (g *Game) GameFilter(filter Filter) bool {
	// True is gamee is to remain, false if it is to be removed
	consoleSet := make(map[string]struct{})
	for _, console := range filter.Console {
		consoleSet[console] = struct{}{}
	}

	_, exists := consoleSet[g.Console]
	if filter.Exclude {
		return !exists
	}
	return exists
}

type Filter struct {
	// Structure for a list of console to filter
	Console []string
	Exclude bool
}
