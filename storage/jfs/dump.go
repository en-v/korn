package jfs

import (
	"encoding/json"
	"os"

	"github.com/en-v/korn/inset"
	"github.com/pkg/errors"
)

func (self *JFS) Dump(holderName string, obj inset.InsetInterface) error {

	filename := self.fullFileName(holderName, obj.GetId())

	data, err := json.Marshal(obj)
	if err != nil {
		return errors.Wrap(err, "Marshalling before Dump")
	}

	_, err = os.Stat(filename)
	if os.IsExist(err) {
		err = os.Remove(filename)
		if err != nil {
			return errors.Wrap(err, "Remove existing before Dump")
		}
	}

	err = os.WriteFile(filename, data, 0777)
	if err != nil {
		return errors.Wrap(err, "Dump to file")
	}

	return nil
}
