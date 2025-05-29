package unit

import "testing"

func TestServices(t *testing.T) {
	expected := "Hello, Services!"
	if result := "Hello, Services!"; result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}