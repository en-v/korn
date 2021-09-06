package snapshot

import (
	"errors"
	"reflect"

	"github.com/en-v/reactor/core"
)

func Make(target IJet, container string) (*Snapshot, error) {
	shot, err := makeSnapshot(target, nil, target.Key())
	if err != nil {
		return nil, err
	}
	shot.container = container
	return shot, nil
}

func makeSnapshot(target interface{}, parent *Snapshot, key string) (*Snapshot, error) {
	var err error

	if target == nil {
		return nil, errors.New("Something is NIL")
	}

	shot := create(parent, key)
	tval := reflect.Indirect(reflect.ValueOf(target))
	ttype := tval.Type()

	for f := 0; f < tval.NumField(); f++ {

		fvalue := tval.Field(f)
		fstruct := ttype.Field(f)

		switch fvalue.Kind() {
		case reflect.Struct:
			shot.branches[fstruct.Name], err = makeSnapshot(fvalue.Interface(), shot, fstruct.Name)
			if err != nil {
				return nil, err
			}

		default:
			tag, tagexists := fstruct.Tag.Lookup(core.TAG)
			if tagexists {
				if tag == core.TAG_ADD || tag == core.TAG_REM {
					return nil, errors.New("Tags cannot to be: " + core.TAG_ADD + " or " + core.TAG_REM)
				}
				shot.fields[fstruct.Name] = makeField(&fvalue, tag)
			}
		}
	}

	return shot, nil
}

func makeField(inv *reflect.Value, rname string) *Field {
	var outv interface{}

	switch inv.Kind() {
	case reflect.Map:
		mp := make(map[interface{}]interface{}, inv.Len())
		for _, k := range inv.MapKeys() {
			mp[k.Interface()] = inv.MapIndex(k).Interface()
		}
		outv = mp

	case reflect.Slice:
		slice := make([]interface{}, inv.Len())
		for i := 0; i < inv.Len(); i++ {
			slice[i] = inv.Index(i).Interface()
		}
		outv = slice

	default:
		outv = inv.Interface()
	}

	return &Field{
		Value:    outv,
		Kind:     inv.Kind(),
		Reaction: rname,
	}
}
