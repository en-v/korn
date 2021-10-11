package korn

import "github.com/en-v/korn/event"

//Engine - make new KORN-engine instance
func Engine(name string) IEngine {
	return IEngine(makeEngine(name))
}

func EmptyHandler() event.Handler {
	return func(event *event.Event) error {
		return nil
	}
}
