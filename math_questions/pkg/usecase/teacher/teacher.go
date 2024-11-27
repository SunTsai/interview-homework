package teacher

import (
	"fmt"
	"time"

	"main/pkg/types"
	"main/pkg/usecase/student"
	"main/pkg/utils"
)

type Teacher struct {
	question types.Question
	answer   float32
}

func New() *Teacher {
	return &Teacher{}
}

func (t *Teacher) AskQuestion(questionCh chan types.Question) {
	num0, num1 := utils.RandNumber(0, 100), utils.RandNumber(0, 100)
	operator := utils.RandOperator()
	for num1 == 0 && operator == "/" {
		num1 = utils.RandNumber(0, 100)
	}

	fmt.Println("Teacher: Guys, are you ready?")
	time.Sleep(3 * time.Second)
	fmt.Printf("Teacher: %d %s %d = ?\n", num0, operator, num1)

	question := types.Question{Num0: num0, Num1: num1, Operator: operator}
	questionCh <- question
	t.question = question
	t.answer = utils.CalculateAnswer(question)
}

func (t *Teacher) CheckAnswer(answererCh <-chan *student.Student, questionCh chan<- types.Question, winnerCh chan string) {
	for answerer := range answererCh {
		if answerer.Answer == t.answer {
			fmt.Printf("Teacher: %s, you are right!\n", answerer.Name)
			close(questionCh)
			winnerCh <- answerer.Name
		} else {
			questionCh <- t.question
		}
	}
}
