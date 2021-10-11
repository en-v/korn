package holder

import (
	"github.com/en-v/korn/event"
	"github.com/pkg/errors"
)

type Reactions struct {
	Items    map[string]*Reaction
	OnAdd    event.Handler
	OnRemove event.Handler
	OnUpdate event.Handler
}

type Reaction struct {
	Name    string
	Handler event.Handler
}

func emptyReactions() *Reactions {
	return &Reactions{
		Items: make(map[string]*Reaction),
	}
}

func (self *Reactions) add(name string, handler event.Handler) {

	switch name {
	case event.KIND_ADD:
		self.OnAdd = handler

	case event.KIND_REMOVE:
		self.OnRemove = handler

	case event.KIND_UPDATE:
		self.OnUpdate = handler

	default:
		self.Items[name] = &Reaction{
			Name:    name,
			Handler: handler,
		}
	}
}

func (self *Reactions) get(name string) (*Reaction, error) {
	reaction, exists := self.Items[name]
	if !exists {
		return nil, errors.New("Reaction does not exist, name = " + name)
	}
	return reaction, nil
}
