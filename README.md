```
██╗░░██╗░█████╗░██████╗░███╗░░██╗
██║░██╔╝██╔══██╗██╔══██╗████╗░██║
█████═╝░██║░░██║██████╔╝██╔██╗██║
██╔═██╗░██║░░██║██╔══██╗██║╚████║
██║░╚██╗╚█████╔╝██║░░██║██║░╚███║
╚═╝░░╚═╝░╚════╝░╚═╝░░╚═╝╚═╝░░╚══╝
```

In-memory reactivity database engine which can store and restore data to/from JSON files or MongoDB.\
Written in Go for Go.
```
KORN = Keep + Observe + React + eNgine
```
# Quick Start :: Just five steps

### 1. Define some observable structure
Nested structures are allowed.
```go
type User struct {
    korn.Inset `korn:"-" bson:",inline"` // required, the Inset is needed for communication between the Engine and current object
    Name    string `korn:"nameChanged"` // a "korn" tag defines an event name which will be invoked after changing
    Enabled bool   `korn:"enabledChanged"`
}
```
As weel the Inseet contains Id and Updated fields. Id field has BSON-tag "_id". Updated field contains last commint time.\
Warning: you cannot to add your own Id and Updated fields to your structure. It makes a panic.

### 2. Init Korn
Create basic KORN actors: the Engine and the Holder.\
Bind reactions names and handlers.\
Make and capture observable targets.

```go
engine := korn.Engine("demo")
holder, err := engine.Holder("users", User{})

holder.Bind("add", addHandler) // required, the "add" event have to be defined
holder.Bind("remove", removeHandler) // required, the "remove" too
holder.Bind("nameChanged", nameChangedHandler) // one regular event minimum requried
holder.Bind("enabledChanged", enabledChangedHandler)
```

### 3. Don't forget about data storing
Add one of two kinds storages if you need.\
One. JSON-files storage (JSF):
```go
engine.Connect("demo", "") 
engine.Restore() // restored data to memory
```
Two. MongoDB storage:
```go
engine.Connect("korn_mongo_demo", "mongodb://localhost") 
engine.Restore()  // restored data to memory
```

### 4. Make and capture your data
Create and catch some data objects.\
If you are using storage then captured objects will store automatically.
```go
users := map[string]*User{
    "root": user("root"), 
    "bob": user("bob"), 
    "guest": nil,
}
holder.Capture(users) // capture targets 
engine.Activate() // activate the engine > reactivity enabling
```

### 5. Perform
To do somethig with any one target. Or many :)

```go
user := holder.Get("bob").(*User) // getting from the holder and cast to origin type pointer
user.Name = "Bob Smith" // do something
...
user.Enabled = false // do something else
user.Commit() // required for reactivity magic and storing :-)
```
Also you can catch errors from reactions.
```go
err := user.Commit() 
if err != nil {
    panic(err)
}
```
For your enjoy you can make a wrapper for the holder and it provides your types casting easy. \
As well all package methods return an error value and your code will stay clean ever :)

***

[Full example code](https://github.com/en-v/korn/blob/main/examples/example.go)
