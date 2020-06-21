package main

import "testing"

//to run this test without problems, you need to add this environment variable:
//export env CGO_ENABLED=0

func TestGreeting(t *testing.T) {
	result := greeting("Message test")
	if result != "<b>Message test</b>" {
		t.Errorf("greeting function did not work, expected <b>Message test</b>! but got %s", result)
	}

}
