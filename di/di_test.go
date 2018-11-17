package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Alice")

	got := buffer.String()
	want := "Hello, Alice"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}