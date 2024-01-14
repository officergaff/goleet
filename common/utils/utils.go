package utils

import (
	"log"
	"os"
	"reflect"
	"testing"
)

func TestError[T any](t *testing.T, got T, want T) {
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Got %v, wanted %v", got, want)
	}
}

func LoadTestInput(file string) []byte {
	dat, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Loading file failed: %v", err)
	}
	return dat
}
