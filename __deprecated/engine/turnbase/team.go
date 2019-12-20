package turnbase

import (
	"time"

	"gitlab.com/megatech/serverex/core/msg"
	"gitlab.com/megatech/serverex/core/svc"
)

type Team struct {
	svc.Service
	msg.Parser
	msg.Dispatcher
	svc.MqClient

	admin TeamInfoAdmin
}

type TeamID int32

type MemberID int64

type MemberInfo struct {
	MemberID
}

type TeamStatus int

const (
	_ TeamStatus = iota
	Preparing
	Confirmed
	Expired
)

type TeamInfo struct {
	TeamID
	Max     int
	Members []MemberInfo
	Created time.Time
	TeamStatus
}

type TeamInfoAdmin struct {
	Teams map[TeamID]TeamInfo
}

func NewTeamInfoAdmin() *TeamInfoAdmin {
	return &TeamInfoAdmin{
		make(map[TeamID]TeamInfo),
	}
}
