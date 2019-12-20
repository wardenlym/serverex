package datamap

import (
	"io/ioutil"
	"sync"

	"gitlab.com/megatech/serverex/lib/fswatcher"
	"gitlab.com/megatech/serverex/lib/log"
)

type DataMap struct {
	sync.Mutex
	data    map[int]interface{}
	watcher *fswatcher.Watcher
}

func NewImmutable(parser func() map[int]interface{}) *DataMap {
	return &DataMap{
		Mutex: sync.Mutex{},
		data:  parser(),
	}
}

func New(path string, parser func([]byte) map[int]interface{}) *DataMap {

	b, err := ioutil.ReadFile(path)
	if err != nil {
	}

	m := parser(b)
	if m == nil {
	}

	p := &DataMap{
		Mutex: sync.Mutex{},
		data:  m,
	}

	watcher := fswatcher.New(path, func(name string) {
		p.Lock()
		b, err := ioutil.ReadFile(path)
		if err != nil {
			log.Error(err)
		}

		m := parser(b)
		if m == nil {
			log.Error(m)
		}
		p.data = m
		p.Unlock()
	})
	p.watcher = watcher
	return p
}

func (p *DataMap) Find(k int) (interface{}, bool) {
	p.Lock()
	defer p.Unlock()
	v, ok := p.data[k]
	return v, ok
}

func (p *DataMap) Keys() []int {
	p.Lock()
	defer p.Unlock()
	keys := []int{}
	for k, _ := range p.data {
		keys = append(keys, k)
	}
	return keys
}
