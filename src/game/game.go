package game

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	// "strings"
	"randomizer/src/csv_data"
	"randomizer/src/pray"
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

func FilterConsole(games []Game, filter Filter) ([]Game, error) {
	// Remove games for a particular console
	var new_games, removed_games []Game
	for i := 0; i < len(games); i++ {
		if games[i].GameFilter(filter) {
			removed_games = append(removed_games, games[i])
		} else {
			new_games = append(new_games, games[i])
		}
	}
	// Check for empty list errors, and return requested list
	var final_list []Game
	var err error
	if filter.Exclude {
		err = ListCheck(new_games)
		final_list = new_games
	} else {
		err = ListCheck(removed_games)
		final_list = removed_games
	}
	return final_list, err
}

func ListCheck(games []Game) error {
	if len(games) == 0 {
		err := errors.New("game list is empty")
		return err
	}
	return nil
}

func CSVErrorTest(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return true
	}
	return false
}

func GenerateCSV(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, record := range csv_data.Data {
		err := writer.Write(record)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	fmt.Println("Default csv file generated.")
}

func LoadCsv(filename string) ([][]string, error) {
	// Load a .csv file and return nested arrays of file
	if CSVErrorTest(filename) {
		GenerateCSV(filename)
	}
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
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

func CSVRandomizer(filename string) {
	for {
		games, err := LoadCsv(filename)
		if err != nil {
			log.Fatal(err)
		}
		full_games := ParseCSVRecords(games)

		choice := rand.Intn(len(full_games))

		game := full_games[choice].GameString()
		next := pray.Prayer(game)
		switch next {
		case "q":
			os.Exit(0)
		case "y":
			continue
		default:
			return
		}
	}
}
