package holder

import (
	"github.com/en-v/korn/inset"
	"github.com/pkg/errors"
)

func (self *_Holder) originById(key string) (inset.InsetInterface, error) {
	origin := self.origins[key]
	if origin == nil {
		return nil, nil
	}

	res, cast := origin.(inset.InsetInterface)
	if !cast {
		return nil, errors.New("Target intercase cannot to be casted to the Thing type")
	}
	return res, nil
}
