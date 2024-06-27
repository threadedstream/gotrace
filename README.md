# gotrace

## Usage 
Here's how to use it
```go
package main

import (
	"fmt"
	"os"

	"github.com/threadedstream/trace"
)

func main() {
	const path = "profile.trace"
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	if trace.IsTraceV2(path) {
		parsedTrace, err := trace.ParseTraceV2(file)
		if err != nil {
			panic(err)
		}
		println(parsedTrace)
	} else {
		parsedTrace, err := trace.Parse(file, "")
		if err != nil {
			panic(err)
		}
		fmt.Println(parsedTrace)
	}
}
```
