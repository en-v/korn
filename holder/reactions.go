package holder

import (
	"reflect"

	"github.com/en-v/korn/core"
	"github.com/en-v/korn/event"
	"github.com/pkg/errors"
)

type Reactions struct {
	Tags     map[string]*Tag
	Items    map[string]*Reaction
	OnAdd    event.Handler
	OnRemove event.Handler
	OnUpdate event.Handler
}

type Reaction struct {
	Name    string
	Handler event.Handler
}

type Tag struct {
	Name  string
	Bound bool
}

func emptyReactions() *Reactions {
	return &Reactions{
		Items: make(map[string]*Reaction),
		Tags:  make(map[string]*Tag),
	}
}

func (self *Reactions) scanTags(t interface{}, zerolevel bool) error {
	if zerolevel {
		self.Tags = make(map[string]*Tag)
	}

	var alredyExists bool
	var err error

	tval := reflect.Indirect(reflect.ValueOf(t))
	ttype := tval.Type()

	for f := 0; f < tval.NumField(); f++ {
		fvalue := tval.Field(f)
		fstruct := ttype.Field(f)

		tag, tagexists := fstruct.Tag.Lookup(core.TAG)
		if tagexists {

			_, alredyExists = self.Tags[tag]
			if alredyExists {
				return errors.New("Action tag alredy exists, tag = " + tag)
			}

			if tag != "-" {
				self.Tags[tag] = &Tag{
					Name:  tag,
					Bound: false,
				}
			}
		}
		if fvalue.Kind() == reflect.Struct {
			err = self.scanTags(fvalue.Interface(), false)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (self *Reactions) add(name string, handler event.Handler) error {

	switch name {
	case event.KIND_ADD:
		self.OnAdd = handler

	case event.KIND_REMOVE:
		self.OnRemove = handler

	case event.KIND_UPDATE:
		self.OnUpdate = handler

	default:
		tag, found := self.Tags[name]
		if !found {
			return errors.New("Tag doesn't exist, tag = " + name)
		}

		tag.Bound = true
		self.Items[name] = &Reaction{
			Name:    name,
			Handler: handler,
		}
	}
	return nil
}

func (self *Reactions) get(name string) (*Reaction, error) {
	reaction, exists := self.Items[name]
	if !exists {
		return nil, errors.New("Reaction does not exist, name = " + name)
	}
	return reaction, nil
}

func (self *Reactions) checkUnbounds() error {
	unbs := ""
	for name, tag := range self.Tags {
		if !tag.Bound {
			if unbs == "" {
				unbs = name
			} else {
				unbs = unbs + ", " + name
			}
		}
	}
	if unbs != "" {
		return errors.New("Unbound tags found, tags = " + unbs)
	}
	return nil
}
