package container

type IJet interface {
	Key() string
	SetKey(string)
	Clone() interface{}
	React()
	Observe(IContainer)
}
