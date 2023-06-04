package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func readNameFromInput(index int) (string, error) {
	fmt.Printf("Enter the name for player #%d: ", index)
	var name string
	fmt.Scanln(&name)
	if len(name) < 2 {
		return name, errors.New("minimum length is 2")
	}
	
	return name, nil
}

func readAliasAnswer() string {
	var answer string
	fmt.Scanln(&answer)
	return answer
}

func inputNumberOfPlayers() (int, error) {
	var numberOfPlayers int
	
	fmt.Print("Enter the number of players: ")
	_, err := fmt.Scanln(&numberOfPlayers)
	
	if numberOfPlayers < 2 {
		err = errors.New("minimum players required 2")
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

type Node struct {
	Value string
	Next  *Node
}

func CreateCyclicalList(values []string) *Node {
	if len(values) == 0 {
		return nil
	}
	
	// Create the first node
	first := &Node{Value: values[0]}
	current := first
	
	// Create the remaining nodes and link them
	for i := 1; i < len(values); i++ {
		next := &Node{Value: values[i]}
		current.Next = next
		current = next
	}
	
	// Make the last node point back to the first node
	current.Next = first
	
	return first
}
