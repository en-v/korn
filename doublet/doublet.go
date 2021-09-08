package doublet

import (
	"encoding/json"

	"reflect"
)

type Doublet struct {
	holder   string
	parent   *Doublet
	key      string
	fields   map[string]*Field
	branches map[string]*Doublet
	diffs    map[string]*Difference
}

type Field struct {
	Value      interface{}
	Kind       reflect.Kind
	Reaction   string
	Observable bool
}

type Difference struct {
	Reaction string
	Name     string
	Old      interface{}
	New      interface{}
	Holder   string
	Path     string
}

func create(parent *Doublet, key string) *Doublet {
	return &Doublet{
		parent:   parent,
		key:      key,
		fields:   make(map[string]*Field),
		branches: make(map[string]*Doublet),
	}

}

func (self *Doublet) ToString() string {
	b, err := json.MarshalIndent(self, "", "   ")
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func (self *Doublet) NextDifference() *Difference {
	for key, diff := range self.diffs {
		clone := &diff
		delete(self.diffs, key)
		return *clone

	}
	return nil
}

func (self *Doublet) GetPath() string {
	if self.parent == nil {
		return "[" + self.key + "]"
	}

	if self.parent.GetPath() != "" {
		return self.parent.GetPath() + "." + self.key
	}

	return self.key
}

func (self *Doublet) Name() string {
	return self.key
}

func (self *Doublet) getRoot() *Doublet {
	if self.parent != nil {
		return self.parent.getRoot()
	}
	return self
}
