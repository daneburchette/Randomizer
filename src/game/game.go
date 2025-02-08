package game

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
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
	if filter.Exclude {
		for i := 0; i < len(filter.Console); i++ {
			if filter.Console[i] == g.Console {
				return false
			} else {
				continue
			}
		}
		return true
	} else {
		for i := 0; i < len(filter.Console); i++ {
			if filter.Console[i] == g.Console {
				return true
			} else {
				continue
			}
		}
		return false
	}
}

type Filter struct {
	// Structure for a list of console to filter
	Console []string
	Exclude bool
}

func FilterConsole(games []Game, filter Filter) ([]Game, error) {
	// REFACTOR! //
	// Use GameFilter method to filter games
	// Remove games for a particular console
	var new_games, removed_games []Game
	for i := 0; i < len(games); i++ {
		for j := 0; j < len(filter.Console); j++ {
			if games[i].Console == filter.Console[j] {
				removed_games = append(removed_games, games[i])
			} else {
				new_games = append(new_games, games[i])
			}
		}
	}
	// Check for empty list errors, and return requested list
	if filter.Exclude {
		if len(new_games) <= 0 {
			return new_games, errors.New("game list is empty")
		}
		return new_games, nil
	}
	if len(removed_games) <= 0 {
		return removed_games, errors.New("game list is empty")
	}
	return removed_games, nil
}
func LoadCsv(filename string) [][]string {
	// Load a .csv file and return nested arrays of file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return records
}

func ParseCSVRecords(input [][]string) []Game {
	// Convert loaded csv into an array of Game structs
	var games []Game
	for i := 1; i < len(input); i++ {
		new_game := NewGame(input[i])
		games = append(games, new_game)
	}
	return games
}
