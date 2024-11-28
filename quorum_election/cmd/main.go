package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"main/pkg/hub"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("Please give a member amount number")
	}

	memberAmountArg := args[1]
	memberAmount, err := strconv.Atoi(memberAmountArg)
	if err != nil {
		log.Fatal("Invalid member amount")
	}

	fmt.Printf("Starting quorum with %d members\n", memberAmount)

	hub.New(memberAmount)
}
