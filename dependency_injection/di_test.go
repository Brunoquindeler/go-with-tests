package dependency_injection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Bruno")

	got := buffer.String()
	want := "Hello, Bruno"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
