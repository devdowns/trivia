package main

import (
	"fmt"
	"sort"
	"strings"
)

type Player map[string]PlayerStats

type PlayerStats struct {
	Name      string
	Points    int
	Questions []string
	Votes     int
	Alias     string
}

func (p *PlayerStats) PrintStats() string {
	var stats strings.Builder
	
	stats.WriteString(fmt.Sprintf("%s -> ", p.Name))
	stats.WriteString(fmt.Sprintf("%d pts ", p.Points))
	
	for _, question := range p.Questions {
		stats.WriteString(question + "\n")
	}
	
	return stats.String()
}

func (p *Player) FindWinners(isReverseQuestion bool) map[string]bool {
	if isReverseQuestion {
		return p.findWinnersByVotesReverse()
	} else {
		return p.findWinnersByVotes()
	}
}

func (p *Player) findWinnersByVotes() map[string]bool {
	playerStats := make([]PlayerStats, 0, len(*p))
	winners := make(map[string]bool)
	
	for _, player := range *p {
		playerStats = append(playerStats, player)
	}
	
	// Sort in descending order to determine the highest score
	sort.Slice(playerStats, func(i, j int) bool {
		return playerStats[i].Votes > playerStats[j].Votes
	})
	
	maxVotes := playerStats[0].Votes
	
	// Get the aliases of all players with the highest score
	for key, player := range *p {
		if player.Votes == maxVotes {
			winners[key] = true
		}
	}
	
	return winners
}

func (p *Player) findWinnersByVotesReverse() map[string]bool {
	playerStats := make([]PlayerStats, 0, len(*p))
	winners := make(map[string]bool)
	
	for _, player := range *p {
		playerStats = append(playerStats, player)
	}
	
	// Sort in ascending order to determine the lowest score
	sort.Slice(playerStats, func(i, j int) bool {
		return playerStats[i].Votes < playerStats[j].Votes
	})
	
	minVotes := playerStats[0].Votes
	
	// Get the aliases of all players with the lowest score
	for key, player := range *p {
		if player.Votes == minVotes {
			winners[key] = true
		}
	}
	
	return winners
}

func (p *Player) findWinnersByPoints() string {
	var playerStats []PlayerStats
	var winners []string
	
	for _, player := range *p {
		playerStats = append(playerStats, player)
	}
	
	// sort in asc order to figure out which is the max score
	sort.Slice(playerStats, func(i, j int) bool {
		return playerStats[i].Points > playerStats[j].Points
	})
	
	maxPoints := playerStats[0].Points
	
	// get the aliases of all the players with the max score
	for _, player := range *p {
		if player.Points == maxPoints {
			winners = append(winners, player.Name)
		}
	}
	
	return strings.Join(winners, ", ")
}

func (p *Player) updateScores(voteWinners map[string]bool, question string) {
	for player, data := range *p {
		// Check if player exists in the winner's map. If so, update their points.
		if _, ok := voteWinners[player]; ok {
			data.Points++
			// uncomment to print all the questions that gave you points
			//data.Questions = append(data.Questions, question)
		}
		data.Votes = 0
		(*p)[player] = data
	}
}

func (p *Player) getAliasesSlice() []string {
	// Get the keys of the Player map
	keys := make([]string, 0, len(*p))
	for key := range *p {
		keys = append(keys, key[0:2])
	}
	
	// Join the keys into a single string
	return keys
}

func (p *Player) getNamesSlice() []string {
	// Get the keys of the Player map
	keys := make([]string, 0, len(*p))
	for _, val := range *p {
		keys = append(keys, val.Name)
	}
	
	// Join the keys into a single string
	return keys
}

func (p *Player) PrintStats() string {
	statsSummary := strings.Builder{}
	for _, val := range *p {
		statsSummary.WriteString(val.PrintStats() + "| ")
	}
	return statsSummary.String()
}

func getPlayerChoices(aliases []string, currentPlayerAlias string) string {
	
	// Create a new slice to store the filtered aliases
	filteredAliases := make([]string, 0, len(aliases))
	
	// Iterate over the aliases and filter out the current player's alias
	for _, alias := range aliases {
		if alias != currentPlayerAlias {
			filteredAliases = append(filteredAliases, alias)
		}
	}
	
	// Join the filtered aliases with commas and return the result
	return strings.Join(filteredAliases, ", ")
}
