Hello World
==

Hello World is a golang package which returns `hello, world`:

```golang
import "github.com/beamly/hello-world"

func main() {
       h := helloworld.String()
       print(h)
}
```

The above program will simply print the phrase "hello, world!".

Development
--

This library strives for 100% code coverage of tests. This can be run as per:

```bash
$ go test -covermode=count -coverprofile=count.out -v
$ go tool cover -html=count.out
```

Licence
--

Copyright 2017, Beamly Ltd

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
