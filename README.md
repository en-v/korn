### THE KOR

A reactivity in-memory data base for Go
KOR > Keep, Observe, React

# Example

Full example's code in ./examples/example.go

```go
// create KOR kit (Engine and first Holder), also you can use "New" for an empty engine create
kor, cont := kor.Kit("single")

// adding handlers to the holder
cont.On("add", universalHandler)
cont.On("remove", universalHandler)
cont.On("string-changed", universalHandler)

// capture targets (map or single) and activate kor
targets := map[string]*Type{"1": new("1"), "2": new("2"), "3": nil, "4": new("4")}
cont.Capture(targets)
kor.Activate()

// to modificate targets and look at results
modify(cont.Get("1").(*Type))
modify(targets["4"])
cont.Remove("2"))
cont.Capture(new("5")))
modify(cont.Get("1").(*Type))
```