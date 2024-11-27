package student

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"

	"main/pkg/question"
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

func (s *Student) ListeningAndAnswer(wg *sync.WaitGroup, questionCh <-chan question.Question, answererCh chan<- *Student) {
	defer wg.Done()
	for q := range questionCh {
		thinkTime := time.Duration(rand.IntN(3)+1) * time.Second
		time.Sleep(thinkTime)

		ans := utils.CalculateAnswer(q)
		if rand.IntN(10)%2 == 0 {
			ans = utils.CalculateAnswer(question.New())
		}

		fmt.Printf("Student %s: %d %s %d = %.2f!\n", s.Name, q.Num0, q.Operator, q.Num1, ans)
		s.Answer = ans
		answererCh <- s
	}
}
