package integration

import "testing"

func TestApi(t *testing.T) {
	expected := "Hello, API!"
	if result := "Hello, API!"; result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}