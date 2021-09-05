package observer

import (
	"reflect"

	"github.com/en-v/reactor/types"
	"github.com/pkg/errors"
)

//Capture - things must be a slice, map or single (interface of types.Thing)
func (this *Observer) Capture(things interface{}) error {

	var err error

	value := reflect.ValueOf(things)
	kind := value.Kind()

	if kind != reflect.Slice && kind != reflect.Map && kind != reflect.Struct {
		return errors.New("Things must be map, slice or struct, current = " + kind.String())
	}

	switch kind {
	case reflect.Struct:
		this.originsSlice = make([]interface{}, 1)
		this.originsSlice[0] = value.Interface()
		this.originsMap = nil

		this.things = make(map[string]*types.Tree, 1)

		item := value.Interface().(types.Thing)
		item.Observe(this)
		this.things[item.Key()], err = types.MakeTree(item)
		if err != nil {
			return errors.Wrap(err, "Single Captured")
		}

	case reflect.Slice:
		this.originsSlice = make([]interface{}, value.Len())
		for i := 0; i < value.Len(); i++ {
			this.originsSlice[i] = value.Index(i).Interface()
		}
		this.originsMap = nil

		this.things = make(map[string]*types.Tree, len(this.originsSlice))
		for _, item := range this.originsSlice {
			item := item.(types.Thing)
			item.Observe(this)
			this.things[item.Key()], err = types.MakeTree(item)
			if err != nil {
				return errors.Wrap(err, "Slice Captured")
			}
		}

	case reflect.Map:
		this.originsMap = make(map[string]interface{}, value.Len())
		for i := 0; i < value.Len(); i++ {
			item := value.Index(i).Interface().(types.Thing)
			item.Observe(this)
			this.originsMap[item.Key()] = item
		}
		this.originsSlice = nil

		this.things = make(map[string]*types.Tree, len(this.originsMap))
		for key, item := range this.originsMap {
			item := item.(types.Thing)
			item.Observe(this)
			this.things[key], err = types.MakeTree(item)
			if err != nil {
				return errors.Wrap(err, "Map Captured")
			}
		}
	}

	return err
}
