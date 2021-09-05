package types

type Handler func(current interface{})

type Reaction struct {
	Name    string
	Handler Handler
}

func MakeReaction(name string, handler Handler) *Reaction {
	return &Reaction{
		Name:    name,
		Handler: handler,
	}
}
