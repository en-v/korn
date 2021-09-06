package event

type Event struct {
	Key       string
	Origin    interface{} // pointer to origin, always
	Name      string      // name of changed field
	Error     error       // an error
	Kind      string      // kind of event: add, remove, change
	Old       interface{} // pointer to old value
	New       interface{} // pointer to new value
	Container string
	Path      string
}

type Handler func(*Event)
