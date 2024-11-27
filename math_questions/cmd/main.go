package main

import (
	"fmt"
	"sync"

	"main/pkg/types"
	"main/pkg/usecase/student"
	"main/pkg/usecase/teacher"
)

func main() {
	questionCh := make(chan types.Question)
	teacher := teacher.New()
	go teacher.AskQuestion(questionCh)

	students := make([]*student.Student, 5)
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
	go teacher.CheckAnswer(answererCh, questionCh, winnerCh)
	wg.Wait()
	close(answererCh)

	winner := <-winnerCh
	for _, student := range students {
		if student.Name != winner {
			fmt.Printf("Student %s: %s, you win.\n", student.Name, winner)
		}
	}
}
