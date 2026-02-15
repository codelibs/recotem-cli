package utils

import (
	"testing"

	"recotem.org/cli/recotem/pkg/openapi"
)

func TestNilOrString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *string
	}{
		{"empty string", "", nil},
		{"non-empty string", "test", stringPtr("test")},
		{"space string", " ", stringPtr(" ")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NilOrString(tt.input)
			if (result == nil && tt.expected != nil) || (result != nil && tt.expected == nil) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
			if result != nil && tt.expected != nil && *result != *tt.expected {
				t.Errorf("expected %v, got %v", *tt.expected, *result)
			}
		})
	}
}

func TestNilOrInt(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *int
	}{
		{"empty string", "", nil},
		{"invalid int", "abc", nil},
		{"valid int", "123", intPtr(123)},
		{"negative int", "-456", intPtr(-456)},
		{"zero", "0", intPtr(0)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NilOrInt(tt.input)
			if (result == nil && tt.expected != nil) || (result != nil && tt.expected == nil) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
			if result != nil && tt.expected != nil && *result != *tt.expected {
				t.Errorf("expected %v, got %v", *tt.expected, *result)
			}
		})
	}
}

func TestNilOrFloat32(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *float32
	}{
		{"empty string", "", nil},
		{"invalid float", "abc", nil},
		{"valid float", "123.45", float32Ptr(123.45)},
		{"negative float", "-67.89", float32Ptr(-67.89)},
		{"zero", "0", float32Ptr(0)},
		{"integer as float", "42", float32Ptr(42)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NilOrFloat32(tt.input)
			if (result == nil && tt.expected != nil) || (result != nil && tt.expected == nil) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
			if result != nil && tt.expected != nil && *result != *tt.expected {
				t.Errorf("expected %v, got %v", *tt.expected, *result)
			}
		})
	}
}

func TestNilOrBool(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *bool
	}{
		{"empty string", "", nil},
		{"invalid bool", "abc", nil},
		{"true", "true", boolPtr(true)},
		{"false", "false", boolPtr(false)},
		{"1", "1", boolPtr(true)},
		{"0", "0", boolPtr(false)},
		{"T", "T", boolPtr(true)},
		{"F", "F", boolPtr(false)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NilOrBool(tt.input)
			if (result == nil && tt.expected != nil) || (result != nil && tt.expected == nil) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
			if result != nil && tt.expected != nil && *result != *tt.expected {
				t.Errorf("expected %v, got %v", *tt.expected, *result)
			}
		})
	}
}

func TestNilOrScheme(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *openapi.SchemeEnum
	}{
		{"empty string", "", nil},
		{"RG", "RG", schemePtr(openapi.RG)},
		{"TG", "TG", schemePtr(openapi.TG)},
		{"TU", "TU", schemePtr(openapi.TU)},
		{"invalid scheme", "INVALID", nil},
		{"lowercase rg", "rg", nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NilOrScheme(tt.input)
			if (result == nil && tt.expected != nil) || (result != nil && tt.expected == nil) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
			if result != nil && tt.expected != nil && *result != *tt.expected {
				t.Errorf("expected %v, got %v", *tt.expected, *result)
			}
		})
	}
}

func TestNilOrTargetMetric(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *openapi.TargetMetricEnum
	}{
		{"empty string", "", nil},
		{"hit", "hit", targetMetricPtr(openapi.Hit)},
		{"map", "map", targetMetricPtr(openapi.Map)},
		{"recall", "recall", targetMetricPtr(openapi.Recall)},
		{"ndcg", "ndcg", targetMetricPtr(openapi.Ndcg)},
		{"invalid metric", "INVALID", nil},
		{"uppercase HIT", "HIT", nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NilOrTargetMetric(tt.input)
			if (result == nil && tt.expected != nil) || (result != nil && tt.expected == nil) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
			if result != nil && tt.expected != nil && *result != *tt.expected {
				t.Errorf("expected %v, got %v", *tt.expected, *result)
			}
		})
	}
}

// Helper functions
func stringPtr(s string) *string {
	return &s
}

func intPtr(i int) *int {
	return &i
}

func float32Ptr(f float32) *float32 {
	return &f
}

func boolPtr(b bool) *bool {
	return &b
}

func schemePtr(s openapi.SchemeEnum) *openapi.SchemeEnum {
	return &s
}

func targetMetricPtr(t openapi.TargetMetricEnum) *openapi.TargetMetricEnum {
	return &t
}
