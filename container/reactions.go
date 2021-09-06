package container

import (
	"sync"

	"github.com/en-v/reactor/core"
	"github.com/en-v/reactor/event"
	"github.com/pkg/errors"
)

type Reactions struct {
	mutex    sync.Mutex
	Regulars map[string]*Reaction
	OnAdd    event.Handler
	OnRemove event.Handler
}

type Reaction struct {
	Name    string
	Handler event.Handler
}

func emptyReactions() *Reactions {
	return &Reactions{
		Regulars: make(map[string]*Reaction),
	}
}

func (this *Reactions) add(name string, handler event.Handler) {
	this.mutex.Lock()
	switch name {
	case core.TAG_ADD:
		this.OnAdd = handler

	case core.TAG_REM:
		this.OnRemove = handler

	default:
		this.Regulars[name] = &Reaction{
			Name:    name,
			Handler: handler,
		}
	}
	this.mutex.Unlock()
}

func (this *Reactions) get(name string) (*Reaction, error) {
	reaction, exists := this.Regulars[name]
	if !exists {
		return nil, errors.New("Reaction does not exist, name = " + name)
	}
	return reaction, nil
}
