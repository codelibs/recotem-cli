package cmd

import "testing"

func TestLogoutCmdStructure(t *testing.T) {
	cmd := newLogoutCmd()

	if cmd.Use != "logout" {
		t.Errorf("expected Use %q, got %q", "logout", cmd.Use)
	}
	if cmd.Short != "Clear authentication tokens" {
		t.Errorf("expected Short %q, got %q", "Clear authentication tokens", cmd.Short)
	}
	if cmd.RunE == nil {
		t.Error("expected RunE to be set")
	}
}
