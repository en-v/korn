package holder

import (
	"github.com/en-v/kor/doublet"
	"github.com/en-v/kor/event"
	"github.com/pkg/errors"
)

type Mode byte

type _Holder struct {
	activated bool
	name      string
	reactions *Reactions
	origins   map[string]interface{}
	doublets  map[string]*doublet.Doublet
	errch     chan error
}

//Make - create a new instance of holder
func Make(name string) *_Holder {
	return &_Holder{
		activated: false,
		name:      name,
		reactions: emptyReactions(),
		errch:     make(chan error),
		origins:   make(map[string]interface{}),
		doublets:  make(map[string]*doublet.Doublet),
	}
}

func (self *_Holder) On(name string, handler event.Handler) {
	if handler == nil {
		panic("Handler cannot to be empty")
	}
	self.reactions.add(name, handler)
}

func (self *_Holder) Activate() error {
	if self.reactions.OnAdd == nil || self.reactions.OnRemove == nil {
		return errors.New("No requred reactions found ('add' and 'remove'), you have to add requred reactions")
	}

	if len(self.reactions.Items) == 0 {
		return errors.New("No regualr rections found")
	}

	self.activated = true
	return nil
}

func (self *_Holder) Shutdown() {
	self.activated = false
}

func (self *_Holder) Name() string {
	return self.name
}

func (self *_Holder) CatchError() error {
	return <-self.errch
}

func (self *_Holder) All() map[string]interface{} {
	return self.origins
}

func (self *_Holder) Count() int {
	return len(self.origins)
}

func (self *_Holder) Reset() error {
	panic("Method not implememted")
}
