package main

import "testing"

func TestValidPalindrome(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "Test Case 1",
			input: "abc",
			want:  false,
		},
		{
			name:  "Test Case 2",
			input: "",
			want:  true,
		},
		{
			name:  "Test Case 3",
			input: "Was it a car or a cat I saw?",
			want:  true,
		},
		{
			name:  "Test Case 4",
			input: "tab a cat",
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidPalindrome(tt.input); got != tt.want {
				t.Errorf("ValidPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
