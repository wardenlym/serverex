package componet

import (
	"gitlab.com/megatech/serverex/core/ecs"
)

type UserID struct {
	ecs.Componet
	UID int64
}
