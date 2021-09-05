package observer

import (
	"time"

	"github.com/en-v/log"
	"github.com/en-v/reactor/types"
)

func (this *Observer) observe() {
	ticker := time.NewTicker(time.Second * 3).C
	this.enabled = true
	for this.enabled {

		if this.mode == _MODE_LAZY {
			log.Debug("cycle")
			for _, tree := range this.things {
				this.checkThing(tree)
			}
		}

		select {
		case <-ticker:
			continue
		case <-this.stop:
			return
		}
	}
}

func (this *Observer) checkThing(thingTree *types.Snapshot) {
	err := thingTree.Compare()
	if err != nil {
		log.Error(err, thingTree.Name())
	}

	for thingTree.DiffExist() {
		diff := thingTree.PullDiff()
		sub, exist := this.reactions[diff]
		if !exist {
			log.Error("Subscription does not exist, name = " + diff)
			continue
		}
		sub.Handler(thingTree.GetOrigin())
	}
}

func (this *Observer) CheckThing(key string) {
	t, e := this.things[key]
	if e {
		this.checkThing(t)
	}
}
