package container

import (
	"github.com/pkg/errors"
)

type OriginsType byte

type Origins struct {
	Items map[string]interface{}
}

func (this *Container) originByKey(key string) (IJet, error) {
	origin := this.origins[key]
	if origin == nil {
		return nil, nil
	}

	res, cast := origin.(IJet)
	if !cast {
		return nil, errors.New("Target intercase cannot to be casted to the Thing type")
	}
	return res, nil
}
