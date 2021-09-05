package types

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/en-v/log"
	"github.com/en-v/reactor/core"
)

type Tree struct {
	root       *Tree
	origin     Thing
	name       string
	leafs      map[string]*Leaf
	branches   map[string]*Tree
	diffsStack map[string]bool
}

type Leaf struct {
	Value    interface{}
	Kind     reflect.Kind
	Reaction string
}

func MakeTree(thing Thing) (*Tree, error) {
	tree, err := makeTree(thing, nil, thing.Key())
	if err != nil {
		return nil, err
	}
	tree.origin = thing
	return tree, nil
}

func makeTree(thing interface{}, root *Tree, name string) (*Tree, error) {

	if thing == nil {
		return nil, errors.New("Something is NIL")
	}

	tree := &Tree{
		root:     root,
		name:     name,
		leafs:    make(map[string]*Leaf),
		branches: make(map[string]*Tree),
	}

	svalue := reflect.Indirect(reflect.ValueOf(thing))
	stype := svalue.Type()

	for f := 0; f < svalue.NumField(); f++ {

		fvalue := svalue.Field(f)
		fstruct := stype.Field(f)
		reacTag, tagExists := fstruct.Tag.Lookup(core.TAG)

		if fstruct.Type.Kind() == reflect.Struct {
			subTree, err := makeTree(fvalue.Interface(), tree, fstruct.Name)
			if err != nil {
				return nil, err
			}
			tree.branches[fstruct.Name] = subTree
			continue
		}

		if tagExists {
			tree.leafs[fstruct.Name] = &Leaf{
				Value:    fvalue.Interface(),
				Kind:     fvalue.Kind(),
				Reaction: reacTag,
			}
		}
	}

	return tree, nil
}

func (this *Tree) ToString() string {
	b, err := json.MarshalIndent(this, "", "   ")
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func (this *Tree) DiffExist() bool {
	return len(this.diffsStack) > 0
}

func (this *Tree) PullDiff() string {
	for key := range this.diffsStack {
		delete(this.diffsStack, key)
		return key
	}
	return ""
}

func (this *Tree) GetRoute() string {
	if this.root == nil {
		return this.name
	}

	if this.root.GetRoute() != "" {
		return this.root.GetRoute() + "." + this.name
	}

	return this.name
}

func (this *Tree) Compare() error {

	if this.diffsStack == nil {
		this.diffsStack = make(map[string]bool)
	} else {
		for key := range this.diffsStack {
			delete(this.diffsStack, key)
		}
	}

	new, err := MakeTree(this.origin)
	if err != nil {
		return err
	}

	err = this.compare(this.diffsStack, new)
	if err != nil {
		return err
	}

	this.branches = new.branches
	this.leafs = new.leafs
	if len(this.diffsStack) > 0 {
		log.Debug(this.name, "Differents found", this.diffsStack)
	}
	return nil
}

func (this *Tree) compare(diffsOut map[string]bool, new *Tree) error {
	for name, leaf := range this.leafs {
		if leaf.Value != new.leafs[name].Value {
			log.Debug(leaf.Value, new.leafs[name].Value)
			diffsOut[leaf.Reaction] = true
		}
	}

	for name, branch := range this.branches {
		err := branch.compare(diffsOut, new.branches[name])
		if err != nil {
			return err
		}
	}

	return nil
}

func (this *Tree) GetOrigin() Thing {
	return this.origin
}

func (this *Tree) Name() string {
	return this.name
}
