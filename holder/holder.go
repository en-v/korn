package holder

import (
	"fmt"

	"github.com/en-v/korn/duplicate"
	"github.com/en-v/korn/event"
	"github.com/en-v/korn/storage"
	"github.com/pkg/errors"
)

type Mode byte

type _Holder struct {
	activated  bool
	name       string
	reactions  *Reactions
	ref        *Reference
	extra      interface{}
	origins    map[string]interface{}
	duplicates map[string]*duplicate.Duplicate
	errch      chan error
	storage    storage.IStorage
	berrs      []error
}

//Make - create a new instance of holder
func Make(name string, referenceObjectNotPointer interface{}) (*_Holder, error) {

	ref, err := makeRef(referenceObjectNotPointer)
	if err != nil {
		return nil, errors.Wrap(err, "Make Reference")
	}

	h := &_Holder{
		ref:        ref,
		activated:  false,
		name:       name,
		reactions:  emptyReactions(),
		errch:      make(chan error),
		origins:    make(map[string]interface{}),
		duplicates: make(map[string]*duplicate.Duplicate),
		berrs:      make([]error, 0),
	}

	err = h.reactions.scanTags(referenceObjectNotPointer, true)
	if err != nil {
		return nil, errors.Wrap(err, "Make Holder")
	}

	return h, nil
}

func (self *_Holder) Bind(name string, handler event.Handler) {
	if handler == nil {
		panic("Handler cannot to be empty")
	}
	err := self.reactions.add(name, handler)
	if err != nil {
		self.berrs = append(self.berrs, err)
	}
}

func (self *_Holder) BindBasic(add event.Handler, remove event.Handler, update event.Handler) {
	self.Bind(event.KIND_ADD, add)
	self.Bind(event.KIND_REMOVE, remove)
	self.Bind(event.KIND_UPDATE, update)
}

func (self *_Holder) CheckBindings() error {
	if len(self.berrs) > 0 {
		s := ""
		for _, e := range self.berrs {
			if s == "" {
				s = e.Error()
			} else {
				s = s + ", " + e.Error()
			}
		}
		self.berrs = make([]error, 0)
		return errors.New("Check Bindings: " + s)
	}
	return nil
}

func (self *_Holder) Activate() error {
	add, remove, update := self.reactions.OnAdd == nil, self.reactions.OnRemove == nil, self.reactions.OnUpdate == nil

	if add || remove || update {
		str := fmt.Sprintf("add %t, remove %t, update %t", add, remove, update)
		return errors.New("No requred reactions found (method is nil: " + str + "), you have to add requred reactions")
	}

	err := self.reactions.checkUnbounds()
	if err != nil {
		return errors.Wrap(err, "Check For Unbound Tags")
	}

	self.activated = true
	return nil
}

func (self *_Holder) SetExtra(extra interface{}) {
	self.extra = extra
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
	self.origins = make(map[string]interface{})
	self.duplicates = make(map[string]*duplicate.Duplicate)
	if self.storage != nil {
		return self.storage.Reset()
	}
	return nil
}
