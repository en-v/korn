### The Korn

A reactivity in-memory data base engine for Go.
[ K-O-R-N : Keep, Observe, React, eNgine ]

## Quick Start

1. Define some observable structure.
```go
type User struct {
    korn.Inset // required, the Inset is needed for communication between the Engine and current object
    Name    string `korn:"nameChanged"` // a "korn" tag defines an event name which will be invoked after changing
    Enabled bool   `korn:"enabledChanged"`
}
```

2. Create basic KORN actors: the Engine and the Holder. 
Define reactions and attach thier handlers.
Make and capture observable targets.
```go
engine, holder := korn.Kit("users")

holder.On("add", addHandler) // required, the "add" event have to be defined
holder.On("remove", removeHandler) // required, the "remove" too
holder.On("nameChanged", nameChangedHandler) // one regular event minimum requried
holder.On("enabledChanged", enabledChangedHandler)

users := map[string]*User{"root": user("root"), "bob": user("bob"), "guest": nil}
holder.Capture(users) // capture targets
engine.Activate() // activa the engine, it needs for reactivity works
```

3. To do somethig with any one target. Or many :)
```go
user := holder.Get("bob").(*Type) // getting from in-memory base and cast to origin type pointer
user.Name = "Bob Smith"
user.Enabled = false
user.Commit() // require for reactivity magic :-)
```
For your enjoy you can make a wrapper for the holder and it will cast your types easy. 

4. PROFIT!11

[Full example's code](https://github.com/en-v/korn/blob/main/examples/example.go)

## To Do
- Query tools (like Mongo style) 
- MongoDB and simple JSON-files auto sync
- ...