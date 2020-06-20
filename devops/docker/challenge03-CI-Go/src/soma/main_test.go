package main

import "testing"

func TestSoma(t *testing.T) {
	result := soma(5, 5)
	if result != 10 {
		t.Errorf("soma function did not work, expected 10! but got %v", result)
	}
}
