package reactor

import (
	"fmt"

	"github.com/en-v/reactor/observer"
	"github.com/pkg/errors"
)

//Reactor - reactivity framework for Go
type Reactor struct {
	observers map[string]*observer.Observer
}

//New - create a new Reactor instance
func New() *Reactor {
	return &Reactor{
		observers: make(map[string]*observer.Observer),
	}
}

//Observer - create a new named instance of Observer
func Observer(name string) *observer.Observer {
	return observer.New(name)
}

//Add - add an observer ot the reactor
func (this Reactor) Add(observer *observer.Observer) error {
	_, exists := this.observers[observer.Name()]
	if exists {
		return errors.New("Observer with current name alredy exists, name = " + observer.Name())
	}
	this.observers[observer.Name()] = observer
	return nil
}

//Get - get an observer by name
func (this Reactor) Get(name string) (*observer.Observer, error) {
	observer, exists := this.observers[name]
	if !exists {
		return nil, errors.New("Observer with current name is not found, name = " + name)
	}

	return observer, nil
}

//Active - activate the reactor instance
func (this *Reactor) Activate() error {
	if len(this.observers) == 0 {
		return errors.New("No observers found")
	}

	for name, obs := range this.observers {
		err := obs.Activate()
		if err != nil {
			go this.Shutdown()
			return errors.Wrap(err, fmt.Sprintf("Observer('%s').Activation", name))
		}
	}
	return nil
}

//Shutdown - deactivate the reactor instance
func (this *Reactor) Shutdown() {
	for _, obs := range this.observers {
		obs.Shutdown()
	}
}
