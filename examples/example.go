package main

import (
	"github.com/en-v/log"
	"github.com/en-v/reactor"
	"github.com/en-v/reactor/event"
)

func main() {
	singleContainer()
	//multipleContainers()
	//withErrorCapturing()
}

func singleContainer() {

	// create reactor kit (Reactor and first Container), also you can use "New" for an empty reactor create
	rtor, first := reactor.Kit("single")
	// adding handlers to the container
	first.On("add", universalHandler)
	first.On("remove", universalHandler)
	first.On("string-changed", universalHandler)
	first.On("int-changed", universalHandler)
	first.On("struct-enabled-changed", universalHandler)
	first.On("struct-slice-changed", universalHandler)
	first.On("struct-map-string-changed", universalHandler)
	first.On("struct-map-int-changed", universalHandler)

	// capture targets (map or single) and activate reactor
	targets := map[string]*Type{"1": new("1"), "2": new("2"), "3": nil, "4": new("4")}
	err(first.Capture(targets))
	err(rtor.Activate())

	// to modificate targets and look at results
	modify(first.Get("1").(*Type))
	modify(targets["4"])
	err(first.Remove("2"))
	err(first.Capture(new("5")))
	modify(first.Get("1").(*Type))
}

func multipleContainers() {

	// create reactor kit (Reactor and first Container), also you can use "New" for an empty reactor create
	rtor, first := reactor.Kit("first")
	second := rtor.Container("second")
	// adding handlers to the container
	first.On("add", universalHandler)
	first.On("remove", universalHandler)
	first.On("string-changed", universalHandler)
	first.On("int-changed", universalHandler)
	first.On("struct-enabled-changed", universalHandler)
	first.On("struct-slice-changed", universalHandler)
	first.On("struct-map-string-changed", universalHandler)
	first.On("struct-map-int-changed", universalHandler)

	second.On("add", universalHandler)
	second.On("remove", universalHandler)
	second.On("string-changed", universalHandler)
	second.On("int-changed", universalHandler)

	// capture targets (map or single) and activate reactor
	err(first.Capture(map[string]*Type{"1": new("1"), "2": new("2"), "3": nil}))
	err(second.Capture(map[string]*Type{"1": new("1"), "2": new("2"), "3": nil}))
	err(rtor.Activate())

	// to modificate targets and look at results
	modify(first.Get("1").(*Type))
	err(first.Remove("2"))
	err(second.Remove("2"))
	err(first.Capture(new("5")))
	modify(first.Get("1").(*Type))
}

func withErrorCapturing() {

}

func err(err error) {
	if err != nil {
		panic(err)
	}
}

func universalHandler(event *event.Event) {
	log.Debugw(event.Kind, "Field", event.Name, "Old", event.Old, "New", event.New, "Container", event.Container, "Path", event.Path)
}

func modify(t *Type) {
	t.String = "NEW_VAL"
	t.Int = 2
	t.Struct.Enabled = false
	t.Struct.MapString["zero"] = 0
	t.Struct.MapInt[0] = 0
	t.Struct.Slice[0] = "zero"
	t.Foo = 6.77
	t.React()
}

// ##################################################################

type Type struct {
	reactor.Jet
	String string `react:"string-changed"`
	Int    int    `react:"int-changed"`
	Foo    float32
	Struct Struct
}

type Struct struct {
	Enabled   bool           `react:"struct-enabled-changed"`
	Slice     []string       `react:"struct-slice-changed"`
	MapString map[string]int `react:"struct-map-string-changed"`
	MapInt    map[int]int    `react:"struct-map-int-changed"`
}

func new(id string) *Type {
	t := &Type{
		String: id,
		Int:    1,
		Foo:    1.56,
		Struct: Struct{
			Enabled: true,
			Slice:   []string{"one", "two", "three"},
			MapString: map[string]int{
				"uno":   1,
				"do":    2,
				"stres": 3,
			},
			MapInt: map[int]int{
				1: 1,
				2: 2,
			},
		},
	}
	t.SetKey(id) // required
	return t
}
