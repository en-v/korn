package holder

import (
	"reflect"

	"github.com/en-v/korn/doublet"
	"github.com/en-v/korn/event"
	"github.com/pkg/errors"
)

func (self *_Holder) Capture(targets interface{}) error {

	value := reflect.ValueOf(targets)

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

	key := value.Interface().(iInset).Key()
	_, exists := self.origins[key]
	if exists {
		return errors.New("Element is alredy captured, key = " + key)
	}

	self.origins[key] = value.Interface()
	jet := self.origins[key].(iInset)
	jet.Link(self)
	self.doublets[jet.Key()], err = doublet.Make(jet, self.name)

	if self.activated && self.reactions.OnAdd != nil {
		self.reactions.OnAdd(&event.Event{
			Key:    key,
			Origin: self.origins[key],
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
			item := value.MapIndex(key).Interface().(iInset)
			_, exists := self.origins[item.Key()]
			if exists {
				return errors.New("Element is alredy captured, key = " + item.Key())
			}
		}
	}

	for _, keyValue := range value.MapKeys() {
		if !value.MapIndex(keyValue).IsNil() {
			item := value.MapIndex(keyValue).Interface()
			ins := item.(iInset)
			key := ins.Key()

			ins.Link(self)

			self.origins[key] = item
			self.doublets[key], err = doublet.Make(ins, self.name)
			if err != nil {
				return errors.Wrap(err, "Map Captured")
			}

			if self.activated && self.reactions.OnAdd != nil {
				self.reactions.OnAdd(&event.Event{
					Key:    key,
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
