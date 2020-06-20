package greeting

import "testing"

func TestGreeting(t *testing.T) {
	result := greeting("Message test")
	if result != "<b>Message test</b>" {
		t.Errorf("greeting function did not work, expected <b>Message test</b>! but got %s", result)
	}

}
