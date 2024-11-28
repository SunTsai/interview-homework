package teacher

import (
	"fmt"
	"time"

	"main/pkg/question"
	"main/pkg/student"
	"main/pkg/utils"
)

type Teacher struct {
	Answer   interface{}
	question question.Question
}

func New() *Teacher {
	return &Teacher{}
}

func (t *Teacher) AskQuestion(questionCh chan question.Question) {
	question := question.New()

	fmt.Println("Teacher: Guys, are you ready?")
	time.Sleep(3 * time.Second)
	fmt.Printf("Teacher: %d %s %d = ?\n", question.Num0, question.Operator, question.Num1)

	questionCh <- question
	t.question = question
	t.Answer = utils.CalculateAnswer(question)
}

func (t *Teacher) CheckAnswer(studentAmount int, answererCh <-chan *student.Student, questionCh chan<- question.Question, winnerCh chan string) {
	answered := make(map[string]bool)
	for answerer := range answererCh {
		if answerer.Answer == t.Answer {
			fmt.Printf("Teacher: %s, you are right!\n", answerer.Name)
			close(questionCh)
			winnerCh <- answerer.Name
		} else {
			fmt.Printf("Teacher: %s, you are wrong!\n", answerer.Name)
			answered[answerer.Name] = true
			if len(answered) == studentAmount {
				close(questionCh)
				winnerCh <- ""
			} else {
				questionCh <- t.question
			}
		}
	}
}
