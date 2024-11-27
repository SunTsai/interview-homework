package question

import (
	"math/rand/v2"
)

type Question struct {
	Num0     int
	Num1     int
	Operator string
}

func New() Question {
	num0, num1 := randNumber(0, 100), randNumber(0, 100)
	operator := randOperator()
	for num1 == 0 && operator == "/" {
		num1 = randNumber(0, 100)
	}
	return Question{Num0: num0, Num1: num1, Operator: operator}
}

func randNumber(min, max int) int {
	return rand.IntN(max-min) + min
}

func randOperator() string {
	ops := []string{"+", "-", "*", "/"}
	return ops[rand.IntN(len(ops))]
}
