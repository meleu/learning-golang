package hello_test

import (
	"bytes"
	"testing"

	"github.com/meleu/hello"
)

func TestPrintTo_PrintsHelloMessageToGivenWriter(t *testing.T) {
	t.Parallel()
	buf := new(bytes.Buffer)
	hello.PrintTo(buf)
	got := buf.String()
	want := "Hello, world\n"
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestPrintPrintsHelloMessageToOutput(t *testing.T) {
	t.Parallel()
	buf := new(bytes.Buffer)
	p := hello.NewPrinter()
	p.Output = buf
	p.Print()
	got := buf.String()

	want := "Hello, world\n"

	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
