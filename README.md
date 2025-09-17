# terr

A traceable error library for golang.

## Example

```golang
import (
    "net/http"

    "github.com/deadblue/terr"
)

func loadData() (err error) {
    resp, err := terr.TraceError(http.Get("http://example.com/"))
    if err != nil {
        return
    }
    defer resp.Close()
    // TODO: Parse resp
}

func divide(m, n int) (err error) {
    if (n == 0) {
        err = terr.New("can not divide by zero")
    }
    return
}

func main() {
    loadData()
    divide(5, 0)
}
```

## License

MIT