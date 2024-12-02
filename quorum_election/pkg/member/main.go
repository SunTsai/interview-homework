package member

import (
	"fmt"

	membertype "interview/quorum/election/pkg/member/type"
)

type Member struct {
	ID         int
	IsAlive    bool
	MemberType membertype.MemberType
}

func New(id int) *Member {
	fmt.Printf("Member %d: Hi\n", id)
	return &Member{ID: id, IsAlive: true, MemberType: membertype.Follower}
}
