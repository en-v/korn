package inset

const NAME = "Inset"
const REQUIRED = "required"

type InsetInterface interface {
	GetId() string
	SetId(string)
	Clone() interface{}
	Commit() error
	Link(interface{})
}
