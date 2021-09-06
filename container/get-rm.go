package container

import (
	"github.com/en-v/reactor/core"
	"github.com/en-v/reactor/event"
	"github.com/pkg/errors"
)

func (this *Container) Get(key string) interface{} {
	target, exists := this.origins[key]
	if exists {
		return target
	}
	return nil
}

func (this *Container) Remove(key string) error {
	origin, exists := this.origins[key]
	if !exists {
		return errors.New("Elements with current key is not found, key = " + key)
	}

	if this.activated && this.reactions.OnRemove != nil {
		this.reactions.OnRemove(&event.Event{
			Key:    key,
			Origin: origin,
			Name:   "*",
			Kind: core.TAG_REM,
			Container: this.name,
		})
	}
	delete(this.snaps, key)
	delete(this.origins, key)
	return nil
}
