### THE REACTOR

A reactivity framework for Go

# Example

Full example's code in ./examples/example.go

```go
// 1st step - create an Observer
obs := reactor.Observer("demo-observer")
obs.ManualMode() // or observer.LazyMode()

// 2nd step - adding handlers to observer
obs.On("someFieldChange", handlerOne)
obs.On("subDisabled", handlerTwo)

// 3rd step - capturing targets, can capture - slice, array, map or single
err := obs.Capture(objs)
// ... error handling

// finally - create reacorm add observer and activare reactor
tor := reactor.New()
err = tor.Add(obs)
// ... error handling
err = tor.Activate()
// ... error handling
```

And after

```go
objs[1].SomeField = "NEW_VAL"
objs[1].IntVal = 144
objs[1].Sub2.Enabled = true
objs[1].React() // no need to call it in the lazy mode

//OR 

obj := bjs[1]
obj.IntVal = 144
obj.Sub2.Enabled = true
obj.React() // no need to call it in the lazy mode

```