### THE REACTOR

A reactivity framework for Go


# Example

Full example's code in ./examples/example.go

```go
// create reactor kit (Reactor and first Container), also you can use "New" for an empty reactor create
rtor, cont := reactor.Kit("single")

// adding handlers to the container
cont.On("add", universalHandler)
cont.On("remove", universalHandler)
cont.On("string-changed", universalHandler)
cont.On("int-changed", universalHandler)
cont.On("struct-enabled-changed", universalHandler)
cont.On("struct-slice-changed", universalHandler)
cont.On("struct-map-string-changed", universalHandler)
cont.On("struct-map-int-changed", universalHandler)

// capture targets (map or single) and activate reactor
targets := map[string]*Type{"1": new("1"), "2": new("2"), "3": nil, "4": new("4")}
cont.Capture(targets)
rtor.Activate()

// to modificate targets and look at results
modify(cont.Get("1").(*Type))
modify(targets["4"])
cont.Remove("2"))
cont.Capture(new("5")))
modify(cont.Get("1").(*Type))
```