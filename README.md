# The Korn

An in-memory reactivity database engine for Go which can store and restore data to JSON files or MongoDB.
[ K-O-R-N : Keep, Observe, React, eNgine ]

## Quick Start

**I.** Define some observable structure.
   
```go
type User struct {
    korn.Inset // required, the Inset is needed for communication between the Engine and current object
    Name    string `korn:"nameChanged"` // a "korn" tag defines an event name which will be invoked after changing
    Enabled bool   `korn:"enabledChanged"`
}
```

**II.** Create basic KORN actors: the Engine and the Holder. 
Bind reactions names and handlers.
Make and capture observable targets.

```go
engine, holder, err := korn.Kit("users")

holder.Bind("add", addHandler) // required, the "add" event have to be defined
holder.Bind("remove", removeHandler) // required, the "remove" too
holder.Bind("nameChanged", nameChangedHandler) // one regular event minimum requried
holder.Bind("enabledChanged", enabledChangedHandler)
```

Add storage if you want.
```go
err = holder.SetRefo(new("0")) // it needs to catch type info, temporary decision, 
err = engine.Connect("demo", "") // JSON-files storage - JFS
// or
err = engine.Connect("korn_mongo_demo", "mongodb://localhost")
err = engine.Restore()
```

```go
users := map[string]*User{"root": user("root"), "bob": user("bob"), "guest": nil}
err = holder.Capture(users) // capture targets
err = engine.Activate() // activa the engine, it needs for reactivity works
```

**III.** To do somethig with any one target. Or many :)

```go
user := holder.Get("bob").(*Type) // getting from in-memory base and cast to origin type pointer
user.Name = "Bob Smith"
user.Enabled = false
user.Commit() // require for reactivity magic :-)
```
or
```go
err := user.Commit() 
if err != nil {
    panic(err)
}
```
For your enjoy you can make a wrapper for the holder and it will cast your types easy. 

**IV.** PROFIT!11

[Full example's code](https://github.com/en-v/korn/blob/main/examples/example.go)

## To Do
- Query tools (like Mongo style) 
- MongoDB and simple JSON-files auto sync
- ...