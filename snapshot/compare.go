package snapshot

import (
	"reflect"

	"github.com/en-v/log"
)

func (this *Snapshot) Compare(target IJet) error {

	if this.diffs == nil {
		this.diffs = make(map[string]*Difference)
	} else {
		for key := range this.diffs {
			delete(this.diffs, key)
		}
	}

	new, err := Make(target, this.container)
	if err != nil {
		return err
	}

	err = this.compare(this.diffs, new)
	if err != nil {
		return err
	}

	this.branches = new.branches
	this.fields = new.fields
	if len(this.diffs) > 0 {
		log.Debug(this.key, "Differents found", this.diffs)
	}
	return nil
}

func (this *Snapshot) compare(diffsOut map[string]*Difference, new *Snapshot) error {
	var found bool

	for fieldName, Old := range this.fields {
		New := new.fields[fieldName]

		switch Old.Kind {
		case reflect.Map:
			found = !reflect.DeepEqual(Old.Value, New.Value)

		case reflect.Array:
			found = !reflect.DeepEqual(Old.Value, New.Value)

		case reflect.Slice:
			found = !reflect.DeepEqual(Old.Value, New.Value)

		default:
			found = Old.Value != New.Value
		}

		if found {
			diffsOut[Old.Reaction] = &Difference{
				Reaction:  Old.Reaction,
				Name:      fieldName,
				Old:       Old.Value,
				New:       New.Value,
				Container: this.getRoot().container,
				Path:      this.GetRoute() + "." + fieldName,
			}
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
