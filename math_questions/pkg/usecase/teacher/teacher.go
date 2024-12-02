package teacher

import (
	"math/rand/v2"

	"main/pkg/types"
)

type Teacher struct {
}

func New() *Teacher {
	return &Teacher{}
}

func (t *Teacher) CreateQuestion() types.Question {
	num0, num1 := randNumber(0, 100), randNumber(0, 100)
	operator := randOperator()
	return types.Question{Num0: num0, Num1: num1, Operator: operator}
}

func randNumber(min, max int) int {
	return rand.IntN(max-min) + min
}

func randOperator() string {
	ops := []string{"+", "-", "*", "/"}
	return ops[rand.IntN(len(ops))]
}
