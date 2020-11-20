# bindec

> Helper type that wraps the `encoding/binary` package

``` go
dec := bindec.NewDecoder(reader)
size, _ := dec.Int16()
header, _ := dec.Bytes(int(size))
```
