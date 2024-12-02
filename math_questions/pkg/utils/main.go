package utils

import (
	"fmt"
	"math"

	"interview/math/questions/pkg/question"
)

func CalculateAnswer(question question.Question) interface{} {
	var ans interface{}

	switch question.Operator {
	case "+":
		ans = question.Num0 + question.Num1
	case "-":
		ans = question.Num0 - question.Num1
	case "*":
		ans = question.Num0 * question.Num1
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

func ParseAnswer(answer interface{}) string {
	switch answer.(type) {
	case int:
		return fmt.Sprintf("%d", answer)
	default:
		return fmt.Sprintf("%.2f", answer)
	}
}
