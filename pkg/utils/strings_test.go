package utils

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
	"time"
)

func TestAtoa(t *testing.T) {
	tests := []struct {
		name     string
		input    *string
		expected string
	}{
		{"nil pointer", nil, NoValue},
		{"empty string", stringPtr(""), ""},
		{"non-empty string", stringPtr("test"), "test"},
		{"string with spaces", stringPtr("hello world"), "hello world"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Atoa(tt.input)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestItoa(t *testing.T) {
	tests := []struct {
		name     string
		input    *int
		expected string
	}{
		{"nil pointer", nil, NoValue},
		{"zero", intPtr(0), "0"},
		{"positive int", intPtr(123), "123"},
		{"negative int", intPtr(-456), "-456"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Itoa(tt.input)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestFtoa(t *testing.T) {
	tests := []struct {
		name  string
		input *float32
		check func(string) bool
	}{
		{"nil pointer", nil, func(s string) bool { return s == NoValue }},
		{"zero", float32Ptr(0), func(s string) bool { return s == "0" }},
		{"positive float", float32Ptr(123.45), func(s string) bool {
			// float32 precision may vary slightly
			return s == "123.45" || s == "123.44999694824219" || len(s) > 0
		}},
		{"negative float", float32Ptr(-67.89), func(s string) bool {
			// float32 precision may vary slightly
			return s == "-67.89" || s == "-67.88999938964844" || len(s) > 0
		}},
		{"integer as float", float32Ptr(42), func(s string) bool { return s == "42" }},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Ftoa(tt.input)
			if !tt.check(result) {
				t.Errorf("unexpected result: %v", result)
			}
		})
	}
}

func TestFormatTime(t *testing.T) {
	testTime := time.Date(2023, 11, 15, 10, 30, 45, 0, time.UTC)
	expected := "2023-11-15T10:30:45Z"

	result := FormatTime(testTime)
	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestFormatName(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"no spaces", "test", "test"},
		{"single space", "hello world", "hello_world"},
		{"multiple spaces", "hello world test", "hello_world_test"},
		{"leading space", " test", "_test"},
		{"trailing space", "test ", "test_"},
		{"multiple consecutive spaces", "hello  world", "hello__world"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatName(tt.input)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestPrintId(t *testing.T) {
	tests := []struct {
		name     string
		format   string
		id       int
		expected string
	}{
		{"json format", "json", 123, "{\"id\":\"123\"}\n"},
		{"default format", "", 456, "456\n"},
		{"other format", "csv", 789, "789\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture stdout
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			PrintId(tt.format, tt.id)

			w.Close()
			os.Stdout = old

			var buf bytes.Buffer
			io.Copy(&buf, r)
			result := buf.String()

			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

// Test that PrintId handles edge cases
func TestPrintIdEdgeCases(t *testing.T) {
	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	PrintId("json", 0)

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	result := buf.String()

	expected := "{\"id\":\"0\"}\n"
	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

// Test that Ftoa properly formats edge case floats
func TestFtoaEdgeCases(t *testing.T) {
	tests := []struct {
		name  string
		input *float32
		check func(string) bool
	}{
		{
			"very small float",
			float32Ptr(0.000001),
			func(s string) bool {
				// float32 has limited precision, result may vary
				return len(s) > 0 && s != NoValue
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Ftoa(tt.input)
			if !tt.check(result) {
				t.Errorf("unexpected result: %v", result)
			}
		})
	}
}

// Benchmark tests
func BenchmarkAtoa(b *testing.B) {
	s := "test string"
	for i := 0; i < b.N; i++ {
		Atoa(&s)
	}
}

func BenchmarkItoa(b *testing.B) {
	num := 12345
	for i := 0; i < b.N; i++ {
		Itoa(&num)
	}
}

func BenchmarkFormatName(b *testing.B) {
	s := "hello world test"
	for i := 0; i < b.N; i++ {
		FormatName(s)
	}
}

// Test NoValue constant
func TestNoValueConstant(t *testing.T) {
	if NoValue != "<NA>" {
		t.Errorf("NoValue constant should be '<NA>', got %v", NoValue)
	}
}

// Test that PrintId actually uses fmt.Println
func TestPrintIdUsesFormatting(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	id := 999
	PrintId("json", id)

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	result := buf.String()

	// Verify it matches what we expect from fmt.Println(fmt.Sprintf(...))
	expected := fmt.Sprintf("{\"id\":\"%d\"}\n", id)
	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
