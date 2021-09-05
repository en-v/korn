package observer

import (
	"sync"

	"github.com/en-v/reactor/types"
)

const (
	_MODE_MANUAL byte = 1
	_MODE_LAZY   byte = 2
)

type Observer struct {
	mutex   sync.Mutex
	name    string
	enabled bool

	reactions map[string]*types.Reaction
	things    map[string]*types.Snapshot

	originsSlice []interface{}
	originsMap   map[string]interface{}

	mode byte

	stop chan byte
}

//New - create a new instance of Observer
func New(name string) *Observer {
	return &Observer{
		name:         name,
		reactions:    make(map[string]*types.Reaction),
		stop:         make(chan byte),
		originsSlice: nil,
		originsMap:   nil,
		mode:         _MODE_MANUAL,
	}
}

//LazyMode - set the lazy mode
//It means that an observing will be in permanent cycle without manual changes confirmation (call  method React()).
//Current mode required more resources
func (this *Observer) LazyMode() {
	this.mode = _MODE_LAZY
}

//ManualMode - set the manul mode
//It means that an observing will be after manual changes confirmation (call method React()).
//Current mode required more developer attention and more lines of code
func (this *Observer) ManualMode() {
	this.mode = _MODE_MANUAL
}

//On - add event handler
func (this *Observer) On(name string, handler types.Handler) {
	this.mutex.Lock()
	this.reactions[name] = types.MakeReaction(name, handler)
	this.mutex.Unlock()
}

//Activate - activate observer
//No need to call this method cos the Reactor will call it
func (this *Observer) Activate() error {
	go this.observe()
	return nil
}

//Activate - deactivate observer
//No need to call this method cos the Reactor will call it
func (this *Observer) Shutdown() {
	this.stop <- 1
}

//Name - getter for observer name only
func (this *Observer) Name() string {
	return this.name
}
