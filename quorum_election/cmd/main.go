package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"interview/quorum/election/pkg/hub"
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
		timeout := time.After(5 * time.Second)
		for {
			select {
			case <-ctx.Done():
				return
			case signal := <-signalCh:
				if len(signal) > 0 {
					fmt.Println(signal)
				}
				timeout = time.After(5 * time.Second)
			case <-timeout:
				fmt.Println("No signal received, canceling the quorum.")
				return
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
