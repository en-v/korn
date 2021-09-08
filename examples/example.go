package main

import (
	"github.com/en-v/kor"
	"github.com/en-v/kor/event"
	"github.com/en-v/kor/query"
	"github.com/en-v/log"
)

func main() {
	singleholder()
	//multipleholders()
	//withErrorCapturing()
}

func singleholder() {

	// create kor kit (kor and holder holder), also you can use "New" for an empty kor create
	kor, holder := kor.Kit("single")
	// adding handlers to the holder
	holder.On("add", universalHandler)
	holder.On("remove", universalHandler)
	holder.On("string-changed", universalHandler)
	holder.On("int-changed", universalHandler)
	holder.On("struct-enabled-changed", universalHandler)
	holder.On("struct-slice-changed", universalHandler)
	holder.On("struct-map-string-changed", universalHandler)
	holder.On("struct-map-int-changed", universalHandler)

	// capture targets (map or single) and activate kor
	targets := map[string]*Type{"1": new("1"), "2": new("2"), "3": nil, "4": new("4")}
	err(holder.Capture(targets))
	err(kor.Activate())

	// to modificate targets and look at results
	modify(holder.Get("1").(*Type))
	modify(targets["4"])
	err(holder.Remove("2"))
	err(holder.Capture(new("5")))
	modify(holder.Get("1").(*Type))
	log.Debugw("All", "Length", len(holder.All()))
	// select
	q := query.S{
		"Struct.Enabled": query.NotEq(false),
		"Foo":            query.Eq(1.145),
	}
	res, e := holder.Select(q)
	err(e)
	log.Debug(res)
}

func multipleholders() {

	// create kor kit (kor and first holder), also you can use "New" for an empty kor create
	kor, first := kor.Kit("first")
	second := kor.Holder("second")
	// adding handlers to the holder
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

	// capture targets (map or single) and activate kor
	err(first.Capture(map[string]*Type{"1": new("1"), "2": new("2"), "3": nil}))
	err(second.Capture(map[string]*Type{"1": new("1"), "2": new("2"), "3": nil}))
	err(kor.Activate())

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
	log.Debugw(string(event.Kind), "Field", event.Name, "Old", event.Old, "New", event.New, "holder", event.Holder, "Path", event.Path)
}

func modify(t *Type) {
	t.String = "NEW_VAL"
	t.Int = 2
	t.Struct.Enabled = false
	t.Struct.MapString["zero"] = 0
	t.Struct.MapInt[0] = 0
	t.Struct.Slice[0] = "zero"
	t.Foo = 6.77
	t.Commit()
}

// ##################################################################

type Type struct {
	kor.Insert `kor:"-"`
	String     string `kor:"string-changed"`
	Int        int    `kor:"int-changed"`
	Foo        float32
	Struct     Struct
}

type Struct struct {
	Enabled   bool           `kor:"struct-enabled-changed"`
	Slice     []string       `kor:"struct-slice-changed"`
	MapString map[string]int `kor:"struct-map-string-changed"`
	MapInt    map[int]int    `kor:"struct-map-int-changed"`
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
