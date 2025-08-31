package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestPrintHelp(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	printHelp()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	expectedStrings := []string{
		"Railgun",
		"Usage:",
		"Options:",
		"--version",
		"--help",
	}

	for _, expected := range expectedStrings {
		if !strings.Contains(output, expected) {
			t.Errorf("Expected help output to contain '%s', but it didn't", expected)
		}
	}
}

func TestVersion(t *testing.T) {
	if version == "" {
		t.Error("Version should not be empty")
	}

	if !strings.Contains(version, "0.1.0") {
		t.Errorf("Expected version to contain '0.1.0', got %s", version)
	}
}
