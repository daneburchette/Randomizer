package auto

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	// local imports
	"randomizer/src/csv_data"
	"randomizer/src/game"
	"randomizer/src/pray"
)

func filterConsole(games []game.Game, filter game.Filter) ([]game.Game, error) {
	// Remove games for a particular console
	var new_games, removed_games []game.Game
	for i := 0; i < len(games); i++ {
		if games[i].GameFilter(filter) {
			removed_games = append(removed_games, games[i])
		} else {
			new_games = append(new_games, games[i])
		}
	}
	// Check for empty list errors, and return requested list
	var final_list []game.Game
	var err error
	if filter.Exclude {
		err = listCheck(new_games)
		final_list = new_games
	} else {
		err = listCheck(removed_games)
		final_list = removed_games
	}
	return final_list, err
}

func listCheck(games []game.Game) error {
	if len(games) == 0 {
		err := errors.New("game list is empty")
		return err
	}
	return nil
}

func csvErrorTest(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return true
	}
	return false
}

func generateCSV(filename string) {
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

func loadCsv(filename string) ([][]string, error) {
	// Load a .csv file and return nested arrays of file
	if csvErrorTest(filename) {
		generateCSV(filename)
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

func parseCSVRecords(input [][]string) []game.Game {
	// Convert loaded csv into an array of Game structs
	var games []game.Game
	for i := 1; i < len(input); i++ {
		new_game := game.NewGame(input[i])
		games = append(games, new_game)
	}
	return games
}

func CSVRandomizer(filename string) {
	for {
		games, err := loadCsv(filename)
		if err != nil {
			log.Fatal(err)
		}
		full_games := parseCSVRecords(games)

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
