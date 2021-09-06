package container

import (
	"github.com/en-v/reactor/core"
	"github.com/en-v/reactor/event"
	"github.com/en-v/reactor/snapshot"
	"github.com/pkg/errors"
)

func (this *Container) compare(shot *snapshot.Snapshot, target IJet) error {

	err := shot.Compare(target)
	if err != nil {
		return errors.Wrap(err, shot.Name())
	}

	for {
		diff := shot.NextDifference()
		if diff == nil {
			return nil
		}

		reaction, err := this.reactions.get(diff.Reaction)
		if err != nil {
			this.catchError(err)
			continue
		}

		reaction.Handler(&event.Event{
			Origin:    target,
			Name:      diff.Name,
			Old:       diff.Old,
			New:       diff.New,
			Container: diff.Container,
			Path:      diff.Path,
			Key:       target.Key(),
			Kind:      core.TAG_CHANGE,
		})
	}
}
