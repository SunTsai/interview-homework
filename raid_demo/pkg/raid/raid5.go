package raid

import (
	"bytes"
	"errors"

	"interview/raid/demo/pkg/types"
)

type RAID5 struct {
	disks      types.RAID
	stripeSize int
}

func NewRAID5(numDisks, stripeSize int) (*RAID5, error) {
	if numDisks < 3 {
		return nil, errors.New("invalid disk numbers")
	}

	disks := make(types.RAID, numDisks)
	for i := range numDisks {
		disks[i] = make(types.Disk, 0)
	}

	return &RAID5{
		disks:      disks,
		stripeSize: stripeSize,
	}, nil
}

func (r *RAID5) Write(data []byte) {
	pos := 0
	numDisks := len(r.disks)

	for len(data) > 0 {
		stripe := min(len(data), r.stripeSize)
		chunk := data[:stripe]
		data = data[stripe:]

		parityIndex := pos % numDisks
		for i := range numDisks {
			if i == parityIndex {
				continue
			}
			r.disks[i] = append(r.disks[i], types.DataStripe(chunk))
		}

		parity := calculateParity(r.disks, pos, chunk)
		r.disks[parityIndex] = append(r.disks[parityIndex], types.DataStripe(parity))

		pos++
	}
}

func (r *RAID5) Read() ([]byte, error) {
	var buffer bytes.Buffer
	numDisks := len(r.disks)

	for i := 0; i < len(r.disks[0]); i++ {
		var stripe []byte
		parityIndex := i % numDisks
		stripeRecovered := false

		for j := 0; j < numDisks; j++ {
			if j == parityIndex {
				continue
			}

			if i < len(r.disks[j]) {
				stripe = append(stripe, r.disks[j][i]...)
			} else {
				if !stripeRecovered {
					parity := r.disks[parityIndex][i]
					stripe = recoverStripe(stripe, parity)
					stripeRecovered = true
				}
			}
		}

		if len(stripe) == 0 {
			return nil, errors.New("read error: unable to recover lost stripe")
		}

		buffer.Write(stripe)
	}

	return buffer.Bytes(), nil
}

func (r *RAID5) Clear(diskIndex int) error {
	if diskIndex < 0 || diskIndex >= len(r.disks) {
		return errors.New("invalid disk index")
	}

	r.disks[diskIndex] = make(types.Disk, 0)
	return nil
}

func calculateParity(disks types.RAID, pos int, chunk []byte) []byte {
	parity := make([]byte, len(chunk))
	for i := range len(disks) {
		if i != pos%len(disks) {
			for j := range len(chunk) {
				parity[j] ^= chunk[j]
			}
		}
	}
	return parity
}

func recoverStripe(stripe []byte, parity []byte) []byte {
	for i := 0; i < len(parity); i++ {
		stripe[i] ^= parity[i]
	}
	return stripe
}
