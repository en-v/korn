package holder

import (
	"github.com/pkg/errors"
)

func (self *_Holder) originByKey(key string) (IInsert, error) {
	origin := self.origins[key]
	if origin == nil {
		return nil, nil
	}

	res, cast := origin.(IInsert)
	if !cast {
		return nil, errors.New("Target intercase cannot to be casted to the Thing type")
	}
	return res, nil
}
