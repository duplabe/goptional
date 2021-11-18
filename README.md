# Go Optional

POC Go generic implementation inspired by [Java Optional](https://www.baeldung.com/java-optional).

Requires Go 1.18.

## Examples

### None and Some

```go
none := None[bool]()
some := Some(42)

// check if the value is empty
none.IsEmpty()          // true
some.IsEmpty()          // false

// check if the value is present
none.IsPresent()        // false
some.IsPresent()        // true

// Get gives back the value and ok is true if it's present
val, ok := none.Get()   // val = false, ok = false
val, ok := some.Get()   // val = 42, ok = true

// get the value with default
none.GetOr(true)        // true
some.GetOr(100)         // 42

// get the value or the types zero value
none.GetOrZero()     // false
some.GetOrZero()     // 42

// GetOrPanic will panic if value is not present
none.GetOrPanic()       // will panic
some.GetOrPanic()       // 42

// IfPresent will run the callback if the value is present
none.IfPresent(func(val bool) { // noop
    // will never run
})
some.IfPresent(func(val int) {  // will print 42
    println(val)
})
```

### Of

```go
val := 42
var p *int

Of(&val)).IsPresent()  // true
Of(p)).IsPresent()     // false
```
