package container

import (
	"github.com/pkg/errors"
)

func (this *Container) LookAt(key string) error {
	if !this.activated {
		return errors.New("Container is not activated")
	}

	snap, exists := this.snaps[key]
	if !exists {
		return errors.New("Snapshot with current key is not found, key = " + key)
	}

	origin, err := this.originByKey(key)
	if err != nil {
		return errors.Wrap(err, "Get the origin as a target by the key")
	}

	err = this.compare(snap, origin)
	if err != nil {
		return errors.Wrap(err, "CheckThing")
	}

	return nil
}
