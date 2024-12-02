package member

import (
	"fmt"
)

type Member struct {
	ID       int
	IsAlive  bool
	IsLeader bool
}

func New(id int) *Member {
	fmt.Printf("Member %d: Hi\n", id)
	return &Member{ID: id, IsAlive: true, IsLeader: false}
}
