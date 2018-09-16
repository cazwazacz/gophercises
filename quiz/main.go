package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"gophercises/quiz/question"
)

// ParseQuestions reads the csv file and returns a map of Questions
func ParseQuestions(filepath string) ([]question.Question, error) {
	problems, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(bufio.NewReader(problems))

	var questions []question.Question
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		questions = append(questions, question.Question{
			Title:  line[0],
			Answer: line[1],
		})
	}

	return questions, nil
}

func incrementScore(score *int) {
	*score++
}

func main() {
	score := 0
	questions, err := ParseQuestions("problems.csv")
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(os.Stdin)

	var answer string

	for _, question := range questions {
		fmt.Println(question.Title)

		answer, _ = reader.ReadString('\n')

		if question.AnswerCorrect(answer) == true {
			incrementScore(&score)
		}
	}

	fmt.Println("Your score is", score)
}
