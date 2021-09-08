package doublet

import (
	"reflect"

	"github.com/en-v/log"
)

func (self *Doublet) Compare(target iInsert) error {

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
		log.Debug(self.key, "Differents found", self.diffs)
	}
	return nil
}

func (self *Doublet) compare(diffsOut map[string]*Difference, newSnap *Doublet) error {
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

		default:
			found = old.Value != new.Value
		}

		if found {
			diffsOut[old.Reaction] = &Difference{
				Reaction: old.Reaction,
				Name:     fieldName,
				Old:      old.Value,
				New:      new.Value,
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
