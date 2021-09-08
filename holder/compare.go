package holder

import (
	"github.com/en-v/kor/doublet"
	"github.com/en-v/kor/event"
	"github.com/pkg/errors"
)

func (self *_Holder) compare(shot *doublet.Doublet, target IInsert) error {

	err := shot.Compare(target)
	if err != nil {
		return errors.Wrap(err, shot.Name())
	}

	for {
		diff := shot.NextDifference()
		if diff == nil {
			return nil
		}

		reaction, err := self.reactions.get(diff.Reaction)
		if err != nil {
			self.catchError(err)
			continue
		}

		reaction.Handler(&event.Event{
			Origin: target,
			Name:   diff.Name,
			Old:    diff.Old,
			New:    diff.New,
			Holder: diff.Holder,
			Path:   diff.Path,
			Key:    target.Key(),
			Kind:   event.KIND_CHANGE,
		})
	}
}
