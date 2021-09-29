package korn

import (
	"github.com/en-v/korn/holder"
	"github.com/pkg/errors"
)

//IEngine - reactivity framework for Go
type IEngine interface {
	//Active - activate the kor instance
	Activate() error
	//Shutdown - deactivate the kor instance
	Shutdown()
	//Holder - get an holder, if a holder is not found than will make a new named holder
	Holder(string) (holder.IHolder, error)
	//Connect - connect to storage (JSON-files or MongoDB)
	//1st param - name, 2nd param - path or connection string, 3nd param - model example
	Connect(string, string) error
	//Restore - restore data from storage into memory
	Restore() error
	//Reset - reset all data from storage and memory
	Reset() error
}

//Empty - create a new KORN instance
func Empty() IEngine {
	return IEngine(makeEngine())
}

//Kit - make a couple of basic elements: the korn engine and the objsect holder
func Kit(name string) (IEngine, holder.IHolder, error) {
	engine := Empty()
	holder, err := engine.Holder(name)
	if err != nil {
		return nil, nil, errors.Wrap(err, "Korn Kit")
	}
	return engine, holder, nil
}
