package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// PrintOutput prints the given value in the specified format
func PrintOutput(format string, v any) {
	switch strings.ToLower(format) {
	case "json":
		printJSON(v)
	case "yaml":
		printYAML(v)
	default:
		printText(v)
	}
}

// PrintList prints a list of values in the specified format
func PrintList(format string, items []map[string]any) {
	switch strings.ToLower(format) {
	case "json":
		printJSON(items)
	case "yaml":
		printYAML(items)
	default:
		for _, item := range items {
			printTextMap(item)
		}
	}
}

func printJSON(v any) {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(v); err != nil {
		fmt.Fprintf(os.Stderr, "JSON encoding error: %v\n", err)
	}
}

func printYAML(v any) {
	enc := yaml.NewEncoder(os.Stdout)
	enc.SetIndent(2)
	if err := enc.Encode(v); err != nil {
		fmt.Fprintf(os.Stderr, "YAML encoding error: %v\n", err)
	}
}

func printText(v any) {
	switch val := v.(type) {
	case map[string]any:
		printTextMap(val)
	case []map[string]any:
		for _, item := range val {
			printTextMap(item)
		}
	default:
		fmt.Println(v)
	}
}

func printTextMap(m map[string]any) {
	parts := make([]string, 0, len(m))
	for _, v := range m {
		parts = append(parts, fmt.Sprintf("%v", v))
	}
	fmt.Println(strings.Join(parts, "\t"))
}

// ToMap converts a struct-like value to a map for output
func ToMap(pairs ...any) map[string]any {
	m := make(map[string]any)
	for i := 0; i+1 < len(pairs); i += 2 {
		key, ok := pairs[i].(string)
		if !ok {
			continue
		}
		m[key] = pairs[i+1]
	}
	return m
}
