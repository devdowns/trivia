package main

import (
	"testing"
)

func Test_main(t *testing.T) {
	// Given
	numberOfPlayers := 2
	names := []string{"dev", "kushuki"}
	questions := []Question{
		{Question: "What is the capital of France?", IsReverse: false},
		{Question: "What is the name of the largest ocean in the world?", IsReverse: true},
	}
	
	// When
	expectedWinner := "dev"
	actualWinner := playGame(numberOfPlayers, names, questions)
	
	// Then
	if actualWinner != expectedWinner {
		t.Errorf("Expected winner to be %s, but got %s", expectedWinner, actualWinner)
	}
}
