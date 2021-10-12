package holder

import (
	"strconv"

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

func (self *_Holder) Single() (interface{}, error) {
	if len(self.origins) != 1 {
		return nil, errors.New("Origins count not 1, current = " + strconv.Itoa(len(self.origins)))
	}

	for origin := range self.origins {
		return origin, nil
	}

	return nil, errors.New("Something went wrong, origins count =  " + strconv.Itoa(len(self.origins)))
}

func (self *_Holder) Remove(id string) error {
	origin, exists := self.origins[id]
	if !exists {
		return errors.New("Elements with current key is not found, key = " + id)
	}

	delete(self.duplicates, id)
	delete(self.origins, id)

	if self.activated && self.reactions.OnRemove != nil {
		err := self.reactions.OnRemove(&event.Event{
			Id:     id,
			Origin: origin,
			Name:   event.KIND_REMOVE,
			Kind:   event.KIND_REMOVE,
			Holder: self.name,
		})
		if err != nil {
			return errors.Wrap(err, "Remove")
		}
	}

	if self.storage != nil {
		err := self.storage.Remove(self.name, id)
		if err != nil {
			return errors.Wrap(err, "Remove From Storage")
		}
	}

	return nil
}
