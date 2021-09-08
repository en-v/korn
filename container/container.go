package container

import (
	"github.com/en-v/korn/event"
	"github.com/en-v/korn/snapshot"
	"github.com/pkg/errors"
)

type Mode byte

type _Holder struct {
	activated bool
	name      string
	reactions *Reactions
	origins   map[string]interface{}
	snaps     map[string]*snapshot.Snapshot
	errs      chan error
}

//Make - create a new instance of Container
func Make(name string) *_Holder {
	return &_Holder{
		activated: false,
		name:      name,
		reactions: emptyReactions(),
		errs:      make(chan error),
		origins:   make(map[string]interface{}),
		snaps:     make(map[string]*snapshot.Snapshot),
	}
}

func (this *_Holder) On(name string, handler event.Handler) {
	if handler == nil {
		panic("Handler cannot to be empty")
	}
	this.reactions.add(name, handler)
}

func (this *_Holder) Activate() error {
	if this.reactions.OnAdd == nil || this.reactions.OnRemove == nil {
		return errors.New("No requred reactions found ('add' and 'remove'), you have to add requred reactions")
	}

	if len(this.reactions.Regulars) == 0 {
		return errors.New("No regualr rections found")
	}

	this.activated = true
	return nil
}

func (this *_Holder) Shutdown() {
	this.activated = false
}

func (this *_Holder) Name() string {
	return this.name
}

func (this *_Holder) CatchError() error {
	return <-this.errs
}

func (this *_Holder) All() map[string]interface{} {
	return this.origins
}

func (this *_Holder) Count() int {
	return len(this.origins)
}

func (this *_Holder) Reset() error {
	panic("Method not implememted")
}
