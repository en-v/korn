package types

import (
	"github.com/en-v/log"
)

type IObserver interface {
	CheckThing(key string)
}

type Injection struct {
	inj_obs IObserver
	inj_key string
}

func (this *Injection) React() {
	if this.inj_obs != nil {
		this.inj_obs.CheckThing(this.inj_key)
		log.Debug("Reactivate")
	} else {
		log.Error("Observer is not setted")
	}
}

func (this *Injection) Observe(obs IObserver) {
	this.inj_obs = obs
}

func (this *Injection) SetKey(key string) {
	this.inj_key = key
}

func (this *Injection) Key() string {
	return this.inj_key
}
