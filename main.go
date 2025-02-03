package main

import (
	"fmt"
	"math/rand"
)

func main() {
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

	choice := rand.Intn(len(games))

	fmt.Println("Place your head on the BatterUp peripheral,")
	fmt.Println("Point your No-No hole in a random direction")
	fmt.Println("and say the prayer be all love to say:")
	fmt.Println()
	for i := 0; i < 3; i++ {
		fmt.Println("Bee-da-bud-a-bud-a")
		fmt.Println()
	}
	for i := 0; i < 3; i++ {
		fmt.Println("Boop")
		fmt.Println()
	}
	fmt.Println("No whammies...")
	fmt.Println()
	fmt.Println("No Whammies!")
	fmt.Println()
	fmt.Println("NO WHAMMIES!")
	fmt.Println()
	fmt.Println("STOP!!!!")
	fmt.Println()
	fmt.Println(games[choice])
}
