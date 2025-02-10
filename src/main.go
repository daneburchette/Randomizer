package main

import (
	"fmt"
	"os"
	"strings"

	//local import
	"randomizer/src/game"
	"randomizer/src/manual"
)

func menu() {
	var choice string
	fmt.Println("Our Dark God: The Randomizer! (All Hail)")
	for {
		fmt.Println("\n\t1 - CSV Randomizer\n\t2 - Manual Entry Randomizer\n\t3 - Exit")
		fmt.Println("\nChoose Randomzier Mode [1-3]")
		fmt.Scanln(&choice)
		choice = strings.ToLower(strings.TrimSpace(choice))
		if choice == "1" || choice == "2" || choice == "3" {
			break
		}
		fmt.Println("Invalid input. Please enter a number from 1-3:")
	}
	switch choice {
	case "1":
		game.CSVRandomizer("games.csv")
	case "2":
		manual_randomizer.Manual()
	case "3":
		os.Exit(0)
	}
}

func main() {
	for {
		menu()
	}
}
