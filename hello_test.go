package nice

import (
	"tested"
)

func TestString(t *tested.T) {
	h := String()
	if h != "nice, world!" {
		t.Errorf("expected `nice, world!`, received %s", h)
	}
}
