package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	//local import
	"randomizer/src/auto"
	"randomizer/src/manual"
)

var MenuList []string = []string{
	"CSV Randomizer",
	"Manual Entry Randomizer",
	"Exit",
}

func menu() {
	var menu_option int
	for {
		fmt.Println("Our Dark God: The Randomizer! (All Hail)")
		for i := 0; i < len(MenuList); i++ {
			fmt.Printf("\t%d - %s", i, MenuList[i])
		}
		fmt.Printf("\nChoose Randomzier Mode [1-%d]", len(MenuList))
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}
		if choice > 0 && choice <= len(MenuList) {
			break
		}
		fmt.Printf("Invalid input. Please enter a number from 1-%d:", len(MenuList))
	}
	switch menu_option {
	case 1:
		filename := choose_csv()
		auto.CSVRandomizer(filename)
	case 2:
		manual_randomizer.Manual()
	default:
		os.Exit(0)
	}
}

func choose_csv() string {
	fmt.Println("Enter name of custom csv file or leave blank for default: ")
	scanner := bufio.NewScanner(os.Stdin)
	input := strings.TrimSpace(scanner.Text())
	if input == "" {
		return "games.csv"
	} else {
		return input
	}
}

func main() {
	menu()
}
