package types

type Thing interface {
	Key() string
	SetKey(string)
	Clone() Thing
	React()
	Observe(IObserver)
}
