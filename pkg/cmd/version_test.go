package cmd

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func captureStdout(t *testing.T, fn func()) string {
	t.Helper()
	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create pipe: %v", err)
	}
	os.Stdout = w

	fn()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, r); err != nil {
		t.Fatalf("failed to copy output: %v", err)
	}
	return buf.String()
}

func TestVersionCmdOutput(t *testing.T) {
	cmd := newVersionCmd("1.2.3", "abc1234", "2024-06-15")
	cmd.SetArgs([]string{})

	output := captureStdout(t, func() {
		err := cmd.Execute()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	if !strings.Contains(output, "recotem version 1.2.3") {
		t.Errorf("expected version string, got %q", output)
	}
	if !strings.Contains(output, "commit: abc1234") {
		t.Errorf("expected commit string, got %q", output)
	}
	if !strings.Contains(output, "built:  2024-06-15") {
		t.Errorf("expected build time string, got %q", output)
	}
}

func TestVersionCmdDifferentValues(t *testing.T) {
	cmd := newVersionCmd("0.0.1-dev", "0000000", "2025-01-01")
	cmd.SetArgs([]string{})

	output := captureStdout(t, func() {
		err := cmd.Execute()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	if !strings.Contains(output, "0.0.1-dev") {
		t.Errorf("expected dev version, got %q", output)
	}
	if !strings.Contains(output, "0000000") {
		t.Errorf("expected zero commit, got %q", output)
	}
}
