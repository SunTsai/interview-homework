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
	Name   string
	Answer float32
}

func New(name string) *Student {
	return &Student{
		Name: name,
	}
}

func (s *Student) ListeningAndAnswer(wg *sync.WaitGroup, questionCh <-chan types.Question, answererCh chan<- *Student) {
	defer wg.Done()
	for question := range questionCh {
		thinkTime := time.Duration(rand.IntN(3)+1) * time.Second
		time.Sleep(thinkTime)

		ans := utils.CalculateAnswer(question)
		fmt.Printf("Student %s: %d %s %d = %f!\n", s.Name, question.Num0, question.Operator, question.Num1, ans)
		s.Answer = ans
		answererCh <- s
	}
}
