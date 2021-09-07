package container

type IJet interface {
	Key() string
	SetKey(string)
	Clone() interface{}
	Commit()
	Observe(IContainer)
}
