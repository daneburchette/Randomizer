package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	// local import
	"randomizer/src/csv_data"
)

// Constants and utilities

const defaultCSVFile = "games.csv"

var MenuList []string = []string{
	"CSV Randomizer, Single",
	"CSV Randomizer, Multi",
	"Manual Entry Randomizer",
	"Exit",
}

func getUserInput(prompt string) string {
	fmt.Println(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func clearScreen() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error clearing the screen:", err)
	}
}

// Main menu and main function

func menu() {
	var menu_option int
	var err error
	for {
		clearScreen()
		fmt.Println("Our Dark God: The Randomizer! (All Hail)")
		for i := 0; i < len(MenuList); i++ {
			fmt.Printf("\t%d - %s\n", (i + 1), MenuList[i])
		}
		fmt.Printf("\nChoose Randomzier Mode [1-%d]: ", len(MenuList))
		input := getUserInput("")
		menu_option, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}
		if menu_option > 0 && menu_option <= len(MenuList) {
			break
		}
		fmt.Printf("Invalid input. Please enter a number from 1-%d: ", len(MenuList))
	}
	fmt.Println(menu_option)
	filename := defaultCSVFile
	switch menu_option {
	case 1:
		csvRandomizerSingle(filename)
	case 2:
		csvRandomizerMulti(filename)
	case 3:
		manual()
	default:
		os.Exit(0)
	}
}

func choose_csv() string {
	input := getUserInput("Enter name of custom csv file or leave blank for default: ")
	if input == "" {
		return defaultCSVFile
	} else {
		return input
	}
}

func main() {
	for {
		menu()
	}
}

// CSV Handling

func csvFileMissing(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return true
	}
	return false
}

func createDefaultCSV(filename string) {
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

func loadCSV(filename string) ([][]string, error) {
	// Load a .csv file and return nested arrays of file
	if csvFileMissing(filename) {
		createDefaultCSV(filename)
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

func parseCSVRecords(input [][]string) []Game {
	// Convert loaded csv into an array of Game structs
	var games []Game
	for i := 1; i < len(input); i++ {
		new_game := newGame(input[i])
		games = append(games, new_game)
	}
	return games
}

func getConsoleList(input []Game) []string {
	var consoleList []string
	for _, game := range input {
		if !contains(consoleList, game.Console) {
			consoleList = append(consoleList, game.Console)
		}
	}
	return consoleList
}

func contains(slice []string, str string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}
	return false
}

func csvRandomizerMulti(filename string) {
	var count int
	clearScreen()
	games, err := loadCSV(filename)
	if err != nil {
		log.Fatal(err)
	}
	for {
		var err error
		input := getUserInput("How many games are you praying for?")
		count, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Please enter a number: ")
			continue
		}
		if count > len(games)-1 {
			fmt.Printf("You cannot pray for more games than are in the list! Choose a number smaller than %d\n", len(games)-1)
			continue
		}
		break
	}
	for {
		clearScreen()
		full_games := parseCSVRecords(games)
		selectedGames := randomizeGames(full_games, count)
		next := prayer(selectedGames)
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

func randomizeGames(games []Game, count int) []string {
	// Helper function for csvRandomizeMulti
	if count > len(games) {
		log.Fatal("Requested more games than available.")
	}
	selectedIndices := make(map[int]struct{})
	var selectedGames []string
	for len(selectedGames) < count {
		randomIndex := rand.Intn(len(games))
		if _, exists := selectedIndices[randomIndex]; !exists {
			selectedGames = append(selectedGames, games[randomIndex].GameString())
			selectedIndices[randomIndex] = struct{}{}
		}
	}
	return selectedGames
}

func csvRandomizerSingle(filename string) {
	for {
		clearScreen()
		games, err := loadCSV(filename)
		if err != nil {
			log.Fatal(err)
		}
		full_games := parseCSVRecords(games)
		// consoles := getConsoleList(full_games)

		// add filter prompts here, possible a bool to bypass filter when loading from menu
		choice := rand.Intn(len(full_games))

		var game []string
		game = append(game, full_games[choice].GameString())
		next := prayer(game)
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

// Game struct and methods

type Game struct {

	// Structure for a retro game
	Title   string
	Console string
}

func newGame(input []string) Game {
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

// Filter struct and methods

type Filter struct {
	// Structure for a list of console to filter
	Console []string
	Exclude bool
}

func filterConsole(games []Game, filter Filter) ([]Game, error) {
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
		err = listCheck(new_games)
		final_list = new_games
	} else {
		err = listCheck(removed_games)
		final_list = removed_games
	}
	return final_list, err
}

func listCheck(games []Game) error {
	if len(games) == 0 {
		err := errors.New("game list is empty")
		return err
	}
	return nil
}

// Manual Entry

func gamePrompt() ([]string, error) {
	var games []string
	fmt.Println("Enter game titles to randomize, one at a time:")
	fmt.Println("(Enter blank when done)")
	for {
		game := getUserInput("")
		if game == "" {
			break
		}
		games = append(games, game)
	}

	if len(games) <= 1 {
		return games, errors.New("please enter at least two games to randomize")
	}

	return games, nil
}

func manual() {
	for {
		clearScreen()
		games, err := gamePrompt()
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}
		choice := rand.Intn(len(games))
		var gameChoices []string
		gameChoices = append(gameChoices, games[choice])
		next := prayer(gameChoices)
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

// Prayer

func prayer(choice []string) string {
	fmt.Println("Place your head on the BatterUp peripheral,")
	fmt.Println("Point your No-No hole in a random direction")
	fmt.Println("and say the prayer be all love to say:")
	fmt.Scanln()
	repeatPrayerPhrase("Bee-da-bud-a-bud-a", 3)
	repeatPrayerPhrase("Boop", 3)
	fmt.Println("No whammies...")
	fmt.Scanln()
	fmt.Println("No Whammies!")
	fmt.Scanln()
	fmt.Println("NO WHAMMIES!")
	fmt.Scanln()
	fmt.Println("STOP!!!!")
	fmt.Scanln()
	for i := 0; i < len(choice); i++ {
		fmt.Println(choice[i])
	}

	var answer string
	fmt.Println("Randomize again? (Y)es/(N)o/(Q)uit:")
	for {
		fmt.Scanln(&answer)
		answer = strings.ToLower(strings.TrimSpace(answer))
		if answer == "y" || answer == "n" || answer == "q" {
			break
		}
		fmt.Println("Invalid input. Please enter (Y)es/(N)o/(Q)uit:")
	}
	return answer
}

func repeatPrayerPhrase(phrase string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(phrase)
		fmt.Scanln()
	}
}
