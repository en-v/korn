package holder

import (
	"github.com/pkg/errors"
)

func (self *_Holder) Work(id string) error {
	if !self.activated {
		return errors.New("holder is not activated")
	}

	snap, exists := self.duplicates[id]
	if !exists {
		return errors.New("doublet with current key is not found, key = " + id)
	}

	origin, err := self.originById(id)
	if err != nil {
		return errors.Wrap(err, "Get the origin as a target by the key")
	}

	if self.storage != nil {
		err = self.dump(origin)
		if err != nil {
			return errors.Wrap(err, "Dump To Storage")
		}
	}

	err = self.compare(snap, origin)
	if err != nil {
		return errors.Wrap(err, "CheckThing")
	}

	return nil
}
