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
# Quick Start [ 5 steps ]

### 1. Defining
Define your own observable structure or structures.\
Nested structures are allowed.
```go
type User struct {
    korn.Inset `korn:"-" bson:",inline"`
    Name    string `korn:"nameChanged"`
    Enabled bool   `korn:"enabledChanged"`
}
```
Use **korn.Inset** as embedded part of the root structure. Tags `korn:"-" bson:",inline"` are required.\
If you forget **korn.Inset** or tags then your app will catch a panic.\
Also **korn.Inset** contains **Id** and **Updated** fields.\
**Id** field has BSON-tag "_id" and "string" type.\ 
**Updated** field contains the last commit date and time.\
*Warning: do NOT add your own Id and Updated fields to your structure. It will cause panic.*\
Tag `korn` contains an action name that will invoke when the field value changes.

### 2. Init Korn
Create basic KORN actors: Engine and Holder.\
Bind reactions names and handlers.\
Make and capture observable targets.

```go
engine := korn.Engine("demo")
holder, err := engine.Holder("users", User{})

holder.Bind("add", addHandler) 
holder.Bind("remove", removeHandler) 
holder.Bind("nameChanged", nameChangedHandler) // one regular event minimum requried
holder.Bind("enabledChanged", enabledChangedHandler)
```
Actions `add` and `remove` must always be defined and bound.
You can use more than one holder.\ 
Holders count is unlimited but you have to remember that name of the holder must be unique.\
If you have many types of data then your scenario is "one holder per one type".

### 3. Don't forget about data storing
Add one of two kinds storage if you need it. The storage using is optional.\
You can use NO data storage then your data will be lost after your app close (it is a useful case for temporary data).

**JSF** - storage based on simle JSON files (one file per object):
```go
engine.Connect("demo", "") 
engine.Restore() 
```
**MDB** - MongoDB storage:
```go
engine.Connect("korn_mongo_demo", "mongodb://localhost") 
engine.Restore() 
```
Call the `Restore()` method if you need to restore the last saved data from the storage.

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
To do something with any one target. Or many :)

```go
user := holder.Get("bob").(*User) // getting from the holder and cast to origin type pointer
user.Name = "Bob Smith" // do something
...
user.Enabled = false // do something else
user.Commit() // required for reactivity magic and storing :-)
```
`Commit()` method is needed to cause reactions and save the changes.
Also you can catch errors from reactions.
```go
err := user.Commit() 
if err != nil {
    panic(err)
}
```
For your enjoyment you can make a wrapper for the holder and it provides your type casting easy. \
As well all package methods return an error value and your code will stay clean ever :)

***

[Full example code](https://github.com/en-v/korn/blob/main/examples/example.go)
