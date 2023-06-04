package types

import (
	"math/rand"
	"os"
	"strings"
)

type Question struct {
	Text      string
	IsReverse bool
}

type Questions []Question

func LoadQuestionsFromFile(filename string) (Questions, error) {
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	
	lines := strings.Split(string(fileContent), "\n")
	questions := make(Questions, 0)
	
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine != "" {
			isReverse := trimmedLine[0] == '*'
			
			if isReverse {
				trimmedLine = trimmedLine[1:]
			}
			
			question := Question{
				Text:      trimmedLine,
				IsReverse: isReverse,
			}
			questions = append(questions, question)
		}
	}
	
	return questions, nil
}

func ShuffleQuestions(questions Questions) {
	rand.Shuffle(len(questions), func(i, j int) {
		questions[i], questions[j] = questions[j], questions[i]
	})
}
