package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

const (
	inputFile  = "./testdata/test1.md"
	resultFile = "test1.md.html"
	goldenFile = "./testdata/test1.md.html"
)

func TestParseContent(t *testing.T) {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		t.Fatal(err)
	}

	got, err := parseContent(input, "")
	if err != nil {
		t.Fatal(err)
	}
	assertEqualGolden(t, got)
}

func TestRun(t *testing.T) {
	var mockStdout bytes.Buffer

	if err := run(inputFile, "", &mockStdout, true); err != nil {
		t.Fatal(err)
	}

	resultFile := strings.TrimSpace(mockStdout.String())

	got, err := os.ReadFile(resultFile)
	if err != nil {
		t.Fatal(err)
	}

	assertEqualGolden(t, got)
	os.Remove(resultFile)
}

func assertEqualGolden(t testing.TB, got []byte) {
	t.Helper()
	want, err := os.ReadFile(goldenFile)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(got, want) {
		t.Logf("golden:\n%s\n", want)
		t.Logf("result:\n%s\n", got)
		t.Error("Result content does not match golden file")
	}
}
