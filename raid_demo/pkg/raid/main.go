package raid

type RAID interface {
	Write(data []byte)
	Read() ([]byte, error)
	Clear(diskIndex int) error
}
