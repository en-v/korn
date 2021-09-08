package holder

import (
	"github.com/en-v/korn/query"
	"github.com/en-v/log"
)

func (self *_Holder) Select(query query.S) (map[string]interface{}, error) {
	s := query.Selectors()
	if len(s) > 0 {
		log.Debug("Do somethig")
	}
	return nil, nil
}
