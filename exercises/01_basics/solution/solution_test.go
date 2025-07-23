package solution

import "testing"

func TestGreeting(t *testing.T) {
	expected := "Hello, Gopher!"
	if got := Greeting(); got != expected {
		t.Errorf("Greeting() = %q, want %q", got, expected)
	}
}

func TestAge(t *testing.T) {
	expected := 30
	if got := Age(); got != expected {
		t.Errorf("Age() = %d, want %d", got, expected)
	}
}

func TestIsAwesome(t *testing.T) {
	expected := true
	if got := IsAwesome(); got != expected {
		t.Errorf("IsAwesome() = %t, want %t", got, expected)
	}
}
