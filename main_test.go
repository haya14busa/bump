package main

import (
	"testing"
)

func TestNextTag(t *testing.T) {
	got, err := nextTag([]string{"v1.0.1", "v1.11.0"}, PATCH)
	if err != nil {
		t.Fatal(err)
	}
	if want := "v1.11.1"; got != want {
		t.Errorf("nextTag({v1.0.1, v1.11.0}, PATCH) = %v, want %v", got, want)
	}
}
