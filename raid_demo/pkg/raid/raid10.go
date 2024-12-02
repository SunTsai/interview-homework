package raid

import (
	"bytes"
	"errors"
	"math/rand/v2"

	"interview/raid/demo/pkg/types"
)

type RAID10 struct {
	mirrors    []types.RAID
	stripeSize int
}

func NewRAID10(numMirrors, disksPerMirror, stripeSize int) (*RAID10, error) {
	if numMirrors < 1 || disksPerMirror < 2 {
		return nil, errors.New("RAID10 requires at least 1 mirror and 2 disks per mirror")
	}

	mirrors := make([]types.RAID, numMirrors)
	for i := range numMirrors {
		mirrors[i] = make(types.RAID, disksPerMirror)
		for j := range mirrors[i] {
			mirrors[i][j] = make(types.Disk, 0)
		}
	}

	return &RAID10{
		mirrors:    mirrors,
		stripeSize: stripeSize,
	}, nil
}

func (r *RAID10) Write(data []byte) {
	pos := 0
	numMirrors := len(r.mirrors)

	for len(data) > 0 {
		stripe := min(len(data), r.stripeSize)
		chunk := data[:stripe]
		data = data[stripe:]

		mirrorIndex := pos % numMirrors
		for j := range r.mirrors[mirrorIndex] {
			r.mirrors[mirrorIndex][j] = append(r.mirrors[mirrorIndex][j], types.DataStripe(chunk))
		}
		pos++
	}
}
func (r *RAID10) Read() ([]byte, error) {
	var buffer bytes.Buffer
	for _, mirror := range r.mirrors {
		if len(mirror) == 0 {
			continue
		}

		var availableDisk types.Disk
		for _, disk := range mirror {
			if len(disk) > 0 {
				availableDisk = disk
				break
			}
		}

		if availableDisk == nil {
			continue
		}

		for _, stripe := range availableDisk {
			buffer.Write(stripe)
		}
	}

	return buffer.Bytes(), nil
}

func (r *RAID10) Clear(diskIndex int) error {
	mirrorIndex := rand.IntN(len(r.mirrors))
	mirror := r.mirrors[mirrorIndex]
	if diskIndex < 0 || diskIndex >= len(mirror) {
		return errors.New("invalid disk index")
	}

	mirror[diskIndex] = make(types.Disk, 0)
	return nil
}
