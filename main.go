package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var (
		numberOfPlayers int
		names           []string
		questions       []Question
	)
	
	const filename = "questions.txt"
	
	questions, err := loadQuestionsFromFile(filename)
	
	if err != nil {
		fmt.Printf("Error reading file %s\n", filename)
		os.Exit(1)
	}
	
	// get number of players for the game
	for numberOfPlayers < 2 {
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
	}
	
	fmt.Println(playGame(numberOfPlayers, names, questions))
	
}
