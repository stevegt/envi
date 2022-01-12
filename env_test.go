package env

import (
	"os"
	"testing"
	// . "github.com/stevegt/goadapt"
)

const key = "alksjfldsa"

func assert(t *testing.T, want, got interface{}) {
	t.Helper() // cause file:line info to show caller
	if want != got {
		t.Fatalf("want: %v %T  got: %v %T", want, want, got, got)
	}
}

func TestString(t *testing.T) {
	os.Unsetenv(key)
	assert(t, String(key, "hello"), "hello")
	os.Setenv(key, "")
	assert(t, String(key, "lsakdf"), "")
	os.Setenv(key, "hello")
	assert(t, String(key, "lsakdf"), "hello")
}

func TestBool(t *testing.T) {
	os.Unsetenv(key)
	assert(t, Bool(key, true), true)
	assert(t, Bool(key, false), false)
	os.Setenv(key, "")
	assert(t, Bool(key, true), true)
	assert(t, Bool(key, false), false)
	os.Setenv(key, "t")
	assert(t, Bool(key, false), true)
}

func TestInt(t *testing.T) {
	os.Unsetenv(key)
	assert(t, Int(key, 3), 3)
	os.Setenv(key, "")
	assert(t, Int(key, 3), 3)
	assert(t, Int(key, 4), 4)
	os.Setenv(key, "4")
	assert(t, Int(key, 5), 4)
}

func TestFloat32(t *testing.T) {
	os.Unsetenv(key)
	assert(t, Float32(key, 3.1), float32(3.1))
	os.Setenv(key, "")
	assert(t, Float32(key, 3.1), float32(3.1))
	assert(t, Float32(key, 4.2), float32(4.2))
	os.Setenv(key, "4.3")
	assert(t, Float32(key, 5.4), float32(4.3))
}

func TestFloat64(t *testing.T) {
	os.Unsetenv(key)
	assert(t, Float64(key, 3.1), float64(3.1))
	os.Setenv(key, "")
	assert(t, Float64(key, 3.1), float64(3.1))
	assert(t, Float64(key, 4.2), float64(4.2))
	os.Setenv(key, "4.3")
	assert(t, Float64(key, 5.4), float64(4.3))
}
