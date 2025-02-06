package main

import (
	"encoding/csv"
	"fmt"
	"log"
	// "math/rand"
	"os"
)

type game struct {
	title   string
	console string
	genre   string
}

func new_game(input []string) game {
	return game{
		title:   input[0],
		console: input[1],
		genre:   input[2],
	}
}

func parse_csv(records [][]string, filter []string) ([]game, error) {
	var games []game
	exclude := false
	if filter[2] == "exclude" {
		exclude = true
	}

	for i := 1; i < len(records); i++ {
		new := new_game(records[i])

		if filter[0] == "" { // No Filter, just add game
			games = append(games, new)
		} else if exclude { // Skip this genre/console if match
			if filter[1] == new.console || filter[1] == new.genre {
				continue
			} else {
				games = append(games, new)
			}
		} else if !exclude { // Skip this genre/console if not match
			if filter[1] != new.console || filter[1] != new.genre {
				continue
			} else {
				games = append(games, new)
			}
		}
	}

	if len(games) <= 0 {
		fmt.Println("There are no games in the list.")
	}
	return games, nil
}

func load_csv(filename string) [][]string {
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

func main() {
	filename := "test.csv"
	records := load_csv(filename)
	filter := get_filter()
	Games := parse_csv(records)
}
