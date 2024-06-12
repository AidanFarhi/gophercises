package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

const FILE_NAME = "problems.csv"

func getLinesFromCSV(fileName string) ([][]string, error) {
	file, err := os.Open(FILE_NAME)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return lines, err
}

func askQuestionAndGetResult(quizItemNumber int, quizItem []string, answerChannel chan<- bool) {
	question, correctAnswer, userAnswer := quizItem[0], quizItem[1], ""
	fmt.Printf("Question #%d: %s = ", quizItemNumber, question)
	fmt.Scanln(&userAnswer)
	answerChannel <- userAnswer == correctAnswer
}

func main() {
	lines, err := getLinesFromCSV(FILE_NAME)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	score := 0
	numQuestions := len(lines)
	timer := time.NewTimer(time.Duration(2) * time.Second)
	answerChannel := make(chan bool)
	for i, quizItem := range lines {
		go askQuestionAndGetResult(i+1, quizItem, answerChannel)
		select {
		case <-timer.C:
			fmt.Printf("\nResults: %d/%d correct\n", score, numQuestions)
			return
		case result := <-answerChannel:
			if result {
				score++
			}
		}
	}
	fmt.Printf("\nResults: %d/%d correct\n", score, numQuestions)
}
