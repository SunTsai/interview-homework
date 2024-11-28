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
	Answer interface{}
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
		if rand.IntN(2) == 0 {
			ans = utils.CalculateAnswer(question.New())
		}

		ansStr := utils.ParseAnswer(ans)
		fmt.Printf("Student %s: %d %s %d = %s!\n", s.Name, q.Num0, q.Operator, q.Num1, ansStr)

		s.Answer = ans
		answererCh <- s
	}
}
