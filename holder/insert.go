package holder

type IInsert interface {
	Key() string
	SetKey(string)
	Clone() interface{}
	Commit()
	Link(Holder)
}
