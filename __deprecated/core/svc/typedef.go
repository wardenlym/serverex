package svc

type IDType = int64

type ConstNamed interface {
	ID() IDType
	Name() string
}

type Named interface {
	ConstNamed
	SetID(id string)
	SetName(name string)
}
type Name struct {
	id   IDType
	name string
}

func (p *Name) ID() IDType {
	return p.id
}

func (p *Name) Name() string {
	return p.name
}

func (p *Name) SetID(id IDType) {
	p.id = id
}

func (p *Name) SetName(name string) {
	p.name = name
}

func NewConstNamed(id IDType, name string) ConstNamed {
	return &constName{
		id:   id,
		name: name,
	}
}

type constName struct {
	id   IDType
	name string
}

func (p *constName) ID() IDType {
	return p.id
}

func (p *constName) Name() string {
	return p.name
}
