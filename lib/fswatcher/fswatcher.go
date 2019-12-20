package fswatcher

import (
	"sync"

	"github.com/fsnotify/fsnotify"
	"gitlab.com/megatech/serverex/lib/log"
)

type Watcher struct {
	stop chan struct{}
	*sync.Mutex
}

func New(path string, onModify func(name string)) *Watcher {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Error(err)
		return nil
	}

	stop := make(chan struct{})

	go func(quit chan struct{}) {
		watcher.Add(path)
		defer watcher.Close()
		for {
			select {
			case <-quit:
				return
			case event := <-watcher.Events:
				if event.Op&fsnotify.Write == fsnotify.Write {
					onModify(event.Name)
				}
			case err := <-watcher.Errors:
				log.Error(err)
			}
		}
	}(stop)

	return &Watcher{
		stop:  stop,
		Mutex: &sync.Mutex{},
	}
}

func (p *Watcher) Stop() {
	p.Lock()
	defer p.Unlock()
	if p.stop != nil {
		close(p.stop)
		p.stop = nil
	}
}
