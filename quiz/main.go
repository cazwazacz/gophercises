package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"gophercises/quiz/question"
	"io"
	"log"
	"os"
	"time"
)

// ParseQuestions reads the csv file and returns a map of Questions
func ParseQuestions(filepath string) ([]question.Question, error) {
	problems, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(problems)

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
	filepath := flag.String("f", "problems.csv", "A filepath to your csv file with problems")
	timeLimit := flag.Int("l", 30, "The time limit for the quiz in seconds")
	flag.Parse()

	score := 0
	questions, err := ParseQuestions(*filepath)
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(os.Stdin)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	for _, question := range questions {
		fmt.Println(question.Title)

		answerCh := make(chan string)
		go func() {
			var answer string
			answer, _ = reader.ReadString('\n')
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println("Time's up!")
			return
		case answer := <-answerCh:
			if question.AnswerCorrect(answer) {
				incrementScore(&score)
			}
		}
	}

	fmt.Println("Your score is", score)
}
