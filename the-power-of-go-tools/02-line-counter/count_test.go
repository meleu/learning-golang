package count_test

import (
	"bytes"
	"testing"

	"github.com/meleu/count"
)

func TestCountReturnsNumberOfLines(t *testing.T) {
	t.Parallel()
	c := count.NewCounter()
	c.Input = bytes.NewBufferString("1\n2\n3")
	got := c.Lines()
	want := 3

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
