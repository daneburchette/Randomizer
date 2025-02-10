package manual_randomizer

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"

	// Local imports
	"randomizer/src/pray"
)

func game_prompt() ([]string, error) {
	var games []string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter game titles to randomize, one at a time:")
	fmt.Println("(Enter blank when done)")
	for {
		var game string
		scanner.Scan()
		game = strings.TrimSpace(scanner.Text())
		if game == "" {
			break
		}
		games = append(games, game)
	}

	if len(games) == 0 {
		return games, errors.New("cannot randomize without games")
	}

	return games, nil
}

func Manual() {
	for {
		games, err := game_prompt()
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}
		choice := rand.Intn(len(games))

		next := pray.Prayer(games[choice])
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
