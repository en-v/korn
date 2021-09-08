package event

import "errors"

type Kind string

const (
	KIND_ADD    = "add"
	KIND_REMOVE = "remove"
	KIND_CHANGE = "change"
)

type Event struct {
	Key    string
	Origin interface{} // pointer to origin, always
	Name   string      // name of changed field
	Error  error       // an error
	Kind   Kind        // kind of event: add, remove, change
	Old    interface{} // pointer to old value
	New    interface{} // pointer to new value
	Holder string
	Path   string
}

type Handler func(*Event)

func IsNotReservedActionName(name string) error {
	if name == KIND_ADD || name == KIND_REMOVE {
		return errors.New("Tags cannot to be: " + KIND_ADD + " or " + KIND_REMOVE)
	}
	return nil
}
