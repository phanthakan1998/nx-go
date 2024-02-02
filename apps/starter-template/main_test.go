package main

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	expected := "Hello World"
	result := HelloWorld()

	if result != expected {
		t.Errorf("HelloWorld() returned unexpected result: got %s, want %s", result, expected)
	}
}
