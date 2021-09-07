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

func (this *Snapshot) compare(diffsOut map[string]*Difference, newSnap *Snapshot) error {
	var found bool

	for fieldName, old := range this.fields {
		if !old.Observable {
			continue
		}
		new := newSnap.fields[fieldName]

		switch old.Kind {
		case reflect.Map:
			found = !reflect.DeepEqual(old.Value, new.Value)

		case reflect.Array:
			found = !reflect.DeepEqual(old.Value, new.Value)

		case reflect.Slice:
			found = !reflect.DeepEqual(old.Value, new.Value)

		default:
			found = old.Value != new.Value
		}

		if found {
			diffsOut[old.Reaction] = &Difference{
				Reaction:  old.Reaction,
				Name:      fieldName,
				Old:       old.Value,
				New:       new.Value,
				Container: this.getRoot().container,
				Path:      this.GetRoute() + "." + fieldName,
			}
		}
	}

	for name, branch := range this.branches {
		err := branch.compare(diffsOut, newSnap.branches[name])
		if err != nil {
			return err
		}
	}

	return nil
}
