# lazytz-go
Little repo to help bridge the gap between lazy time zone offset providers and golang.

# Installation

```shell
go get -u github.com/dcarbone/lazytz-go
```

# Usage

```go
package main

import (
	"fmt"
	"os"

	"github.com/dcarbone/lazytz-go"
)

func main() {
	if len(os.Args) < 1 {
		fmt.Println("Must provide timezone abbreviation as first argument.")
		os.Exit(1)
	}

	abbr := os.Args[1]
	
	info, ok := lazytz.Get(abbr)
	if !ok {
		fmt.Printf("No timezone registered with abbreviation %q\n", abbr)
		fmt.Println("Current list:")
		fmt.Println(lazytz.GetAbbreviations())
		os.Exit(1)
    }
	
	fmt.Println("Timezone info:")
	fmt.Println(info)
}
```