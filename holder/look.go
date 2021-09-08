package holder

import (
	"github.com/pkg/errors"
)

func (self *_Holder) LookAt(key string) error {
	if !self.activated {
		return errors.New("holder is not activated")
	}

	snap, exists := self.doublets[key]
	if !exists {
		return errors.New("doublet with current key is not found, key = " + key)
	}

	origin, err := self.originByKey(key)
	if err != nil {
		return errors.Wrap(err, "Get the origin as a target by the key")
	}

	err = self.compare(snap, origin)
	if err != nil {
		return errors.Wrap(err, "CheckThing")
	}

	return nil
}
