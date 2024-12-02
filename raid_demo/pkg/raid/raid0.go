package raid

import (
	"bytes"
	"errors"

	"interview/raid/demo/pkg/types"
)

type RAID0 struct {
	disks      types.RAID
	stripeSize int
}

func NewRAID0(numDisks, stripeSize int) (*RAID0, error) {
	if numDisks < 2 {
		return nil, errors.New("invalid disk numbers")
	}

	disks := make(types.RAID, numDisks)
	for i := range numDisks {
		disks[i] = make(types.Disk, 0)
	}

	return &RAID0{
		disks:      disks,
		stripeSize: stripeSize,
	}, nil
}

func (r *RAID0) Write(data []byte) {
	pos := 0
	numDisks := len(r.disks)

	for len(data) > 0 {
		stripe := min(len(data), r.stripeSize)
		chunk := data[:stripe]
		data = data[stripe:]

		diskIndex := pos % numDisks
		r.disks[diskIndex] = append(r.disks[diskIndex], types.DataStripe(chunk))
		pos++
	}
}
func (raid *RAID0) Read() ([]byte, error) {
	var buffer bytes.Buffer
	for _, disk := range raid.disks {
		for _, stripe := range disk {
			buffer.Write(stripe)
		}
	}
	return buffer.Bytes(), nil
}

func (r *RAID0) Clear(diskIndex int) error {
	if diskIndex < 0 || diskIndex >= len(r.disks) {
		return errors.New("invalid disk index")
	}

	r.disks[diskIndex] = make(types.Disk, 0)
	return nil
}
