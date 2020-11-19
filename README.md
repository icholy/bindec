# binread

> Helper type that wraps the `encoding/binary` package

``` go
r := binread.NewReader(reader)
size, _ := r.ReadInt16()
header, _ := r.ReadBytes(int(size))
```
