package helloworld

import (
	"testing"
)

func TestString(t *testing.T) {
	h := String()
	if h != "hello, world!" {
		t.Errorf("expected `hello, world!`, received %s", h)
	}
}
