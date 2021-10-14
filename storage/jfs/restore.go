package jfs

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"strings"

	"github.com/en-v/korn/inset"
	"github.com/pkg/errors"
)

func (self *JFS) Restore(folder string, reft reflect.Type) (map[string]interface{}, error) {
	path := self.path + "/" + folder

	list, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, errors.Wrap(err, "JFS Restore")
	}

	res := make(map[string]interface{}, len(list))

	for _, f := range list {

		if !strings.Contains(f.Name(), FILE_EXTENSION) {
			continue
		}

		filePath := path + "/" + f.Name()

		data, err := ioutil.ReadFile(filePath)
		if err != nil {
			return nil, errors.Wrap(err, "JFS Restore")
		}

		item := reflect.New(reft).Interface()

		err = json.Unmarshal(data, item)
		if err != nil {
			return nil, errors.Wrap(err, "JSF Restore")
		}

		iset, cast := item.(inset.InsetInterface)
		if !cast {
			return nil, errors.New("Restored object cant to be casted to InsetInterface")
		}
		res[iset.GetId()] = iset
	}
	return res, nil
}
