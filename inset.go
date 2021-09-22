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

func (self *Inset) Commit() error {
	if self.___holder == nil {
		return errors.New("Holder Is Null.")
	}
	err := self.___holder.LookAt(self.___key)
	if err != nil {
		return err
	}
	return nil
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
