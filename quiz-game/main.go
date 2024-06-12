package main

import (
	"encoding/csv"
	"fmt"
	"os"
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

func askQuestionAndGetResult(quizItemNumber int, quizItem []string) bool {
	question, correctAnswer, userAnswer := quizItem[0], quizItem[1], ""
	fmt.Printf("Question #%d: %s = ", quizItemNumber, question)
	fmt.Scanln(&userAnswer)
	return userAnswer == correctAnswer
}

func main() {
	lines, err := getLinesFromCSV(FILE_NAME)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	score := 0
	numQuestions := len(lines)
	for i, quizItem := range lines {
		if askQuestionAndGetResult(i+1, quizItem) {
			score++
		}
	}
	fmt.Printf("Results: %d/%d correct\n", score, numQuestions)
}
