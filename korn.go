package korn

//Engine - make new KORN-engine instance
func Engine(name string) IEngine {
	return IEngine(makeEngine(name))
}
