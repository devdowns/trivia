package main

import (
	"fmt"
	"github.com/devdowns/trivia/types"
	"strings"
)

func playGame(names []string, questions types.Questions) string {
	var (
		playerTurn  *types.Turn
		playersMap  types.Player
		gameWinners string
		winningMsg  strings.Builder
	)
	
	// setup
	playersMap = types.RegisterPlayers(names)
	
	types.ConfigurePlayerChoices()
	
	playerTurn = types.AssignPlayerTurns(names)
	
	for _, question := range questions {
		
		clearScreen()
		
		fmt.Printf("ScoreBoard: %s\n", playersMap.GetScoreBoard())
		fmt.Printf("Player %s draws a card\n", playerTurn.PlayerName)
		fmt.Printf("Question #%d :%s\n", playerTurn.TurnNumber, question.Text)
		for _, currentPlayerName := range names {
			ok := false
			answer := ""
			
			// repeats until you pick a valid player alias to nominate
			for !ok {
				fmt.Printf("Votes: %s\n", playersMap.GetVotes())
				currentPlayerAlias := playersMap[currentPlayerName[0:2]].Alias
				fmt.Printf("You're up %s, what's your answer %s?\n",
					playersMap[currentPlayerAlias].Name,
					playersMap[currentPlayerAlias].Choices)
				
				answer = readAliasAnswer()
				
				// ensure you can't choose yourself
				if answer != currentPlayerAlias {
					_, ok = playersMap[answer]
				}
			}
			
			// update player stats
			playersMap.UpdatePlayerPoints(answer)
			
		}
		
		// determine which strategy should be used to find the winner
		voteWinners := playersMap.FindRoundWinners(question.IsReverse)
		
		// update player scores after each question
		playersMap.UpdateScores(voteWinners, question.Text)
		
		playerTurn.Next()
	}
	
	clearScreen()
	
	gameWinners = playersMap.FindGameWinners()
	
	winningMsg.WriteString(fmt.Sprintf(
		"Congratulations %s you won!!!\n\nScoreboard: %s",
		gameWinners,
		playersMap.GetScoreBoard()))
	
	return winningMsg.String()
}
