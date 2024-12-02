package main

import (
	"fmt"

	"main/pkg/usecase/student"
	"main/pkg/usecase/teacher"
)

func main() {
	teacher := teacher.New()
	fmt.Println("Teacher: Guys, are you ready?")
	question := teacher.CreateQuestion()
	fmt.Printf("Teacher: %d %s %d = ?\n", question.Num0, question.Operator, question.Num1)

	student := student.New()
	ans := student.Answer(question)
	fmt.Println(ans)
}
