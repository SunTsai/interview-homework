package student

import (
	"math"

	"main/pkg/types"
)

type Student struct {
}

func New() *Student {
	return &Student{}
}

func (s *Student) Answer(q types.Question) float32 {
	if q.Num1 == 0 && q.Operator == "/" {
		return float32(math.NaN())
	}

	switch q.Operator {
	case "+":
		return float32(q.Num0) + float32(q.Num1)
	case "-":
		return float32(q.Num0) - float32(q.Num1)
	case "*":
		return float32(q.Num0) * float32(q.Num1)
	case "/":
		return float32(q.Num0) / float32(q.Num1)
	default:
		return 0
	}
}
