package types

import (
	"fmt"
	"sort"
	"strings"
)

type Player map[string]PlayerStats

var playersMap Player
var playerAliases []string

/*
	Create custom choices for player, these include all other player alias
	but their own
*/
func ConfigurePlayerChoices() {
	for key := range playersMap {
		data := playersMap[key]
		data.Choices = fmt.Sprintf("( %s )", getPlayerChoices(playerAliases, data.Alias))
		playersMap[key] = data
	}
	fmt.Println(playersMap)
}

func RegisterPlayers(names []string) Player {
	playersMap = make(Player)
	for _, name := range names {
		alias := name[:2]
		
		playersMap[alias] = PlayerStats{
			Name:      name,
			Points:    0,
			Questions: nil,
			Votes:     0,
			Alias:     alias,
		}
		playerAliases = append(playerAliases, alias)
	}
	return playersMap
}

func (p *Player) FindRoundWinners(isReverseQuestion bool) map[string]bool {
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

func (p *Player) FindGameWinners() string {
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

func (p *Player) UpdateScores(voteWinners map[string]bool, question string) {
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

func (p *Player) GetScoreBoard() string {
	// Convert the map to a slice of player stats
	var playerStats []PlayerStats
	for _, stat := range *p {
		playerStats = append(playerStats, stat)
	}
	
	// Sort the player stats in descending order based on points
	sort.Slice(playerStats, func(i, j int) bool {
		return playerStats[i].Points > playerStats[j].Points
	})
	
	// Generate the scoreboard
	statsSummary := strings.Builder{}
	for _, stat := range playerStats {
		statsSummary.WriteString(stat.printScore() + "| ")
	}
	return statsSummary.String()
}

func (p *Player) GetVotes() string {
	// Convert the map to a slice of player stats
	var playerStats []PlayerStats
	for _, stat := range *p {
		playerStats = append(playerStats, stat)
	}
	
	// Sort the player stats in descending order based on points
	sort.Slice(playerStats, func(i, j int) bool {
		return playerStats[i].Votes > playerStats[j].Votes
	})
	
	// Generate the scoreboard
	statsSummary := strings.Builder{}
	for _, stat := range playerStats {
		statsSummary.WriteString(stat.printVotes() + "| ")
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

func (p *Player) UpdatePlayerPoints(playerName string) {
	data := (*p)[playerName]
	data.Votes++
	(*p)[playerName] = data
}
