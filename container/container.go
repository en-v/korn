package container

import (
	"github.com/en-v/reactor/event"
	"github.com/en-v/reactor/snapshot"
	"github.com/pkg/errors"
)

type ContainerMode byte

type Container struct {
	activated bool
	name      string
	reactions *Reactions
	origins   map[string]interface{}
	snaps     map[string]*snapshot.Snapshot
	errs      chan error
}

//Make - create a new instance of Container
func Make(name string) *Container {
	return &Container{
		activated: false,
		name:      name,
		reactions: emptyReactions(),
		errs:      make(chan error),
		origins:   make(map[string]interface{}),
		snaps:     make(map[string]*snapshot.Snapshot),
	}
}

func (this *Container) On(name string, handler event.Handler) {
	if handler == nil {
		panic("Handler cannot to be empty")
	}
	this.reactions.add(name, handler)
}

func (this *Container) Activate() error {
	if this.reactions.OnAdd == nil || this.reactions.OnRemove == nil {
		return errors.New("No requred reactions found ('add' and 'remove'), you have to add requred reactions")
	}

	if len(this.reactions.Regulars) == 0 {
		return errors.New("No regualr rections found")
	}

	this.activated = true
	return nil
}

func (this *Container) Shutdown() {
	this.activated = false
}

func (this *Container) Name() string {
	return this.name
}

func (this *Container) CatchError() error {
	return <-this.errs
}

func (this *Container) All() map[string]interface{} {
	return this.origins
}
