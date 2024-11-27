package main

import (
	"fmt"
	"sync"
	"time"

	"main/pkg/types"
	"main/pkg/usecase/student"
	"main/pkg/usecase/teacher"
)

func main() {
	questionCh := make(chan types.Question)
	winnerCh := make(chan string)
	done := make(chan struct{})

	students := make([]*student.Student, 5)
	for i := range len(students) {
		students[i] = student.New(string('A'+i), questionCh, winnerCh, done)
	}

	wg := &sync.WaitGroup{}
	wg.Add(len(students))
	for _, student := range students {
		go student.Answer(wg)
	}

	teacher := teacher.New()
	fmt.Println("Teacher: Guys, are you ready?")
	time.Sleep(3 * time.Second)
	teacher.AskQuestion(questionCh)

	winner := <-winnerCh
	fmt.Printf("Teacher: %s, you are right!\n", winner)

	close(done)
	wg.Wait()
	for _, student := range students {
		if student.Name != winner {
			fmt.Printf("Student %s: %s, you win.\n", student.Name, winner)
		}
	}
}
