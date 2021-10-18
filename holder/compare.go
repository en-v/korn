package holder

import (
	"github.com/en-v/korn/duplicate"
	"github.com/en-v/korn/event"
	"github.com/en-v/korn/inset"
	"github.com/pkg/errors"
)

func (self *_Holder) compare(shot *duplicate.Duplicate, target inset.InsetInterface) error {

	err := shot.Compare(target)
	if err != nil {
		return errors.Wrap(err, shot.Name())
	}

	if shot.HasDifferences() {
		if self.activated && self.reactions.OnUpdate != nil {
			err = self.reactions.OnUpdate(&event.Event{
				Id:     target.GetId(),
				Origin: target,
				Name:   event.KIND_UPDATE,
				Kind:   event.KIND_UPDATE,
				Holder: self.name,
				Extra:  self.extra,
			})
			if err != nil {
				return errors.Wrap(err, "compare-onUpdate, "+shot.Name())
			}
		}

		for {
			diff := shot.NextDifference()
			if diff == nil {
				return nil
			}

			reaction, err := self.reactions.get(diff.Reaction)
			if err != nil {
				return err
			}

			err = reaction.Handler(&event.Event{
				Origin:   target,
				Name:     diff.Name,
				Previous: diff.Previous,
				Current:  diff.Current,
				Holder:   diff.Holder,
				Path:     diff.Path,
				Id:       target.GetId(),
				Kind:     event.KIND_UPDATE,
				Extra:    self.extra,
			})
			if err != nil {
				return errors.Wrap(err, "Compare.Handler."+diff.Name)
			}
		}
	}
	return nil
}
