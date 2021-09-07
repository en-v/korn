package query

import (
	"strings"

	"github.com/en-v/log"
)

type S map[string]Condition // selector

func (this S) Selectors() []Selector {
	res := make([]Selector, len(this))
	i := 0
	for q, s := range this {
		res[i] = Selector{
			Path:     strings.Split(q, "."),
			Operator: s.O,
			Value:    s.V,
		}
		i++
	}
	log.Debug(res)
	return res
}

type Selector struct {
	Path     []string
	Operator Operator
	Value    interface{}
}

func (this *Selector) Match(v interface{}) bool {
	return false
}
