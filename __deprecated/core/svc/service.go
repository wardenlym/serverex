package svc

import (
	"sync"
	"sync/atomic"

	"gitlab.com/megatech/serverex/lib/fn"
	"gitlab.com/megatech/serverex/lib/log"
)

type Runnable interface {
	Run()
	Stop()
	Quit() chan error
	Status() string
}

type Service interface {
	ConstNamed
	Runnable
}

var id_acc IDType

func NewServiceID() IDType {
	id := atomic.AddInt64(&id_acc, 1)
	return id
}

type basicService struct {
	ConstNamed
	Runnable
}

func NewBasicService(name string, loop func(stop chan bool, quit chan error)) Service {
	return &basicService{
		ConstNamed: NewConstNamed(NewServiceID(), name),
		Runnable:   NewBasicRunnable(loop),
	}
}

func NewBasicRunnable(loop func(stop chan bool, quit chan error)) Runnable {
	return &runner{
		stop: make(chan bool),
		quit: make(chan error),
		loop: loop,
		stat: "init",
	}
}

type runner struct {
	sync.Mutex
	stop chan bool
	quit chan error
	loop func(stop chan bool, quit chan error)
	stat string
}

func (p *runner) Run() {
	p.Lock()
	p.stat = "running"
	p.Unlock()

	err := fn.ProtectCall(
		func() {
			p.loop(p.stop, p.quit)
		})

	if err != nil {
		log.Error("panic:", err)
	}

	p.Lock()
	p.stat = "stopped"
	p.Unlock()
}

func (p *runner) Stop() {
	p.Lock()
	defer p.Unlock()
	if p.stop != nil {
		close(p.stop)
		p.stat = "shutting"
		p.stop = nil
	}
}

func (p *runner) Quit() chan error {
	return p.quit
}

func (p *runner) Status() string {
	p.Lock()
	defer p.Unlock()
	return p.stat
}
