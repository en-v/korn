### The Korn

A reactivity in-memory data base engine for Go.
[ K-O-R-N : Keep, Observe, React, eNgine ]

## Quick Start

1. Define some observable structure.
```go
type User struct {
    korn.Inset // required, inset needs for connection between engine and current object
    Name    string `korn:"nameChanged"`
    Enabled bool   `korn:"enabledChanged"`
}
```

2. Create the Engine and the Holder. 
Attach reactions handlers.
Make and capture a little bit observable targets.
```go
engine, holder := korn.Kit("users")

holder.On("add", addHandler) // required
holder.On("remove", removeHandler) // required
holder.On("nameChanged", nameChangedHandler) 
holder.On("enabledChanged", enabledChangedHandler)

users := map[string]*User{"root": user("root"), "bob": user("bob"), "guest": nil}
holder.Capture(users)
engine.Activate() // activa the engine, it needs for reactivity works
```

3. To do somethig with any one target. Or many :)
```go
user := holder.Get("bob").(*Type) // getting from in-memory base and cast to origin type pointer
user.Name = "Bob Smith"
user.Enabled = false
user.Commit() // require for reactivity magic :-)
```

4. PROFIT!11

[Full example's code](https://github.com/en-v/korn/blob/main/examples/example.go)

## To Do
- Query tools (like Mongo style) 
- MongoDB and simple JSON-files auto sync
- ...