# go-batcher

Simple batch library for Golang.

## Install

Go get:

```
$ go get github.com/travisjeffery/go-batcher
```

## Example

``` go
batcher := batcher.New(5 * time.Second, func(batch []interface{}) {
    // do something with the batch
})

batcher.Batch(interface{}{})
batcher.Batch(interface{}{})
batcher.Batch(interface{}{})

// etc.
```

## Author

Travis Jeffery

- [Twitter](http://twitter.com/travisjeffery)
- [Medium](http://medium.com/@travisjeffery)
- [Homepage](http://travisjeffery.com)

## License

MIT
