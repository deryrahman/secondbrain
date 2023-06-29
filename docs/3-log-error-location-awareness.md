# [RFC 3] Log Error Location Awareness

Usually log is used whenever some function encounter error. However, on the same trace which path of the error is happening, same error message will be printed on multiple layer. This makes redundant message. This RFC discuss about how log should handle error so that it only print the message where the error is happening.

## Design

Suppose there's 3 layer which the error happen on the very bottom layer
```
Layer 1
   |
Layer 2
   |
Layer 3 --> error happens here
```

And every layer, log the error. Since the error is happening only on the layer 3, then the log should only print the message from log's caller on layer 3

## Expectation

```go
... Layer 1
err := layer2()
log.Error(err)
return err
...

... Layer 2
err := layer3()
log.Error(err)
return err
...

... Layer 3
err := something()
log.Error(err) --> only print the log here
return err
...
```

Things need to be considered:
- err should aware if it has already logged somewhere, if it's already printed, then it will never be printed (singleton)
- log.Error should have capability to mutate the err object, so that the err can be marked as already logged
- err message should not be changed
- but as err should not have any additional data structure, then poin 1 and 2 are not valid. only log which has  capability to aware which error has been logged

**Case for multiple error**
As there's multiple error, that multiple error should only print the respected error's element. Not all

```go
... Layer 2
err2 := happenHere()
err = errors.Join(err, err2)

err3 := layer3()
err = errors.Join(err, err3)

log.Error(err) --> should only print err2
return err
...

... Layer 3
err := something()
log.Error(err) --> print err
return err
...
```

## Approach

**Address based marker on log struct**
```go
func (l Log) Error(err error) {
  if ee, ok := err.(*joinError); ok {
    for _, e := range ee.Unwrap() {
      l.Error(e)
    }
    return
  }
  addr := fmt.Sprintf("%p", err)
  if l.isLogged[addr] == nil {
    // proceed
    l.isLogger[addr] = true
  }
}
```

Pros: simple implementation
Cons: complex way to handle reused mem address


**Accumulate and print the error array**
Logger contains worker which will print the error at interval. Same err won't be printed twice.
```go
func NewLog() Log {
  // start the worker
  // which will run l.print()
  ...
}

func (Log) Error(err) {
  l.errs = append(l.errs, err)
}

func (Log) print() {
  for _, err := range l.errs {
    addr := fmt.Sprintf("%p", err)
    if !isVisited[addr] {
      // proceed
      isVisited[addr] = true
    }
  }
  l.errs = []errors.Error{}
}
```

Pros: simple implemnentation
Cons: additional worker, which delayed the printing process

**Entirely new approach using custom error**
Realizing that the implementation above can lead to unecessary complexity, creating new custom error should be sufficient.

By nature, error flow can be modelled like this
```
      err            layer 1
    /  |    \
  err err    err     layer 2
 ---- ----    |
             err     layer 3
             ----

```

- Only error which lays on the leaf will be printed. Named it as `errors.RootCause(err)`
- RootCause is the error which are not come from this project or a new error
- Only printed the log.ErrorCause(err) on the very upper layer

```go
func (Log) ErrorCause(err error) {
  if ee, ok := err.(*joinError); ok {
    for _, e := range ee.Unwrap() {
      l.ErrorCause(e)
    }
    return
  }
  // proceed
}
```

- `RootCause` should store the line number, function name, and filename