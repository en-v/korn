package korn

import (
	"fmt"

	"github.com/en-v/korn/holder"
	"github.com/en-v/korn/storage"
	"github.com/pkg/errors"
)

type _Engine struct {
	storage storage.IStorage
	holders map[string]holder.IHolder
}

func makeEngine() *_Engine {
	return &_Engine{
		holders: make(map[string]holder.IHolder),
	}
}

func (self *_Engine) Holder(name string) (holder.IHolder, error) {
	obs, exists := self.holders[name]
	if exists {
		return obs, nil
	}

	new := holder.Make(name)

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

	for name, obs := range self.holders {
		err := obs.Activate()
		if err != nil {
			go self.Shutdown()
			return errors.Wrap(err, fmt.Sprintf("holder('%s').Activation", name))
		}
	}
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

func (self *_Engine) Connect(storageName string, connectionString string) error {
	new, err := storage.Make(storageName, connectionString)
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
			return errors.Wrap(err, "Raise on Engine")
		}
	}
	return nil
}

func (self *_Engine) Reset() error {
	panic("not implemented")
}
