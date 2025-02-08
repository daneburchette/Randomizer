package manual_randomizer

import (
	"errors"
	"fmt"
	"math/rand"
	// Local imports
	"randomizer/src/pray"
)

func game_prompt() ([]string, error) {
	var games []string

	fmt.Println("Enter game titles to randomize, one at a time:")
	fmt.Println("(Enter blank when done)")
	for {
		var game string
		fmt.Scanln(&game)
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

func main() {
	for {
		games, err := game_prompt()
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}
		choice := rand.Intn(len(games))
		pray.Prayer(games[choice])

		fmt.Println("Do you want to randomize again? (yes/no)")
		var response string
		fmt.Scanln(&response)
		if response != "yes" && response != "y" {
			break
		}
	}
}
