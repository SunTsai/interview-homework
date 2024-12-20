package main

import (
	"fmt"
	"math/rand/v2"

	"interview/raid/demo/pkg/raid"
)

func main() {
	data := []byte("This is a RAID demonstration.")
	fmt.Println("Original data:")
	fmt.Println(string(data))

	numDisks := 4
	stripeSize := 4
	for _, n := range []int{0, 1, 10, 5, 6} {
		raid, err := initializeRAID(n, numDisks, stripeSize)
		if err != nil {
			fmt.Printf("Failed to initialize RAID%d: %v\n", n, err)
			continue
		}

		raid.Write(data)

		diskIndex := rand.IntN(numDisks)
		if err := raid.Clear(diskIndex); err != nil {
			fmt.Printf("Failed to clear the disk for RAID%d: %v\n", n, err)
			continue
		}

		reconstructedData, err := raid.Read()
		if err != nil {
			fmt.Printf("Failed to read RAID%d: %v\n", n, err)
			continue
		}
		fmt.Printf("\nReconstructed data from RAID %d\n", n)
		fmt.Println(string(reconstructedData))
	}
}

func initializeRAID(n, numDisks, stripeSize int) (raid.RAID, error) {
	var targetRaid raid.RAID
	var err error

	switch n {
	case 0:
		targetRaid, err = raid.NewRAID0(numDisks, stripeSize)
	case 1:
		targetRaid, err = raid.NewRAID1(numDisks, stripeSize)
	case 10:
		targetRaid, err = raid.NewRAID10(numDisks/2, numDisks, stripeSize)
	default:
		err = fmt.Errorf("unknown RAID level: %d", n)
	}

	if err != nil {
		return nil, err
	}
	return targetRaid, nil
}
