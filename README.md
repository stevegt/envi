# envi

A simple library for reading environment variables: converts to Go
types, provides default values.  

Useful for e.g. feature flags managed in the local environment. 

## Example

(Also see env_test.go.)

```go
package main

import (
	"fmt"
	"os"

	"github.com/stevegt/envi"
)

func main() {
	fmt.Println(envi.String("SHELL", "foo"))
	fmt.Println(envi.Int("SHLVL", 99))
	fmt.Println(envi.Bool("IAmNotSet", true))
	fmt.Println(envi.Float32("somefloat", 2.34))
	os.Setenv("somefloat", "1.23")
	fmt.Println(envi.Float64("somefloat", 2.34))
}
```
