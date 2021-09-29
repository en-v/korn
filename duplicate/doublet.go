package duplicate

import (
	"encoding/json"

	"reflect"
)

type Duplicate struct {
	holder   string
	parent   *Duplicate
	id       string
	fields   map[string]*Field
	branches map[string]*Duplicate
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

func create(parent *Duplicate, key string) *Duplicate {
	return &Duplicate{
		parent:   parent,
		id:       key,
		fields:   make(map[string]*Field),
		branches: make(map[string]*Duplicate),
	}

}

func (self *Duplicate) ToString() string {
	b, err := json.MarshalIndent(self, "", "   ")
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func (self *Duplicate) NextDifference() *Difference {
	for key, diff := range self.diffs {
		clone := &diff
		delete(self.diffs, key)
		return *clone

	}
	return nil
}

func (self *Duplicate) GetPath() string {
	if self.parent == nil {
		return "[" + self.id + "]"
	}

	if self.parent.GetPath() != "" {
		return self.parent.GetPath() + "." + self.id
	}

	return self.id
}

func (self *Duplicate) Name() string {
	return self.id
}

func (self *Duplicate) getRoot() *Duplicate {
	if self.parent != nil {
		return self.parent.getRoot()
	}
	return self
}
