# enviance

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

	env "github.com/stevegt/enviance"
)

func main() {
	fmt.Println(env.String("SHELL", "foo"))
	fmt.Println(env.Int("SHLVL", 99))
	fmt.Println(env.Bool("IAmNotSet", true))
	os.Setenv("somefloat", "1.23")
	fmt.Println(env.Float32("somefloat", 2.34))
	fmt.Println(env.Float64("somefloat", 2.34))
}
```
