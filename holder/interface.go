package holder

import (
	"github.com/en-v/korn/event"
	"github.com/en-v/korn/query"
)

type IHolder interface {
	//Capture - capture the target for observation and containing.
	//The target is an object which provide an intefcae "types.Target".
	//If kor is activated "add" reaction will be invoked
	Capture(target interface{}) error

	//Get - return the target by key.
	//The target is an object which provide an intefcae "types.Target".
	Get(key string) interface{}

	//All - return all contained and abservable targets.
	//The target is an object which provide an intefcae "types.Target".
	All() map[string]interface{}

	//Remove - remove the target by key.
	//The target is an object which provide an intefcae "types.Target".
	//If kor is activated "remove" reaction will be invoked
	Remove(key string) error

	//Bind - add an event handler
	//If an event handler with current name alredy exists then it will be removed and written as a new.
	//event.Handler -> github.com/en-vkor/event/Handler
	//Make a panic if the handlir is empty (nil)
	Bind(string, event.Handler)

	//Activate - activate the holder
	//No need to call self method cos the kor will call it
	Activate() error

	//Shutdown - deactivate holder
	//No need to call self method cos the kor will call it
	Shutdown()

	//Name - getter for holder name only
	Name() string

	//CatchError - waiting for internal errors
	CatchError() error

	Select(query.S) (map[string]interface{}, error)
	Count() int

	Reset() error

	LookAt(key string) error
}
