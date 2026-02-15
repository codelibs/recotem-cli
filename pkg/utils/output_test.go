package utils

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

func TestPrintOutputJSON(t *testing.T) {
	data := map[string]any{"name": "test", "value": 42}
	output := captureStdout(t, func() {
		PrintOutput("json", data)
	})

	if !strings.Contains(output, `"name": "test"`) {
		t.Errorf("expected JSON with name field, got %s", output)
	}
	if !strings.Contains(output, `"value": 42`) {
		t.Errorf("expected JSON with value field, got %s", output)
	}
}

func TestPrintOutputYAML(t *testing.T) {
	data := map[string]any{"name": "test"}
	output := captureStdout(t, func() {
		PrintOutput("yaml", data)
	})

	if !strings.Contains(output, "name: test") {
		t.Errorf("expected YAML output, got %s", output)
	}
}

func TestPrintOutputText(t *testing.T) {
	output := captureStdout(t, func() {
		PrintOutput("text", "hello world")
	})

	if strings.TrimSpace(output) != "hello world" {
		t.Errorf("expected 'hello world', got %q", output)
	}
}

func TestPrintOutputCaseInsensitive(t *testing.T) {
	data := map[string]any{"k": "v"}
	output := captureStdout(t, func() {
		PrintOutput("JSON", data)
	})

	if !strings.Contains(output, `"k": "v"`) {
		t.Errorf("expected JSON output for uppercase format, got %s", output)
	}
}

func TestPrintListJSON(t *testing.T) {
	items := []map[string]any{
		{"id": 1, "name": "a"},
		{"id": 2, "name": "b"},
	}
	output := captureStdout(t, func() {
		PrintList("json", items)
	})

	if !strings.Contains(output, `"name": "a"`) {
		t.Errorf("expected JSON list output, got %s", output)
	}
}

func TestPrintListText(t *testing.T) {
	items := []map[string]any{
		{"id": 1},
		{"id": 2},
	}
	output := captureStdout(t, func() {
		PrintList("text", items)
	})

	lines := strings.Split(strings.TrimSpace(output), "\n")
	if len(lines) != 2 {
		t.Errorf("expected 2 lines, got %d: %q", len(lines), output)
	}
}

func TestToMap(t *testing.T) {
	m := ToMap("name", "test", "value", 42)

	if m["name"] != "test" {
		t.Errorf("expected name=test, got %v", m["name"])
	}
	if m["value"] != 42 {
		t.Errorf("expected value=42, got %v", m["value"])
	}
}

func TestToMapOddArgs(t *testing.T) {
	// Odd number of args: last key is ignored
	m := ToMap("a", 1, "b")
	if m["a"] != 1 {
		t.Errorf("expected a=1, got %v", m["a"])
	}
	if _, ok := m["b"]; ok {
		t.Error("expected 'b' to be missing with odd args")
	}
}

func TestPrintOutputTextMap(t *testing.T) {
	data := map[string]any{"col1": "val1", "col2": "val2"}
	output := captureStdout(t, func() {
		PrintOutput("text", data)
	})

	if !strings.Contains(output, "val1") || !strings.Contains(output, "val2") {
		t.Errorf("expected map values in text output, got %s", output)
	}
}
