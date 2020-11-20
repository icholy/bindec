# bindec

> Helper type that wraps the `encoding/binary` package

``` go
dec := bindec.NewDecoder(reader, binary.LittleEndian)
size, _ := dec.Int16(nil)
header, _ := dec.Bytes(int(size))
```
