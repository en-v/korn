package main

import (
	"strconv"
	"time"

	"github.com/en-v/log"
	"github.com/en-v/reactor"
	"github.com/en-v/reactor/observer"
	"github.com/en-v/reactor/types"
)

func main() {

	rch := make(chan int)

	size := 10000
	objs := make([]*SomeType, size)
	for i := 0; i < size; i++ {
		objs[i] = newThing(strconv.Itoa(i))

	}

	observer := observer.New("demo-observer")
	observer.ManualMode() // also can observer.LazyMode() - expensive mode, but no need to invoke React() methond on the Thing

	observer.On("someFieldChange", handlerOne)
	observer.On("subDisabled", handlerTwo)

	err := observer.Capture(objs) // can capture - slice, map or single
	if err != nil {
		panic(err)
	}

	reactor := reactor.New()
	err = reactor.Add(observer)
	if err != nil {
		panic(err)
	}

	err = reactor.Activate()
	if err != nil {
		panic(err)
	}

	go modify(objs)

	<-rch
}

func handlerOne(obj interface{}) {
	c := obj.(*SomeType)
	log.Debug("You see this text cos the field has changed", c.Key())
}

func handlerTwo(obj interface{}) {
	c := obj.(*SomeType)
	log.Debug("You see this text cos the field has changed", c.Key())
}

func modify(arr []*SomeType) {
	time.Sleep(time.Second * 5)
	log.Debug("SET NEW VALUE")
	arr[1].SomeField = "NEW_VAL"
	arr[1].IntVal = 144
	arr[1].Sub2.Enabled = true
	arr[1].React()
}

// ##################################################################

type SomeType struct {
	types.Injection
	id         string
	SomeField  string `react:"someFieldChange"`
	IntVal     int    `react:"setIntVal"`
	SomeField1 string `react:"someFieldChange"`
	IntVal1    int    `react:"setIntVal"`
	SomeField2 string `react:"someFieldChange"`
	IntVal2    int    `react:"setIntVal"`
	Sub        Sub
	Sub2       Sub
	Sub3       Sub
	Sub4       Sub
	Sub5       Sub
}

type Sub struct {
	SubField   string `react:"subFieldChange"`
	Enabled    bool   `react:"subDisabled"`
	SubField1  string `react:"subFieldChange"`
	Enabled1   bool   `react:"subDisabled"`
	SubField2  string `react:"subFieldChange"`
	Enabled2   bool   `react:"subDisabled"`
	SubField3  string `react:"subFieldChange"`
	Enabled3   bool   `react:"subDisabled"`
	SubField11 string `react:"subFieldChange"`
	Enabled11  bool   `react:"subDisabled"`
	SubField12 string `react:"subFieldChange"`
	Enabled12  bool   `react:"subDisabled"`
	SubField13 string `react:"subFieldChange"`
	Enabled13  bool   `react:"subDisabled"`
}

func newThing(id string) *SomeType {
	t := &SomeType{
		id:        id,
		SomeField: "fv." + id,
		Sub: Sub{
			SubField: "sysyss",
		},
	}
	t.SetKey(t.id)
	return t
}

func (this *SomeType) Clone() types.Thing {
	n := *this
	return &n
}
