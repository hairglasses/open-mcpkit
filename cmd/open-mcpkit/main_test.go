package main

import "testing"

func TestRunManifest(t *testing.T) {
	if err := run([]string{"manifest"}); err != nil {
		t.Fatalf("run() error = %v", err)
	}
}

func TestRunCall(t *testing.T) {
	if err := run([]string{"call", "sample_echo", "--param", "message=hello"}); err != nil {
		t.Fatalf("run() error = %v", err)
	}
}
