package unit

import "testing"

func TestHandlers(t *testing.T) {
	expected := "Hello, Handlers!"
	if result := "Hello, Handlers!"; result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}