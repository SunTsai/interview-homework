package teacher

import (
	"fmt"
	"math/rand/v2"

	"main/pkg/types"
	"main/pkg/utils"
)

type Teacher struct {
	answer float32
}

func New() *Teacher {
	return &Teacher{}
}

func (t *Teacher) AskQuestion(questionCh chan types.Question) {
	num0, num1 := randNumber(0, 100), randNumber(0, 100)
	operator := randOperator()
	// TODO: if num1 == 0 and op == "/"
	fmt.Printf("Teacher: %d %s %d = ?\n", num0, operator, num1)

	question := types.Question{Num0: num0, Num1: num1, Operator: operator}
	questionCh <- question
	t.answer = utils.CalculateAnswer(question)
}

func randNumber(min, max int) int {
	return rand.IntN(max-min) + min
}

func randOperator() string {
	ops := []string{"+", "-", "*", "/"}
	return ops[rand.IntN(len(ops))]
}
