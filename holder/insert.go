package holder

type iInset interface {
	Key() string
	SetKey(string)
	Clone() interface{}
	Commit()
	Link(IHolder)
}
