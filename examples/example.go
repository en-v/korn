package main

import (
	"errors"
	"time"

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
	engine := korn.Engine("korn-test")
	holder, err := engine.Holder("single", Type{})
	catch(err)

	// adding handlers to the holder
	holder.BindBasic(universalHandler, universalHandler, updateHandler)
	holder.Bind("string-changed", universalHandler)
	holder.Bind("int-changed", universalHandler)
	holder.Bind("time-changed", errorHandler)
	holder.Bind("struct-enabled-changed", universalHandler)
	holder.Bind("struct-slice-changed", universalHandler)
	holder.Bind("struct-map-string-changed", universalHandler)
	holder.Bind("struct-map-int-changed", universalHandler)
	catch(holder.CheckBindings())

	// use storages
	switch storageType {
	case 1:
		err = engine.Connect("")
	case 2:
		err = engine.Connect("mongodb://localhost")
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
	err = one.Commit()
	if err != nil {
		log.Error(err)
	}

	one.String = "IT WORKS AGAIN!"
	err = one.Commit()
	if err != nil {
		log.Error(err)
	}

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
	first.BindBasic(universalHandler, universalHandler, updateHandler)
	first.Bind("string-changed", universalHandler)
	first.Bind("int-changed", universalHandler)
	first.Bind("time-changed", universalHandler)
	first.Bind("struct-enabled-changed", universalHandler)
	first.Bind("struct-slice-changed", universalHandler)
	first.Bind("struct-map-string-changed", universalHandler)
	first.Bind("struct-map-int-changed", universalHandler)
	catch(first.CheckBindings())

	second.BindBasic(universalHandler, universalHandler, universalHandler)
	second.Bind("string-changed", universalHandler)
	second.Bind("int-changed", universalHandler)
	second.Bind("time-changed", universalHandler)
	second.Bind("struct-enabled-changed", universalHandler)
	second.Bind("struct-slice-changed", universalHandler)
	second.Bind("struct-map-string-changed", universalHandler)
	second.Bind("struct-map-int-changed", universalHandler)
	catch(second.CheckBindings())

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

func universalHandler(event *event.Event) error {
	log.Debugw(string(event.Kind), "Field", event.Name, "Old", event.Previous, "New", event.Current, "holder", event.Holder, "Path", event.Path)
	return nil
}

func updateHandler(event *event.Event) error {
	log.Trace(string(event.Kind), "Origin", event.Origin)
	return nil
}

func errorHandler(event *event.Event) error {
	return errors.New("ERROR HANDLER, " + event.Name)
}

func modify(t *Type) {
	t.String = "NEW_VAL"
	t.Int = 2
	t.Struct.Enabled = false
	t.Struct.MapString["zero"] = 0
	t.Struct.MapInt[0] = 0
	t.Struct.Slice[0] = "zero"

	t.Time = time.Now()
	err := t.Commit()
	if err != nil {
		log.Error(err)
	}
}

// ##################################################################

type Type struct {
	korn.Inset `korn:"-" bson:",inline"`
	String     string `korn:"string-changed"`
	Int        int    `korn:"int-changed"`
	//Foo        float32   `korn:"foo-unbound"`
	//Bar        float32   `korn:"bar-unbound"`
	Time   time.Time `korn:"time-changed"`
	Struct Struct
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
		Time:   time.Now(),
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
