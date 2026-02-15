package cmd

import (
	"bytes"
	"testing"
)

func TestCompletionValidShells(t *testing.T) {
	shells := []string{"bash", "zsh", "fish", "powershell"}
	for _, shell := range shells {
		t.Run(shell, func(t *testing.T) {
			root := NewRootCmd("1.0.0", "abc", "2024-01-01")
			// Suppress output from leaking to test stdout.
			var devnull bytes.Buffer
			root.SetOut(&devnull)
			root.SetErr(&devnull)
			root.SetArgs([]string{"completion", shell})

			err := root.Execute()
			if err != nil {
				t.Errorf("expected no error for shell %q, got %v", shell, err)
			}
		})
	}
}

func TestCompletionInvalidShell(t *testing.T) {
	root := NewRootCmd("1.0.0", "abc", "2024-01-01")
	var buf bytes.Buffer
	root.SetOut(&buf)
	root.SetErr(&buf)
	root.SetArgs([]string{"completion", "invalid"})

	err := root.Execute()
	if err == nil {
		t.Error("expected error for invalid shell, got nil")
	}
}

func TestCompletionNoArgs(t *testing.T) {
	root := NewRootCmd("1.0.0", "abc", "2024-01-01")
	var buf bytes.Buffer
	root.SetOut(&buf)
	root.SetErr(&buf)
	root.SetArgs([]string{"completion"})

	err := root.Execute()
	if err == nil {
		t.Error("expected error for missing shell argument, got nil")
	}
}
