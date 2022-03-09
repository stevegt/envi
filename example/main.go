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
	os.Setenv("somefloat", "1.23")
	fmt.Println(envi.Float32("somefloat", 2.34))
	fmt.Println(envi.Float64("somefloat", 2.34))
}
