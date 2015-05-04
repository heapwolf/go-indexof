# SYOPSIS
The IndexOf() method returns the first index at which a given element can
be found in the slice, or -1 if it is not present. The only requirement of
this function is that the second argument must be a slice.

# INSTALL
```bash
go get github.com/hij1nx/go-indexof
```

# USAGE
```go
import "github.com/hij1nx/go-indexof"

var v = "bazz"
var a = []string{"bar", "foo", "bazz"}

var i = IndexOf(v, a) // "i" is assigned "2"

var v = 20
var a = []int{10, 20, 30}

var i = IndexOf(v, a) // "i" is assigned "1"
```

