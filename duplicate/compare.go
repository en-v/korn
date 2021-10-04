package duplicate

import (
	"reflect"

	"github.com/en-v/korn/inset"
	"github.com/en-v/log"
)

func (self *Duplicate) Compare(target inset.InsetInterface) error {

	if self.diffs == nil {
		self.diffs = make(map[string]*Difference)
	} else {
		for key := range self.diffs {
			delete(self.diffs, key)
		}
	}

	new, err := Make(target, self.holder)
	if err != nil {
		return err
	}

	err = self.compare(self.diffs, new)
	if err != nil {
		return err
	}

	self.branches = new.branches
	self.fields = new.fields
	if len(self.diffs) > 0 {
		log.Debug(self.id, "Differents found", self.diffs)
	}
	return nil
}

func (self *Duplicate) compare(diffsOut map[string]*Difference, newSnap *Duplicate) error {
	var found bool

	for fieldName, old := range self.fields {
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

		case reflect.Struct:
			found = !reflect.DeepEqual(old.Value, new.Value)

		default:
			found = old.Value != new.Value
		}

		if found {
			diffsOut[old.Reaction] = &Difference{
				Reaction: old.Reaction,
				Name:     fieldName,
				Previous: old.Value,
				Current:  new.Value,
				Holder:   self.getRoot().holder,
				Path:     self.GetPath() + "." + fieldName,
			}
		}
	}

	for name, branch := range self.branches {
		err := branch.compare(diffsOut, newSnap.branches[name])
		if err != nil {
			return err
		}
	}

	return nil
}
