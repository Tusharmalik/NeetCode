package main

import "testing"

func TestEncodeDecode(t *testing.T) {
	tests := []struct {
		name  string
		input []string
	}{
		{
			name:  "Test Case 1",
			input: []string{"Go💙Go💙", "Go💙"},
		},
		{
			name:  "Test Case 2",
			input: []string{"Go💙"},
		},
		{
			name:  "Test Case 3",
			input: []string{"Hello", "World"},
		},
		{
			name:  "Test Case 4",
			input: []string{"", "Empty String", "Another String"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solution := Solution{}
			encoded := solution.Encode(tt.input)
			decoded := solution.Decode(encoded)

			if !equal(decoded, tt.input) {
				t.Errorf("got %v, want %v", decoded, tt.input)
			}
		})
	}
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
