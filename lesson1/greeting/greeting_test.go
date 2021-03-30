package greeting_test

import (
	"testing"

	"github.com/nj-eka/gb_go_1/lesson1/greeting"
)

func TestHello(t *testing.T) {
	hello := "Hello, world."
	if out := greeting.SayHello(); out != hello {
		t.Errorf("Hello() = %q, want %q", out, hello)
	}
}

func TestGo(t *testing.T) {
	go1 := "Don't communicate by sharing memory, share memory by communicating."
	if out := greeting.SayGo(); out != go1 {
		t.Errorf("Go() = %q, want %q", out, go1)
	}
}
