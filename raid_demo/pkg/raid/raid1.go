package raid

import (
	"bytes"
	"errors"

	"interview/raid/demo/pkg/types"
)

type RAID1 struct {
	disks      types.RAID
	stripeSize int
}

func NewRAID1(numDisks, stripeSize int) (*RAID1, error) {
	if numDisks < 2 {
		return nil, errors.New("invalid disk numbers")
	}

	disks := make(types.RAID, numDisks)
	for i := range numDisks {
		disks[i] = make(types.Disk, 0)
	}

	return &RAID1{
		disks:      disks,
		stripeSize: stripeSize,
	}, nil
}

func (r *RAID1) Write(data []byte) {
	for len(data) > 0 {
		stripe := min(len(data), r.stripeSize)
		chunk := data[:stripe]
		data = data[stripe:]

		for i := range r.disks {
			r.disks[i] = append(r.disks[i], types.DataStripe(chunk))
		}
	}
}

func (r *RAID1) Read() ([]byte, error) {
	for _, disk := range r.disks {
		if len(disk) > 0 {
			var buffer bytes.Buffer
			for _, stripe := range disk {
				buffer.Write(stripe)
			}
			return buffer.Bytes(), nil
		}
	}
	return nil, errors.New("all disks are empty")
}

func (r *RAID1) Clear(diskIndex int) error {
	if diskIndex < 0 || diskIndex >= len(r.disks) {
		return errors.New("invalid disk index")
	}

	r.disks[diskIndex] = make(types.Disk, 0)
	return nil
}
