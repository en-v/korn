package container

import (
	"github.com/en-v/log"
	"github.com/en-v/reactor/query"
)

func (this *Container) Select(query query.S) (map[string]interface{}, error) {
	s := query.Selectors()
	if len(s) > 0 {
		log.Debug("Do somethig")
	}
	return nil, nil
}
