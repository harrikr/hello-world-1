Hello World
==

Hello World is a golang package which returns `hello, world`:

```golang
import "github.com/zeebox/hello-world-poc"

func main() {
	h := helloworld.String
	print(h)
}
```

The above program will simply print the phrase "hello, world".

Development
--

This library strives for 100% code coverage of tests. This can be run as per:

```bash
$ go test -covermode=count -coverprofile=count.out -v
$ go tool cover -html=count.out
```
