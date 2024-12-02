package level

type RAID interface {
	Write(data []byte) error
	Read() []byte
	Clear() error
}
