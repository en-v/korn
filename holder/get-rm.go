package holder

import (
	"github.com/en-v/korn/event"
	"github.com/pkg/errors"
)

func (self *_Holder) Get(key string) interface{} {
	origin, exists := self.origins[key]
	if exists {
		return origin
	}
	return nil
}

func (self *_Holder) Remove(key string) error {
	origin, exists := self.origins[key]
	if !exists {
		return errors.New("Elements with current key is not found, key = " + key)
	}

	if self.activated && self.reactions.OnRemove != nil {
		self.reactions.OnRemove(&event.Event{
			Key:    key,
			Origin: origin,
			Name:   event.KIND_REMOVE,
			Kind:   event.KIND_REMOVE,
			Holder: self.name,
		})
	}

	delete(self.doublets, key)
	delete(self.origins, key)

	return nil
}
