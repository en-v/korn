package kor

import (
	"errors"

	"github.com/en-v/kor/holder"
	"github.com/en-v/log"
)

type Insert struct {
	___holder holder.Holder
	___key    string
}

func (self *Insert) Commit() {
	if self.___holder != nil {
		err := self.___holder.LookAt(self.___key)
		if err != nil {
			panic(err)
		}
	} else {
		panic(errors.New("holder is not setted"))
	}
}

func (self *Insert) Link(hldr holder.Holder) {
	self.___holder = hldr
}

func (self *Insert) SetKey(key string) {
	self.___key = key
}

func (self *Insert) Key() string {
	return self.___key
}

func (self *Insert) Clone() interface{} {
	clone := *self
	log.Debug(&self)
	return &clone
}
