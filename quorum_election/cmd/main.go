package main

import (
	"context"
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

	hub := hub.New(memberAmount)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	signalCh := make(chan string)
	go hub.Heartbeat(ctx, signalCh)
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			case signal := <-signalCh:
				fmt.Println(signal)
			}
		}
	}(ctx)

	hub.ElectLeader()

	for {
		var command string
		var id int
		if _, err := fmt.Scanln(&command, &id); err != nil {
			fmt.Println("Failed to read command")
			continue
		}

		switch command {
		case "kill":
			hub.RemoveMember(id)
		default:
			fmt.Println("Unknown command")
		}
	}
}
