package main

import (
	"testing"
)

func TestMyId(t *testing.T) {

	id := myId()
	length := len([]rune(id))

	if length < 64 {
		t.Errorf("Invalid sha256sum: %q", id)
	}
}
