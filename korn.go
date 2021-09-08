package korn

import (
	"fmt"

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
	Holder(string) holder.IHolder
}

type _Engine struct {
	holders map[string]holder.IHolder
}

//Empty - create a new KORN instance
func Empty() IEngine {
	engine := &_Engine{
		holders: make(map[string]holder.IHolder),
	}
	return IEngine(engine)
}

//Kit - make a couple of basic elements: the korn engine and the objsect holder
func Kit(name string) (IEngine, holder.IHolder) {
	engine := Empty()
	holder := engine.Holder(name)
	return engine, holder
}

func (self *_Engine) Holder(name string) holder.IHolder {
	obs, exists := self.holders[name]
	if exists {
		return obs
	}

	self.holders[name] = holder.Make(name)
	return self.holders[name]
}

func (self *_Engine) Activate() error {
	if len(self.holders) == 0 {
		return errors.New("No holders found")
	}

	for name, obs := range self.holders {
		err := obs.Activate()
		if err != nil {
			go self.Shutdown()
			return errors.Wrap(err, fmt.Sprintf("holder('%s').Activation", name))
		}
	}
	return nil
}

func (self *_Engine) Shutdown() {
	for _, obs := range self.holders {
		obs.Shutdown()
	}
}
