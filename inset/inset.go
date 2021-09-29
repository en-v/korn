package inset

type InsetInterface interface {
	GetId() string
	SetId(string)
	Clone() interface{}
	Commit() error
	Link(interface{})
}
