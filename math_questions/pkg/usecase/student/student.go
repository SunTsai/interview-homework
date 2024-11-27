package student

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"

	"main/pkg/types"
	"main/pkg/utils"
)

type Student struct {
	Name       string
	QuestionCh <-chan types.Question
	WinnerCh   chan string
	Done       chan struct{}
}

func New(name string, questionCh <-chan types.Question, winnerCh chan string, done chan struct{}) *Student {
	return &Student{
		Name:       name,
		QuestionCh: questionCh,
		WinnerCh:   winnerCh,
		Done:       done,
	}
}

func (s *Student) Answer(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case question := <-s.QuestionCh:
			thinkTime := time.Duration(rand.IntN(3)+1) * time.Second
			time.Sleep(thinkTime)

			ans := utils.CalculateAnswer(question)
			fmt.Printf("Student %s: %d %s %d = %f!\n", s.Name, question.Num0, question.Operator, question.Num1, ans)
			s.WinnerCh <- s.Name
			return
		case <-s.Done:
			return
		}
	}
}
