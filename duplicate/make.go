package duplicate

import (
	"errors"
	"reflect"

	"github.com/en-v/korn/core"
	"github.com/en-v/korn/event"
	"github.com/en-v/korn/inset"
	"github.com/en-v/log"
)

func Make(target inset.InsetInterface, holder string) (*Duplicate, error) {
	shot, err := makedoublet(target, nil, target.GetId())
	if err != nil {
		return nil, err
	}
	shot.holder = holder
	return shot, nil
}

func makedoublet(target interface{}, parent *Duplicate, key string) (*Duplicate, error) {
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
			tag, tagexists := fstruct.Tag.Lookup(core.TAG)
			if tagexists && tag == "-" {
				continue
			}

			if tagedFieldsExist(fstruct.Type) {
				shot.branches[fstruct.Name], err = makedoublet(fvalue.Interface(), shot, fstruct.Name)
				if err != nil {
					return nil, err
				}

			} else {
				if tagexists {
					err = event.IsNotReservedActionName(tag)
					if err != nil {
						return nil, err
					}
					shot.fields[fstruct.Name] = makeField(&fvalue, tag, true)
				} else {
					shot.fields[fstruct.Name] = makeField(&fvalue, "", false)
				}
			}

		default:
			tag, tagexists := fstruct.Tag.Lookup(core.TAG)

			if tagexists {
				err = event.IsNotReservedActionName(tag)
				if err != nil {
					return nil, err
				}
				shot.fields[fstruct.Name] = makeField(&fvalue, tag, true)
			} else {
				shot.fields[fstruct.Name] = makeField(&fvalue, "", false)
			}
		}
	}

	return shot, nil
}

func tagedFieldsExist(t reflect.Type) bool {
	for f := 0; f < t.NumField(); f++ {
		_, ok := t.Field(f).Tag.Lookup(core.TAG)
		if ok {
			return true
		}
	}
	return false
}

func makeField(inv *reflect.Value, rname string, observable bool) *Field {
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
		log.Trace(inv)
		outv = inv.Interface()
	}

	return &Field{
		Value:      outv,
		Kind:       inv.Kind(),
		Reaction:   rname,
		Observable: observable,
	}
}
