package jfs

import (
	"os"

	"github.com/pkg/errors"
)

const FILE_EXTENSION = ".json"

//JFS - JSON Files Storage
type JFS struct {
	ready bool
	path  string
}

//Make - make JSON Files Storage
func Make(dirname string, path string) (*JFS, error) {
	if path == "" {
		path = "./"
	}
	fullpath := path + "/" + dirname

	err := dirPrepare(fullpath)
	if err != nil {
		return nil, errors.Wrap(err, "Make JSF")
	}

	return &JFS{
		ready: true,
		path:  fullpath,
	}, nil
}

func (self *JFS) Shutdown() {
	self.ready = false
}

func (self *JFS) IsReady() bool {
	return self.ready
}

func (self *JFS) fullFileName(holder string, id string) string {
	return self.path + "/" + holder + "/" + id + FILE_EXTENSION
}

func (self *JFS) Prepare(holderName string) error {
	err := dirPrepare(self.path + "/" + holderName)
	if err != nil {
		return errors.Wrap(err, "Prepare")
	}
	return nil
}

func dirPrepare(fullpath string) error {
	_, err := os.Stat(fullpath)
	if os.IsExist(err) {
		return nil
	}

	err = os.MkdirAll(fullpath, 0777)
	if err != nil {
		return errors.Wrap(err, "Make Directory")
	}
	return nil
}
