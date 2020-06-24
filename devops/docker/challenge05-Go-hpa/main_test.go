package main

import "testing"

func TestSqrt(t *testing.T) {
	result := sqrt()
	if result != "Code.education Rocks" {
		t.Errorf("sqrt function did not work, expected Code.education Rocks, but got %s", result)
	}
}
