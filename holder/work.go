package holder

import (
	"github.com/pkg/errors"
)

func (self *_Holder) Work(id string) error {
	if !self.activated {
		return errors.New("Holder is not activated")
	}

	snap, exists := self.duplicates[id]
	if !exists {
		return errors.New("Doublet with current Id is not found, Id = " + id)
	}

	origin, err := self.originById(id)
	if err != nil {
		return errors.Wrap(err, "Holder Works, Get origin as a target by Id")
	}

	if self.storage != nil {
		err = self.dump(origin)
		if err != nil {
			return errors.Wrap(err, "Holder Works")
		}
	}

	err = self.compare(snap, origin)
	if err != nil {
		return errors.Wrap(err, "CheckThing")
	}

	return nil
}
