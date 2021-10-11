package korn

import (
	"fmt"

	"github.com/en-v/korn/holder"
	"github.com/en-v/korn/storage"
	"github.com/en-v/log"
	"github.com/pkg/errors"
)

//IEngine - reactivity framework for Go
type IEngine interface {
	//Active - activate the kor instance
	Activate() error
	//Shutdown - deactivate the kor instance
	Shutdown()
	//Holder - get an holder, if a holder is not found than will make a new named holder
	//As weel ypu have to set reference object (struct empty instance, not pointer)
	Holder(string, interface{}) (holder.IHolder, error)
	//Connect - connect to storage (JSON-files or MongoDB)
	//path or connection string
	Connect(string) error
	//Restore - restore data from storage into memory
	Restore() error
	//Reset - reset all data from storage and memory
	Reset() error
}

type _Engine struct {
	name    string
	storage storage.IStorage
	holders map[string]holder.IHolder
}

func makeEngine(name string) *_Engine {
	return &_Engine{
		name:    name,
		holders: make(map[string]holder.IHolder),
	}
}

func (self *_Engine) Holder(name string, referenceObjectNotPointer interface{}) (holder.IHolder, error) {
	obs, exists := self.holders[name]
	if exists {
		return obs, nil
	}

	new, err := holder.Make(name, referenceObjectNotPointer)
	if err != nil {
		return nil, errors.Wrap(err, "Make Holder")
	}

	if self.storage != nil {
		err := new.SetStore(self.storage)
		if err != nil {
			return nil, errors.Wrap(err, "Holder (make), name = "+name)
		}
	}

	self.holders[name] = new
	return self.holders[name], nil
}

func (self *_Engine) Activate() error {
	if len(self.holders) == 0 {
		return errors.New("No holders found")
	}
	log.Trace("1")
	for name, holder := range self.holders {
		log.Trace(name)
		err := holder.Activate()
		log.Trace("3")
		if err != nil {
			go self.Shutdown()
			return errors.Wrap(err, fmt.Sprintf("holder('%s').Activation", name))
		} else {
			log.Debugw("Holder activated", "Name", name)
		}
	}
	log.Trace("2")
	return nil
}

func (self *_Engine) Shutdown() {
	for _, obs := range self.holders {
		obs.Shutdown()
	}

	if self.storage != nil {
		self.storage.Shutdown()
	}
}

func (self *_Engine) Connect(connectionString string) error {
	new, err := storage.Make(self.name, connectionString)
	if err != nil {
		return errors.Wrap(err, "Connect to the storage")
	}

	self.storage = new
	for h := range self.holders {
		err := self.holders[h].SetStore(self.storage)
		if err != nil {
			return errors.Wrap(err, "Connect holder, name = "+self.holders[h].Name())
		}
	}
	return nil
}

func (self *_Engine) Restore() error {

	if self.storage == nil {
		return errors.New("No active storage")
	}

	for h := range self.holders {
		err := self.holders[h].Restore()
		if err != nil {
			return errors.Wrap(err, "Restore on Engine")
		}
	}
	return nil
}

func (self *_Engine) Reset() error {
	panic("not implemented")
}
