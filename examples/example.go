package main

import (
	"github.com/en-v/korn"
	"github.com/en-v/korn/event"
	"github.com/en-v/korn/query"
	"github.com/en-v/log"
)

func main() {
	log.Init(log.ORANGE, "KORN")
	SingleHolder(1)   // example #1.1
	SingleHolder(2)   // example #1.2
	MultipleHolders() // example #2.1
}

func SingleHolder(storageType int) {

	// create engine kit (engine and holder holder), also you can use "New" for an empty engine create
	engine := korn.Engine("demo")
	holder, err := engine.Holder("single", Type{})
	catch(err)

	// adding handlers to the holder
	holder.Bind("add", universalHandler)
	holder.Bind("remove", universalHandler)
	holder.Bind("string-changed", universalHandler)
	holder.Bind("int-changed", universalHandler)
	holder.Bind("struct-enabled-changed", universalHandler)
	holder.Bind("struct-slice-changed", universalHandler)
	holder.Bind("struct-map-string-changed", universalHandler)
	holder.Bind("struct-map-int-changed", universalHandler)

	// use storages
	switch storageType {
	case 1:
		err = engine.Connect("demo", "")
	case 2:
		err = engine.Connect("korn_mongo_demo", "mongodb://localhost")
	}
	catch(err)

	err = engine.Restore()
	catch(err)

	err = engine.Activate()
	catch(err)

	first := holder.Get("1")
	if first == nil {
		targets := map[string]*Type{"1": new("1"), "2": new("2"), "3": nil, "4": new("4")}
		err = holder.Capture(targets)
		catch(err)
		first = holder.Get("1")
	}
	log.Trace(first)

	modify(first.(*Type))
	//modify(targets["4"])
	//catch(holder.Remove("2"))
	//catch(holder.Capture(new("5")))
	modify(holder.Get("1").(*Type))
	log.Debugw("All", "Length", len(holder.All()))
	// select
	q := query.S{
		"Struct.Enabled": query.NotEq(false),
		"Foo":            query.Eq(1.145),
	}
	res, err := holder.Select(q) // experimental feature, in-rpogress
	catch(err)

	one := holder.Get("1").(*Type)
	one.String = "IT WORKS"
	one.Commit()

	one.String = "IT WORKS AGAIN!"
	one.Commit()

	log.Debug(res)
}

func MultipleHolders() {

	// create engine kit (engine and first holder), also you can use "New" for an empty engine create
	engine := korn.Engine("demo2")
	first, err := engine.Holder("first", Type{})
	catch(err)
	second, err := engine.Holder("second", Type{})
	catch(err)
	// adding handlers to the holder
	first.Bind("add", universalHandler)
	first.Bind("remove", universalHandler)
	first.Bind("string-changed", universalHandler)
	first.Bind("int-changed", universalHandler)
	first.Bind("struct-enabled-changed", universalHandler)
	first.Bind("struct-slice-changed", universalHandler)
	first.Bind("struct-map-string-changed", universalHandler)
	first.Bind("struct-map-int-changed", universalHandler)

	second.Bind("add", universalHandler)
	second.Bind("remove", universalHandler)
	second.Bind("string-changed", universalHandler)
	second.Bind("int-changed", universalHandler)

	// capture targets (map or single) and activate kor
	catch(first.Capture(map[string]*Type{"1": new("1"), "2": new("2"), "3": nil}))
	catch(second.Capture(map[string]*Type{"1": new("1"), "2": new("2"), "3": nil}))
	catch(engine.Activate())

	// to modificate targets and look at results
	modify(first.Get("1").(*Type))
	catch(first.Remove("2"))
	catch(second.Remove("2"))
	catch(first.Capture(new("5")))
	modify(first.Get("1").(*Type))
}

func catch(err error) {
	if err != nil {
		panic(err)
	}
}

func universalHandler(event *event.Event) {
	log.Debugw(string(event.Kind), "Field", event.Name, "Old", event.Previous, "New", event.Current, "holder", event.Holder, "Path", event.Path)
}

func modify(t *Type) {
	t.String = "NEW_VAL"
	t.Int = 2
	t.Struct.Enabled = false
	t.Struct.MapString["zero"] = 0
	t.Struct.MapInt[0] = 0
	t.Struct.Slice[0] = "zero"
	t.Foo = 6.77
	err := t.Commit()
	if err != nil {
		panic(err)
	}
}

// ##################################################################

type Type struct {
	korn.Inset `korn:"-" bson:",inline"`
	String     string `korn:"string-changed"`
	Int        int    `korn:"int-changed"`
	Foo        float32
	Struct     Struct
}

type Struct struct {
	Enabled   bool           `korn:"struct-enabled-changed"`
	Slice     []string       `korn:"struct-slice-changed"`
	MapString map[string]int `korn:"struct-map-string-changed"`
	MapInt    map[int]int    `korn:"struct-map-int-changed"`
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
	t.SetId(id) // required
	return t
}
