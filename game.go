package main

import (
	"fmt"
	"strings"
)

func playGame(numberOfPlayers int, names []string, questions []Question) string {
	var (
		playersMap  Player
		winningMsg  strings.Builder
		gameWinners string
		players     *Node
	)
	
	playersMap = make(Player)
	
	/*
		Register the players, the key to the players map is the alias,
		the alias is composed of the first two letters of each names
		to expedite the speed at which the game is played
	*/
	for i := 0; i < numberOfPlayers; i++ {
		
		playersMap[names[i][0:2]] =
			PlayerStats{
				Name:      names[i],
				Points:    0,
				Questions: nil,
				Votes:     0,
				Alias:     names[i][0:2]}
	}
	
	players = CreateCyclicalList(playersMap.getNamesSlice())
	aliasesSlice := playersMap.getAliasesSlice()
	
	shuffleQuestions(questions)
	
	for i, question := range questions {
		clearScreen()
		i += 1
		
		fmt.Printf("Player %s draws a card\n", players.Value)
		fmt.Printf("Question #%d \n%s\n", i, question.Question)
		for currentPlayer := range playersMap {
			ok := false
			answer := ""
			
			// repeats until you pick a valid player alias to nominate
			for !ok {
				currentPlayerAlias := playersMap[currentPlayer].Alias
				// prints player aliases available for you to nominate from
				personalizedChoices := getPlayerChoices(aliasesSlice, currentPlayerAlias)
				//fmt.Printf("Nominate one player using their alias: %s\n", personalizedChoices)
				fmt.Printf("You're up %s, what's your answer %s?\n",
					playersMap[currentPlayer].Name,
					personalizedChoices)
				
				answer = readAliasAnswer()
				
				// ensure you can't choose yourself
				if answer != currentPlayerAlias {
					_, ok = playersMap[answer]
				}
			}
			
			// update player stats
			data := playersMap[answer]
			data.Votes++
			playersMap[answer] = data
			
		}
		
		// determine which strategy should be used to find the winner
		voteWinners := playersMap.FindWinners(question.IsReverse)
		
		// update player scores after each question
		playersMap.updateScores(voteWinners, question.Question)
		
		players = players.Next
	}
	
	clearScreen()
	
	gameWinners = playersMap.findWinnersByPoints()
	
	winningMsg.WriteString(fmt.Sprintf("Congratulations %s you won!!!\n\n", gameWinners))
	
	winningMsg.WriteString("Player Stats\n")
	
	winningMsg.WriteString(playersMap.PrintStats())
	
	return winningMsg.String()
}
