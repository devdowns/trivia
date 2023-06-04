package types

import (
	"fmt"
	"strings"
)

type PlayerStats struct {
	Name      string
	Points    int
	Questions []string
	Votes     int
	Alias     string
	Choices   string
}

func (p *PlayerStats) printStats() string {
	var stats strings.Builder
	
	stats.WriteString(fmt.Sprintf("%s -> ", p.Name))
	stats.WriteString(fmt.Sprintf("%d pts ", p.Points))
	
	for _, question := range p.Questions {
		stats.WriteString(question + "\n")
	}
	
	return stats.String()
}
