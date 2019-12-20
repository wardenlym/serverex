package system

import "gitlab.com/megatech/serverex/core/ecs"

type WorldSystem struct {
	ecs.System
	delayed []func(ecs.TimeDelta)
}

func NewWorldSystem() WorldSystem {
	return WorldSystem{
		delayed: []func(ecs.TimeDelta){},
	}
}

func (s WorldSystem) Update(ea ecs.EntityAdmin, delta ecs.TimeDelta) {
	s.delayedUpdate(delta)
}

func (s WorldSystem) delayedUpdate(delta ecs.TimeDelta) {
	for _, f := range s.delayed {
		f(delta)
	}
	s.delayed = []func(ecs.TimeDelta){}
}

func (s WorldSystem) delayToNextFrame(f func(ecs.TimeDelta)) {
	s.delayed = append(s.delayed, f)
}
