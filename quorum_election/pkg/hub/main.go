package hub

import (
	"context"
	"fmt"
	"time"

	"interview/quorum/election/pkg/member"
)

type Hub struct {
	members []*member.Member
}

func New(memberAmount int) *Hub {
	members := make([]*member.Member, memberAmount)
	for i := range memberAmount {
		members[i] = member.New(i)
	}

	return &Hub{members: members}
}

func (h *Hub) Heartbeat(ctx context.Context, signalCh chan string) {
	inactiveMembers := make(map[int]bool)

	for _, sender := range h.members {
		inactiveMembers[sender.ID] = false
		go func(ctx context.Context, sender *member.Member, inactiveMembers map[int]bool) {
			for {
				select {
				case <-ctx.Done():
					return
				default:
					if sender.IsAlive {
						for _, receiver := range h.members {
							if sender.ID != receiver.ID && !receiver.IsAlive {
								if !inactiveMembers[receiver.ID] {
									signalCh <- fmt.Sprintf("Member %d: failed heartbeat with Member %d", sender.ID, receiver.ID)
									inactiveMembers[receiver.ID] = true
								}
							}
						}
					}
					time.Sleep(time.Second)
				}
			}
		}(ctx, sender, inactiveMembers)
	}
}

func (h *Hub) ElectLeader() {
	// Member 0: I want to be leader
	// > Member 2: Accept member 0 to be leader
	// > Member 1: I want to be leader
	// > Member 1: Accept member 0 to be leader

	// If member 自願當 leader
	// 問剩下 member 同不同意
	// 若其中一個 member 票數過半
	// 直接選他當 leader

}

func (h *Hub) RemoveMember(id int) {
	h.members[id].IsAlive = false
	// if id is leader
	// elect new leader
}
