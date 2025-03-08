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
	"slices"
	"strconv"
	"strings"

	// Local import
	"randomizer/src/csv_data"
)

// Constants and utility variables
const defaultCSVFile = "games.csv"

// Menu options for the user
var menuOptions []string = []string{
	"CSV Randomizer, Single",
	"CSV Randomizer, Multi",
	"Manual Entry Randomizer",
	"Exit",
}

// getUserInput prompts the user with a message and returns their input as a string
func getUserInput(prompt string) string {
	fmt.Println(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

// clearScreen clears the console screen depending on the operating system
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

// menu displays the main menu to the user and handles their selection
func menu() {
	var selectedOption int
	var err error
	for {
		clearScreen()
		fmt.Println("Our Dark God: The Randomizer! (All Hail)")
		for i, menuOption := range menuOptions {
			fmt.Printf("\t%d - %s\n", i+1, menuOption)
		}
		fmt.Printf("\nChoose Randomizer Mode [1-%d]: ", len(menuOptions))
		input := getUserInput("")
		selectedOption, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}
		if selectedOption > 0 && selectedOption <= len(menuOptions) {
			break
		}
		fmt.Printf("Invalid input. Please enter a number from 1-%d: ", len(menuOptions))
	}
	fmt.Println(selectedOption)
	filename := defaultCSVFile
	switch selectedOption {
	case 1:
		csvRandomizer(filename, false, false)
	case 2:
		csvRandomizer(filename, true, false)
	case 3:
		manualEntryRandomizer()
	default:
		os.Exit(0)
	}
}

// chooseCSV prompts the user to choose a custom CSV file or use the default one
func chooseCSV() string {
	input := getUserInput("Enter name of custom CSV file or leave blank for default: ")
	if input == "" {
		return defaultCSVFile
	} else {
		return input
	}
}

// main function which runs the program and calls the menu in a loop
func main() {
	for {
		menu()
	}
}

// CSV Handling Functions

// csvFileMissing checks if a CSV file exists
func csvFileMissing(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return true
	}
	return false
}

// createDefaultCSV generates a default CSV file if none exists
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
	fmt.Println("Default CSV file generated.")
}

// loadCSV loads the CSV file and returns its content as a 2D string array
func loadCSV(filename string) ([][]string, error) {
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

// parseCSVRecords converts the CSV records into a slice of Game structs
func parseCSVRecords(input [][]string) []Game {
	var games []Game
	for i := 1; i < len(input); i++ {
		newGame := newGame(input[i])
		games = append(games, newGame)
	}
	return games
}

// getConsoleList extracts a list of unique consoles from the games
func getConsoleList(input []Game) []string {
	var consoleList []string
	for _, game := range input {
		if !slices.Contains(consoleList, game.Console) {
			// if !contains(consoleList, game.Console) {
			consoleList = append(consoleList, game.Console)
		}
	}
	return consoleList
}

// contains checks if a string is present in a slice of strings
// func contains(slice []string, str string) bool {
// 	for _, item := range slice {
// 		if item == str {
// 			return true
// 		}
// 	}
// 	return false
// }

// csvRandomizer allows the user to select multiple games randomly from the CSV file
func csvRandomizer(filename string, multi bool, filter bool) {
	var gameCount int = 1
	clearScreen()
	games, err := loadCSV(filename)
	if err != nil {
		fmt.Println("Error loading CSV file: ", err)
		return
	}
	if multi {
		for {
			var err error
			input := getUserInput("How many games would you like to randomize?")
			gameCount, err = strconv.Atoi(input)
			if err != nil || gameCount <= 0 {
				fmt.Println("Please enter a valid positive number.")
				continue
			} else if gameCount > len(games)-1 {
				fmt.Printf("You cannot select more games than are in the list! Choose a smaller number.\n")
				continue
			}
			break
		}
	}
	if filter {
		fmt.Println("Not Yet Implemented!")
		return
	}
	for {
		clearScreen()
		fullGames := parseCSVRecords(games)
		selectedGames := randomizeGames(fullGames, gameCount)
		if prayer(selectedGames) {
			continue
		} else {
			break
		}
	}
}

// randomizeGames selects a specified number of random games from the list
func randomizeGames(games []Game, count int) []string {
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

// Game struct represents a retro game with a title and console
type Game struct {
	Title   string
	Console string
}

// newGame creates a new Game struct from a set of CSV values
func newGame(input []string) Game {
	return Game{
		Title:   input[0],
		Console: input[1],
	}
}

// GameString formats the Game struct into a readable string
func (g *Game) GameString() string {
	if g.Console == "" {
		return fmt.Sprintf("%s", g.Title)
	}
	return fmt.Sprintf("%s (%s)", g.Title, g.Console)
}

// GameFilter filters a game based on a filter
func (g *Game) GameFilter(filter Filter) bool {
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

// Filter struct for filtering games based on console and exclusion
type Filter struct {
	Console []string
	Exclude bool
}

// filterConsole applies a filter to a list of games and returns the filtered list
func filterConsole(games []Game, filter Filter) ([]Game, error) {
	var remainingGames, removedGames []Game
	for i := range len(games) {
		if games[i].GameFilter(filter) {
			removedGames = append(removedGames, games[i])
		} else {
			remainingGames = append(remainingGames, games[i])
		}
	}
	var finalList []Game
	var err error
	if filter.Exclude {
		err = listCheck(remainingGames)
		finalList = remainingGames
	} else {
		err = listCheck(removedGames)
		finalList = removedGames
	}
	return finalList, err
}

// listCheck checks if the list of games is empty
func listCheck(games []Game) error {
	if len(games) == 0 {
		return errors.New("game list is empty")
	}
	return nil
}

// Manual Entry Functions

// gamePrompt asks the user to enter game titles manually
func gamePrompt() ([]Game, error) {
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
	var fullGames []Game
	if len(games) <= 1 {
		return fullGames, errors.New("please enter at least two games to randomize")
	}
	for i := range len(games) {
		fullGames = append(fullGames, Game{Title: games[i], Console: ""})
	}
	return fullGames, nil
}

// manualEntryRandomizer handles the manual entry randomization process
func manualEntryRandomizer() {
	for {
		clearScreen()
		games, err := gamePrompt()
		if err != nil {
			fmt.Println("Error: ", err)
			fmt.Scanln()
			continue
		}
		clearScreen()
		selectedGames := randomizeGames(games, 1)
		if prayer(selectedGames) {
			continue
		} else {
			break
		}
	}
}

// Prayer Functions

// prayer simulates the randomizer prayer process and returns user's choice
func prayer(choice []string) bool {
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
	for i := range len(choice) {
		fmt.Println(choice[i])
	}
	return shouldContinue()
}

// repeatPrayerPhrase repeats the prayer phrase a specified number of times
func repeatPrayerPhrase(phrase string, times int) {
	for range times {
		fmt.Println(phrase)
		fmt.Scanln()
	}
}

// shouldContinue asks the user if they want to randomize again
func shouldContinue() bool {
	var input string
	for {
		input = getUserInput("\nRandomize again (Y)es/(N)o/(Q)uit:")
		if input == "y" || input == "n" || input == "q" {
			break
		}
		fmt.Println("Invalid input. Please enter (Y)es/(N)o/(Q)uit:")
	}
	switch input {
	case "q":
		os.Exit(0)
		return false
	case "y":
		return true
	default:
		return false
	}
}
