package storage

import (
	"reflect"
	"strings"

	"github.com/en-v/korn/inset"
	"github.com/en-v/korn/storage/jfs"
	"github.com/en-v/korn/storage/mdb"
	"github.com/pkg/errors"
)

type IStorage interface {
	//Shutdown - shutdown the storage
	Shutdown()
	//Dump - save an object to named space (folder or collection)
	Dump(string, inset.InsetInterface) error
	//IsReady - is storage ready?
	IsReady() bool
	//Prepare - prepare named placed (folder or collection)
	Prepare(string) error
	//Restore - restore data from named space to memory
	Restore(string, reflect.Type) (map[string]interface{}, error)
	//Remove - remove an object from holder by id
	Remove(string, string) error
	//Reset - remove from objects and delete all data
	Reset() error
}

func Make(name string, param string) (IStorage, error) {
	if isMongo(param) {
		s, err := mdb.Make(name, param)
		if err != nil {
			return nil, errors.Wrap(err, "Mongo Storage Init")
		}
		return s, nil

	} else {
		s, err := jfs.Make(name, param)
		if err != nil {
			return nil, errors.Wrap(err, "JSON Storage Init")
		}
		return s, nil
	}
}

func isMongo(p string) bool {
	return strings.Index(p, "mongodb://") == 0
}
