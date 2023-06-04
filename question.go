package main

import (
	"math/rand"
	"os"
	"strings"
)

type Question struct {
	Question  string
	IsReverse bool
}

func loadQuestionsFromFile(filename string) ([]Question, error) {
	fileContent, err := os.ReadFile(filename)
	
	if err != nil {
		return nil, err
	}
	
	lines := strings.Split(string(fileContent), "\n")
	var questions []Question
	
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine != "" {
			isReverse := trimmedLine[0] == '*'
			
			if isReverse {
				trimmedLine = trimmedLine[1:]
			}
			
			question := Question{
				Question:  trimmedLine,
				IsReverse: isReverse,
			}
			questions = append(questions, question)
		}
	}
	
	return questions, nil
}

func shuffleQuestions(questions []Question) {
	rand.Shuffle(len(questions), func(i, j int) {
		questions[i], questions[j] = questions[j], questions[i]
	})
}
