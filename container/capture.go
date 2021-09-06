package container

import (
	"reflect"

	"github.com/en-v/reactor/core"
	"github.com/en-v/reactor/event"
	"github.com/en-v/reactor/snapshot"
	"github.com/pkg/errors"
)

func (this *Container) Capture(targets interface{}) error {

	value := reflect.ValueOf(targets)

	switch value.Kind() {
	case reflect.Struct:
		return this.captureStruct(value)
	case reflect.Ptr:
		return this.captureStruct(value)
	case reflect.Map:
		return this.captureMap(value)

	default:
		return errors.New("Targets kind must be a map or struct, current = " + value.Kind().String())
	}
}

func (this *Container) captureStruct(value reflect.Value) (err error) {

	key := value.Interface().(IJet).Key()
	_, exists := this.origins[key]
	if exists {
		return errors.New("Element is alredy captured, key = " + key)
	}

	this.origins[key] = value.Interface()
	jet := this.origins[key].(IJet)
	jet.Observe(this)
	this.snaps[jet.Key()], err = snapshot.Make(jet, this.name)

	if this.activated && this.reactions.OnAdd != nil {
		this.reactions.OnAdd(&event.Event{
			Key:       key,
			Origin:    this.origins[key],
			Name:      "*",
			Kind:      core.TAG_ADD,
			Container: this.name,
		})
	}

	if err != nil {
		return errors.Wrap(err, "Single Capturing Error")
	}
	return nil
}

func (this *Container) captureMap(value reflect.Value) (err error) {
	for _, key := range value.MapKeys() {
		if !value.MapIndex(key).IsNil() {
			item := value.MapIndex(key).Interface().(IJet)
			_, exists := this.origins[item.Key()]
			if exists {
				return errors.New("Element is alredy captured, key = " + item.Key())
			}
		}
	}

	for _, key := range value.MapKeys() {
		if !value.MapIndex(key).IsNil() {
			item := value.MapIndex(key).Interface()
			jet := item.(IJet)
			jet.Observe(this)
			this.origins[jet.Key()] = item
			this.snaps[jet.Key()], err = snapshot.Make(jet, this.name)
			if err != nil {
				return errors.Wrap(err, "Map Captured")
			}
			if this.activated && this.reactions.OnAdd != nil {
				this.reactions.OnAdd(&event.Event{
					Key:       jet.Key(),
					Origin:    value.Interface(),
					Name:      "*",
					Kind:      core.TAG_ADD,
					Container: this.name,
				})
			}
		}
	}

	return nil
}
