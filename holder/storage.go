package holder

import (
	"reflect"

	"github.com/en-v/korn/inset"
	"github.com/en-v/korn/storage"
	"github.com/pkg/errors"
)

func (self *_Holder) SetStore(s storage.IStorage) error {
	if self.reftype == nil {
		return errors.New("You have to set reference object before store setting, Holder = " + self.name)
	}

	err := s.Prepare(self.name)
	if err != nil {
		return errors.Wrap(err, "Set Store")
	}
	self.storage = s
	return nil
}

func (self *_Holder) SetRefo(refobj interface{}) error {
	refo, cast := refobj.(inset.InsetInterface)
	if !cast {
		return errors.New("Reference Object cant to be casted to InsetInterface")
	}
	self.reftype = reflect.ValueOf(refo).Elem().Type()
	return nil
}

func (self *_Holder) Restore() error {
	if self.storage != nil {
		items, err := self.storage.Restore(self.name, self.reftype)
		if err != nil {
			return errors.Wrap(err, "Restore on Holder")
		}

		for _, i := range items {
			i.(inset.InsetInterface).Link(self)
		}

		err = self.Capture(items)
		if err != nil {
			return errors.Wrap(err, "Capture Restored Items")
		}
		return nil

	}
	return errors.New("No active storage")
}

func (self *_Holder) dump(obj inset.InsetInterface) error {

	if self.storage == nil {
		return errors.New("No active storage")
	}

	err := self.storage.Dump(self.name, obj)
	if err != nil {
		return errors.Wrap(err, "Dump To Storage")
	}
	return nil
}
