package types

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/en-v/log"
	"github.com/en-v/reactor/core"
)

type Snapshot struct {
	parent     *Snapshot
	origin     Thing
	name       string
	fields     map[string]*Field
	branches   map[string]*Snapshot
	diffsStack map[string]bool
}

type Field struct {
	Value    interface{}
	Kind     reflect.Kind
	Reaction string
}

func MakeSnapshot(thing Thing) (*Snapshot, error) {
	shot, err := snapshot(thing, nil, thing.Key())
	if err != nil {
		return nil, err
	}
	shot.origin = thing
	return shot, nil
}

func snapshot(thing interface{}, parent *Snapshot, name string) (*Snapshot, error) {

	if thing == nil {
		return nil, errors.New("Something is NIL")
	}

	shot := &Snapshot{
		parent:   parent,
		name:     name,
		fields:   make(map[string]*Field),
		branches: make(map[string]*Snapshot),
	}

	svalue := reflect.Indirect(reflect.ValueOf(thing))
	stype := svalue.Type()

	for f := 0; f < svalue.NumField(); f++ {

		fvalue := svalue.Field(f)
		fstruct := stype.Field(f)
		tag, tagexists := fstruct.Tag.Lookup(core.TAG)

		if fstruct.Type.Kind() == reflect.Struct {
			subTree, err := snapshot(fvalue.Interface(), shot, fstruct.Name)
			if err != nil {
				return nil, err
			}
			shot.branches[fstruct.Name] = subTree
			continue
		}

		if tagexists {
			shot.fields[fstruct.Name] = &Field{
				Value:    fvalue.Interface(),
				Kind:     fvalue.Kind(),
				Reaction: tag,
			}
		}
	}

	return shot, nil
}

func (this *Snapshot) ToString() string {
	b, err := json.MarshalIndent(this, "", "   ")
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func (this *Snapshot) DiffExist() bool {
	return len(this.diffsStack) > 0
}

func (this *Snapshot) PullDiff() string {
	for key := range this.diffsStack {
		delete(this.diffsStack, key)
		return key
	}
	return ""
}

func (this *Snapshot) GetRoute() string {
	if this.parent == nil {
		return this.name
	}

	if this.parent.GetRoute() != "" {
		return this.parent.GetRoute() + "." + this.name
	}

	return this.name
}

func (this *Snapshot) Compare() error {

	if this.diffsStack == nil {
		this.diffsStack = make(map[string]bool)
	} else {
		for key := range this.diffsStack {
			delete(this.diffsStack, key)
		}
	}

	new, err := MakeSnapshot(this.origin)
	if err != nil {
		return err
	}

	err = this.compare(this.diffsStack, new)
	if err != nil {
		return err
	}

	this.branches = new.branches
	this.fields = new.fields
	if len(this.diffsStack) > 0 {
		log.Debug(this.name, "Differents found", this.diffsStack)
	}
	return nil
}

func (this *Snapshot) compare(diffsOut map[string]bool, new *Snapshot) error {
	for name, leaf := range this.fields {
		if leaf.Value != new.fields[name].Value {
			log.Debug(leaf.Value, new.fields[name].Value)
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

func (this *Snapshot) GetOrigin() Thing {
	return this.origin
}

func (this *Snapshot) Name() string {
	return this.name
}
