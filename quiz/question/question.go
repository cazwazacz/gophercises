package question

import (
	"strings"
)

// Question holds the question and answer information
type Question struct {
	Title  string
	Answer string
}

// AnswerCorrect checks if the answer is correct
func (q *Question) AnswerCorrect(attempt string) bool {
	return q.Answer == strings.TrimSpace(attempt)
}