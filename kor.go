package kor

import (
	"fmt"

	"github.com/en-v/kor/holder"
	"github.com/pkg/errors"
)

//Kor - reactivity framework for Go
type Kor interface {
	//Active - activate the kor instance
	Activate() error
	//Shutdown - deactivate the kor instance
	Shutdown()
	//Makeholder - get an holder, if an holder is not found than make a new named holder
	Holder(string) holder.Holder
}

//kor - reactivity framework for Go
type _Kor struct {
	holders map[string]holder.Holder
}

//New - create a new kor instance
func New() Kor {
	r := &_Kor{
		holders: make(map[string]holder.Holder),
	}
	return Kor(r)
}

func Kit(name string) (Kor, holder.Holder) {
	rct := New()
	obs := rct.Holder(name)
	return rct, obs
}

//Holder - create a new named instance of holder
func (self *_Kor) Holder(name string) holder.Holder {
	obs, exists := self.holders[name]
	if exists {
		return obs
	}

	self.holders[name] = holder.Make(name)
	return self.holders[name]
}

//Active - activate the kor instance
func (self *_Kor) Activate() error {
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

//Shutdown - deactivate the kor instance
func (self *_Kor) Shutdown() {
	for _, obs := range self.holders {
		obs.Shutdown()
	}
}
