package main

import "testing"

func TestParenMatcher(t *testing.T) {
	match := parenMatcher("(())")
	if match != true {
		t.Error("Expected true, for string (()) got ", match)
	}

	match = parenMatcher(")(")
	if match != false {
		t.Error("Expected false, for string (()) got ", match)
	}
}
