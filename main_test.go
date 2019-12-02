package main

import (
	"bytes"
	"testing"
)

func TestRun(t *testing.T) {
	stdout := new(bytes.Buffer)
	run(stdout)
	if stdout.String() == "" {
		t.Error("got empty result")
	}
}
