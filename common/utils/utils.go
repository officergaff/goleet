package utils

import (
	"testing"
)

func TestError[T any](t *testing.T, got T, want T) {
	t.Fatalf("Got %v, wanted %v", got, want)
}
