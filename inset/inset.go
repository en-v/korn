package inset

const NAME = "Inset"

type InsetInterface interface {
	GetId() string
	SetId(string)
	Clone() interface{}
	Commit() error
	Link(interface{})
}
