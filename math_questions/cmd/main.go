package main

import (
	"fmt"
	"sync"

	"main/pkg/question"
	"main/pkg/student"
	"main/pkg/teacher"
)

func main() {
	questionCh := make(chan question.Question)
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
	if len(winner) == 0 {
		fmt.Printf("Teacher: Boooo~ Answer is %.2f.\n", teacher.Answer)
	} else {
		for _, student := range students {
			if student.Name != winner {
				fmt.Printf("Student %s: %s, you win.\n", student.Name, winner)
			}
		}
	}
}
