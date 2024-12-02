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

	for _, config := range []struct {
		level      int
		numDisks   int
		stripeSize int
		numMirrors int
	}{
		{0, 3, 4, 0},
		{1, 3, 4, 0},
		{10, 4, 4, 2},
		{5, 3, 4, 0},
		{6, 3, 4, 0},
	} {
		raid, err := initializeRAID(config.level, config.numDisks, config.stripeSize, config.numMirrors)
		if err != nil {
			fmt.Printf("Failed to initialize RAID%d: %v\n", config.level, err)
			continue
		}

		raid.Write(data)

		diskIndex := rand.IntN(config.numDisks)
		if err := raid.Clear(diskIndex); err != nil {
			fmt.Printf("Failed to clear the disk for RAID%d: %v\n", config.level, err)
			continue
		}

		reconstructedData, err := raid.Read()
		if err != nil {
			fmt.Printf("Failed to read RAID%d: %v\n", config.level, err)
			continue
		}
		fmt.Printf("\nReconstructed data from RAID %d\n", config.level)
		fmt.Println(string(reconstructedData))
	}
}

func initializeRAID(n, numDisks, stripeSize int, numMirrors ...int) (raid.RAID, error) {
	var targetRaid raid.RAID
	var err error

	switch n {
	case 0:
		targetRaid, err = raid.NewRAID0(numDisks, stripeSize)
	case 1:
		targetRaid, err = raid.NewRAID1(numDisks, stripeSize)
	case 10:
		targetRaid, err = raid.NewRAID10(numMirrors[0], numDisks, stripeSize)
	case 5:
		targetRaid, err = raid.NewRAID5(3, stripeSize)
	default:
		err = fmt.Errorf("unknown RAID level: %d", n)
	}

	if err != nil {
		return nil, err
	}
	return targetRaid, nil
}
