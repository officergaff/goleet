package utils

import "testing"

func TestError(t *testing.T, got string, want string) {
	t.Fatalf("Got %v, wanted %v", got, want)
}
