package cmd

import "testing"

func TestParseIntList(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expected  []int
		expectErr bool
	}{
		{
			name:     "single value",
			input:    "1",
			expected: []int{1},
		},
		{
			name:     "multiple values",
			input:    "1,2,3",
			expected: []int{1, 2, 3},
		},
		{
			name:     "values with spaces",
			input:    "1, 2, 3",
			expected: []int{1, 2, 3},
		},
		{
			name:     "trailing comma",
			input:    "1,2,",
			expected: []int{1, 2},
		},
		{
			name:     "leading comma",
			input:    ",1,2",
			expected: []int{1, 2},
		},
		{
			name:      "invalid value",
			input:     "1,abc,3",
			expectErr: true,
		},
		{
			name:      "float value",
			input:     "1,2.5,3",
			expectErr: true,
		},
		{
			name:     "empty string",
			input:    "",
			expected: []int{},
		},
		{
			name:     "negative values",
			input:    "-1,-2",
			expected: []int{-1, -2},
		},
		{
			name:     "zero",
			input:    "0",
			expected: []int{0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parseIntList(tt.input)
			if tt.expectErr {
				if err == nil {
					t.Errorf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(result) != len(tt.expected) {
				t.Fatalf("expected %d items, got %d: %v", len(tt.expected), len(result), result)
			}
			for i, v := range result {
				if v != tt.expected[i] {
					t.Errorf("at index %d: expected %d, got %d", i, tt.expected[i], v)
				}
			}
		})
	}
}
