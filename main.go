package main

import (
	"errors"
	"fmt"
	"github.com/devdowns/trivia/types"
	"log"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	var (
		numberOfPlayers int
		names           []string
		questions       types.Questions
	)
	
	const filename = "questions.txt"
	
	questions, err := types.LoadQuestionsFromFile(filename)
	
	if err != nil {
		fmt.Printf("Error reading file %s\n", filename)
		os.Exit(1)
	}
	
	// get number of players for the game
	for numberOfPlayers < 3 {
		numberOfPlayers, err = inputNumberOfPlayers()
		if err != nil {
			fmt.Println(err)
			numberOfPlayers = 0
		}
	}
	
	// register player data
	for i := 0; i < numberOfPlayers; i++ {
		name, err := readNameFromInput(i + 1)
		if err != nil {
			log.Fatal(err)
		}
		names = append(names, name)
		fmt.Printf("Welcome in %s!\n", name)
	}
	
	fmt.Printf("Loaded %d questions\n", len(questions))
	
	fmt.Println("Press Enter to Start...")
	fmt.Scanln() // Wait for the user to press Enter
	
	fmt.Println(playGame(names, questions))
	
}

func readNameFromInput(index int) (string, error) {
	fmt.Printf("Enter the name for player #%d: ", index)
	var name string
	fmt.Scanln(&name)
	if len(name) < 2 {
		return name, errors.New("minimum length is 2")
	}
	
	return name, nil
}

func getAnswer() string {
	var answer string
	fmt.Scanln(&answer)
	return answer
}

func inputNumberOfPlayers() (int, error) {
	var numberOfPlayers int
	
	fmt.Print("Enter the number of players: ")
	_, err := fmt.Scanln(&numberOfPlayers)
	
	if numberOfPlayers < 3 {
		err = errors.New("minimum players required 3")
	}
	
	return numberOfPlayers, err
}

func clearScreen() {
	switch runtime.GOOS {
	case "linux", "darwin": // Unix-like systems
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows": // Windows
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Println("Unsupported platform. Unable to clear console screen.")
	}
}
