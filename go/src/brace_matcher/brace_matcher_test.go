package main

import "testing"

func TestBraceMatcher(t *testing.T) {
	match := braceMatcher("(())")
	if match != true {
		t.Error("Expected true, for string (()) got ", match)
	}

	match = braceMatcher(")(")
	if match != false {
		t.Error("Expected false, for string (()) got ", match)
	}
}
