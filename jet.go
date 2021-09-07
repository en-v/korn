package reactor

import (
	"errors"

	"github.com/en-v/log"
	"github.com/en-v/reactor/container"
)

type Jet struct {
	___container container.IContainer
	___key       string
}

func (this *Jet) Commit() {
	if this.___container != nil {
		err := this.___container.LookAt(this.___key)
		if err != nil {
			panic(err)
		}
	} else {
		panic(errors.New("Container is not setted"))
	}
}

func (this *Jet) Observe(obs container.IContainer) {
	this.___container = obs
}

func (this *Jet) SetKey(key string) {
	this.___key = key
}

func (this *Jet) Key() string {
	return this.___key
}

func (this *Jet) Clone() interface{} {
	clone := *this
	log.Debug(&this)
	return &clone
}
