package utils

import (
	"reflect"
	"testing"
)

func TestError[T any](t *testing.T, got T, want T) {
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Got %v, wanted %v", got, want)
	}
}
