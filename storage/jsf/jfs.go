package jsf

import (
	"os"

	"github.com/pkg/errors"
)

//JSF - JSON Files Storage
type JSF struct {
	ready bool
	path  string
}

//Make - make JSON Files Storage
func Make(dirname string, path string) (*JSF, error) {
	if path == "" {
		path = "./"
	}
	fullpath := path + "/" + dirname

	err := dirPrepare(fullpath)
	if err != nil {
		return nil, errors.Wrap(err, "Make JSF")
	}

	return &JSF{
		ready: true,
		path:  fullpath,
	}, nil
}

func (self *JSF) Shutdown() {
	self.ready = false
}

func (self *JSF) IsReady() bool {
	return self.ready
}

func (self *JSF) fullFileName(holder string, id string) string {
	return self.path + "/" + holder + "/" + id + ".json"
}

func (self *JSF) Prepare(holderName string) error {
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
