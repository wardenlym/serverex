package turnbase

import "gitlab.com/megatech/serverex/core/svc"

type Scene struct {
	svc.Named
}

type SceneOption struct {
	svc.Named
}

type SceneManager interface {
	Create(*SceneOption) *Scene
	Remove(*Scene)
	GetByID(id string) *Scene
}

////////////////////////////////////////////
