package korn

import (
	"errors"

	"github.com/en-v/korn/holder"
	"github.com/en-v/log"
)

type Inset struct {
	___holder holder.IHolder
	___key    string
}

func (self *Inset) Commit() {
	if self.___holder != nil {
		err := self.___holder.LookAt(self.___key)
		if err != nil {
			panic(err)
		}
	} else {
		panic(errors.New("holder is not setted"))
	}
}

func (self *Inset) Link(hldr holder.IHolder) {
	self.___holder = hldr
}

func (self *Inset) SetKey(key string) {
	self.___key = key
}

func (self *Inset) Key() string {
	return self.___key
}

func (self *Inset) Clone() interface{} {
	clone := *self
	log.Debug(&self)
	return &clone
}
