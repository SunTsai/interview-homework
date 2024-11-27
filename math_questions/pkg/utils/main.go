package utils

import (
	"math"
	"math/rand/v2"

	"main/pkg/types"
)

func RandNumber(min, max int) int {
	return rand.IntN(max-min) + min
}

func RandOperator() string {
	ops := []string{"+", "-", "*", "/"}
	return ops[rand.IntN(len(ops))]
}

func CalculateAnswer(question types.Question) float32 {
	var ans float32

	switch question.Operator {
	case "+":
		ans = float32(question.Num0) + float32(question.Num1)
	case "-":
		ans = float32(question.Num0) - float32(question.Num1)
	case "*":
		ans = float32(question.Num0) * float32(question.Num1)
	case "/":
		if question.Num1 == 0 {
			ans = float32(math.NaN())
		} else {
			ans = float32(question.Num0) / float32(question.Num1)
		}
	default:
		ans = float32(math.NaN())
	}
	return ans
}
