package level

import (
	"bytes"
	"errors"

	"interview/raid/demo/pkg/types"
)

type RAID0 struct {
	disks      types.RAID
	stripeSize int
}

func NewRAID0(numDisks, stripeSize int) *RAID0 {
	disks := make(types.RAID, numDisks)
	for i := range numDisks {
		disks[i] = make(types.Disk, 0)
	}

	return &RAID0{
		disks:      disks,
		stripeSize: stripeSize,
	}
}

func (raid0 *RAID0) Write(data []byte) error {
	numDisks := len(raid0.disks)
	if numDisks == 0 {
		return errors.New("invalid disk length")
	}

	pos := 0
	for len(data) > 0 {
		stripe := min(len(data), raid0.stripeSize)
		chunk := data[:stripe]
		data = data[stripe:]

		diskIndex := pos % numDisks
		raid0.disks[diskIndex] = append(raid0.disks[diskIndex], types.DataStripe(chunk))
		pos++
	}
	return nil
}

func (raid0 *RAID0) Read() []byte {
	var buffer bytes.Buffer
	for _, disk := range raid0.disks {
		for _, stripe := range disk {
			buffer.Write(stripe)
		}
	}
	return buffer.Bytes()
}

func (raid0 *RAID0) Clear(diskIndex int) error {
	if diskIndex < 0 || diskIndex >= len(raid0.disks) {
		return errors.New("invalid disk index")
	}

	raid0.disks[diskIndex] = make(types.Disk, 0)
	return nil
}
