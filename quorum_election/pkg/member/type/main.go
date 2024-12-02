package membertype

type MemberType int

const (
	Follower MemberType = iota
	Candidate
	Leader
)
