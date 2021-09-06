package reactor

import (
	"fmt"

	"github.com/en-v/reactor/container"
	"github.com/pkg/errors"
)

//IReactor - reactivity framework for Go
type IReactor interface {
	//Active - activate the reactor instance
	Activate() error
	//Shutdown - deactivate the reactor instance
	Shutdown()
	//MakeContainer - get an container, if an container is not found than make a new named container
	Container(string) container.IContainer
}

//Reactor - reactivity framework for Go
type _Reactor struct {
	containers map[string]container.IContainer
}

//New - create a new Reactor instance
func New() IReactor {
	r := &_Reactor{
		containers: make(map[string]container.IContainer),
	}
	return IReactor(r)
}

func Kit(name string) (IReactor, container.IContainer) {
	rct := New()
	obs := rct.Container(name)
	return rct, obs
}

//Container - create a new named instance of Container
func (this *_Reactor) Container(name string) container.IContainer {
	obs, exists := this.containers[name]
	if exists {
		return obs
	}

	this.containers[name] = container.Make(name)
	return this.containers[name]
}

//Active - activate the reactor instance
func (this *_Reactor) Activate() error {
	if len(this.containers) == 0 {
		return errors.New("No containers found")
	}

	for name, obs := range this.containers {
		err := obs.Activate()
		if err != nil {
			go this.Shutdown()
			return errors.Wrap(err, fmt.Sprintf("Container('%s').Activation", name))
		}
	}
	return nil
}

//Shutdown - deactivate the reactor instance
func (this *_Reactor) Shutdown() {
	for _, obs := range this.containers {
		obs.Shutdown()
	}
}
