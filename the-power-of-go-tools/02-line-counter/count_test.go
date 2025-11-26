package count_test

import (
	"bytes"
	"testing"

	"github.com/meleu/count"
)

// func TestCountReturnsNumberOfLines(t *testing.T) {
// 	t.Parallel()
// 	c := count.NewCounter()
// 	c.Input = bytes.NewBufferString("1\n2\n3")
// 	got := c.Lines()
// 	want := 3
//
// 	if want != got {
// 		t.Errorf("want %d, got %d", want, got)
// 	}
// }

func TestLinesCountsLinesInInput(t *testing.T) {
	t.Parallel()
	inputBuf := bytes.NewBufferString("1\n2\n3")
	c, err := count.NewCounter(
		count.WithInput(inputBuf),
	)
	if err != nil {
		t.Fatal(err)
	}

	want := 3
	got := c.Lines()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
