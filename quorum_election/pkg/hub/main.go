package hub

import (
	"main/pkg/member"
)

type Hub struct {
	Members []*member.Member
}

func New(memberAmount int) *Hub {
	members := make([]*member.Member, memberAmount)
	for i := range memberAmount {
		members[i] = member.New(i)
	}

	return &Hub{Members: members}
}
