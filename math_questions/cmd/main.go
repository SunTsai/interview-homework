package main

import (
	"fmt"
	"sync"

	"interview/math/questions/pkg/question"
	"interview/math/questions/pkg/student"
	"interview/math/questions/pkg/teacher"
	"interview/math/questions/pkg/utils"
)

func main() {
	questionCh := make(chan question.Question)
	teacher := teacher.New()
	go teacher.AskQuestion(questionCh)

	const studentCount = 5
	students := make([]*student.Student, studentCount)
	for i := range len(students) {
		students[i] = student.New(string('A' + i))
	}

	wg := &sync.WaitGroup{}
	wg.Add(len(students))
	answererCh := make(chan *student.Student)
	for _, student := range students {
		go student.ListeningAndAnswer(wg, questionCh, answererCh)
	}

	winnerCh := make(chan string)
	go teacher.CheckAnswer(studentCount, answererCh, questionCh, winnerCh)
	wg.Wait()
	close(answererCh)

	winner := <-winnerCh
	if len(winner) == 0 {
		ansStr := utils.ParseAnswer(teacher.Answer)
		fmt.Printf("Teacher: Boooo~ Answer is %s.\n", ansStr)
	} else {
		for _, student := range students {
			if student.Name != winner {
				fmt.Printf("Student %s: %s, you win.\n", student.Name, winner)
			}
		}
	}
}
