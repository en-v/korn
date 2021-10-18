package event

import "errors"

type Kind string

const (
	KIND_ADD    = "add"
	KIND_REMOVE = "remove"
	KIND_UPDATE = "update"
)

type Event struct {
	Id       string
	Origin   interface{} // pointer to origin, always
	Name     string      // name of changed field
	Error    error       // an error
	Kind     Kind        // kind of event: add, remove, change
	Previous interface{} // pointer to old value
	Current  interface{} // pointer to new value
	Holder   string
	Path     string
	Extra    interface{}
}

type Handler func(*Event) error

func IsNotReservedActionName(name string) error {
	if name == KIND_ADD || name == KIND_REMOVE || name == KIND_UPDATE {
		return errors.New("Tags cannot to be: " + KIND_ADD + " or " + KIND_REMOVE + " or " + KIND_UPDATE)
	}
	return nil
}

func (event *Event) HasExtra() bool {
	return event.Extra != nil
}
