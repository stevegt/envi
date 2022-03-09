package envi

import (
	"os"
	"strconv"
	// . "github.com/stevegt/goadapt"
)

func String(key string, def string) (res string) {
	res, ok := os.LookupEnv(key)
	if !ok {
		return def
	}
	return
}

func Bool(key string, def bool) (res bool) {
	s, ok := os.LookupEnv(key)
	if !ok {
		return def
	}
	res, err := strconv.ParseBool(s)
	if err != nil {
		return def
	}
	return
}

func Int(key string, def int) (res int) {
	s, ok := os.LookupEnv(key)
	if !ok {
		return def
	}
	res, err := strconv.Atoi(s)
	if err != nil {
		return def
	}
	return
}

func Float32(key string, def float32) (res float32) {
	s, ok := os.LookupEnv(key)
	if !ok {
		return def
	}
	res64, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return def
	}
	res = float32(res64)
	return
}

func Float64(key string, def float64) (res float64) {
	s, ok := os.LookupEnv(key)
	if !ok {
		return def
	}
	res, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return def
	}
	return
}
