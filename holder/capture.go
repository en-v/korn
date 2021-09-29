package holder

import (
	"reflect"

	"github.com/en-v/korn/duplicate"
	"github.com/en-v/korn/event"
	"github.com/en-v/korn/inset"
	"github.com/en-v/log"
	"github.com/pkg/errors"
)

func (self *_Holder) Capture(t interface{}) error {

	value := reflect.ValueOf(t)

	switch value.Kind() {
	case reflect.Struct:
		return self.captureStruct(value)
	case reflect.Ptr:
		return self.captureStruct(value)
	case reflect.Map:
		return self.captureMap(value)

	default:
		return errors.New("Targets kind must be a map or struct, current = " + value.Kind().String())
	}
}

func (self *_Holder) captureStruct(value reflect.Value) (err error) {

	id := value.Interface().(inset.InsetInterface).GetId()
	_, exists := self.origins[id]
	if exists {
		return errors.New("Element is alredy captured, key = " + id)
	}

	self.origins[id] = value.Interface()
	ins := self.origins[id].(inset.InsetInterface)
	ins.Link(self)
	if self.storage != nil {
		err = self.dump(ins)
		if err != nil {
			return errors.Wrap(err, "Capture Struct, Dump, Id = "+ins.GetId())
		}
	}
	self.duplicates[ins.GetId()], err = duplicate.Make(ins, self.name)

	if self.activated && self.reactions.OnAdd != nil {
		self.reactions.OnAdd(&event.Event{
			Id:     id,
			Origin: self.origins[id],
			Name:   "*",
			Kind:   event.KIND_ADD,
			Holder: self.name,
		})
	}

	if err != nil {
		return errors.Wrap(err, "Single Capturing Error")
	}
	return nil
}

func (self *_Holder) captureMap(value reflect.Value) (err error) {
	for _, key := range value.MapKeys() {
		if !value.MapIndex(key).IsNil() {
			item := value.MapIndex(key).Interface().(inset.InsetInterface)
			_, exists := self.origins[item.GetId()]
			if exists {
				return errors.New("Element is alredy captured, key = " + item.GetId())
			}
		}
	}

	for _, keyValue := range value.MapKeys() {
		if !value.MapIndex(keyValue).IsNil() {
			item := value.MapIndex(keyValue).Interface()
			ins := item.(inset.InsetInterface)
			id := ins.GetId()

			ins.Link(self)
			if self.storage != nil {
				err = self.dump(ins)
				if err != nil {
					return errors.Wrap(err, "Capture Struct, Dump, Id = "+id)
				}
			}

			log.Trace(id)
			self.origins[id] = item
			self.duplicates[id], err = duplicate.Make(ins, self.name)
			if err != nil {
				return errors.Wrap(err, "Map Captured")
			}

			if self.activated && self.reactions.OnAdd != nil {
				self.reactions.OnAdd(&event.Event{
					Id:     id,
					Origin: value.Interface(),
					Name:   event.KIND_ADD,
					Kind:   event.KIND_ADD,
					Holder: self.name,
				})
			}
		}
	}

	return nil
}
