package ecs

type entityID uint32

type Entity entityID
type EntityTuple []Entity

type EntityAdmin interface {
	Create() Entity
	Destroy(Entity)
	Select(...ComponetType) EntityTuple
	Get(Entity, ComponetType) interface{}
	AddComponet(Entity, interface{})
	DelComponet(Entity, ComponetType)
}

// NewDefaultEntityAdmin .
func NewDefaultEntityAdmin() EntityAdmin {
	return &admin{}
}

type admin struct {
	id_acc   Entity
	entities []Entity
}

func (a *admin) Create() Entity {
	a.id_acc = a.id_acc + 1
	return a.id_acc
}

func (a *admin) Destroy(Entity) {

}

func (a *admin) Select(t ...ComponetType) EntityTuple {
	if len(t) == 0 {
		tuple := make([]Entity, len(a.entities))
		copy(tuple, a.entities)
		return tuple
	}
	return nil
}

func (a *admin) Get(Entity, ComponetType) interface{} {
	return nil
}

func (a *admin) AddComponet(Entity, interface{}) {
}

func (a *admin) DelComponet(Entity, ComponetType) {
}
