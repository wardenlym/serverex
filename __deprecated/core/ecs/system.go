package ecs

import (
	"time"
)

type TimeDelta = time.Duration

type System interface {
	Update(admin EntityAdmin, delta TimeDelta)
}

type SystemAdmin interface {
	Register(System)
	Update(admin EntityAdmin, delta TimeDelta)
}
