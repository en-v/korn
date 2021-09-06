package snapshot

import (
	"encoding/json"

	"reflect"
)

type Snapshot struct {
	container string
	parent    *Snapshot
	key       string
	fields    map[string]*Field
	branches  map[string]*Snapshot
	diffs     map[string]*Difference
}

type Field struct {
	Value    interface{}
	Kind     reflect.Kind
	Reaction string
}

type Difference struct {
	Reaction  string
	Name      string
	Old       interface{}
	New       interface{}
	Container string
	Path      string
}

func create(parent *Snapshot, key string) *Snapshot {
	return &Snapshot{
		parent:   parent,
		key:      key,
		fields:   make(map[string]*Field),
		branches: make(map[string]*Snapshot),
	}

}

func (this *Snapshot) ToString() string {
	b, err := json.MarshalIndent(this, "", "   ")
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func (this *Snapshot) NextDifference() *Difference {
	for key, diff := range this.diffs {
		clone := &diff
		delete(this.diffs, key)
		return *clone

	}
	return nil
}

func (this *Snapshot) GetRoute() string {
	if this.parent == nil {
		return "[" + this.key + "]"
	}

	if this.parent.GetRoute() != "" {
		return this.parent.GetRoute() + "." + this.key
	}

	return this.key
}

func (this *Snapshot) Name() string {
	return this.key
}

func (this *Snapshot) getRoot() *Snapshot {
	if this.parent != nil {
		return this.parent.getRoot()
	}
	return this
}
