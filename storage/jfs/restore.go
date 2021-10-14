package jfs

import (
	"io/ioutil"
	"reflect"

	"github.com/en-v/log"
	"github.com/pkg/errors"
)

func (self *JFS) Restore(folder string, reft reflect.Type) (map[string]interface{}, error) {
	path := self.path + "/" + folder + "/"
	list, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, errors.Wrap(err, "JFS Restore")
	}

	for _, f := range list {
		log.Trace(path + " = " + f.Name())
		data, err := ioutil.ReadFile(f.Name())
		if err != nil {
			return nil, errors.Wrap(err, "JFS Restore")
		}
		str := string(data)
		log.Trace(str)
	}
	return nil, nil
}
