package main

import(
	"testing"
)

func TestPrintMessage(t *testing.T) {
	expected := "Hello World"
	result := printMessage(expected)
	if result != expected {
		t.Errorf("Expected %q, but got %q", expected, result)
	}
}

