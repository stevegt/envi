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
