package main

import (
	"fmt"
	"log"
	"math/rand/v2"

	"interview/raid/demo/pkg/level"
)

func main() {
	data := []byte("This is a RAID demonstration.")
	fmt.Println("Original data:")
	fmt.Println(string(data))

	raid := level.NewRAID0(3, 4)
	if err := raid.Write(data); err != nil {
		log.Fatal("Failed to write the data: ", err)
	}
	diskIndex := rand.IntN(3)
	if err := raid.Clear(diskIndex); err != nil {
		log.Fatal("Failed to clear the disk: ", err)
	}

	fmt.Println("\nReconstructed data from RAID 0:")
	fmt.Println(string(raid.Read()))
}
