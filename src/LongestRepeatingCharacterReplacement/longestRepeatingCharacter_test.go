package main

import "testing"

func TestLongestRepeatingCharacterReplacement(t *testing.T) {
	tests := []struct {
		name    string
		text    string
		replace int
		want    int
	}{
		{
			name:    "test 1",
			text:    "XYYX",
			replace: 2,
			want:    4,
		},
		{
			name:    "test 2",
			text:    "AAABABB",
			replace: 1,
			want:    5,
		},
		{
			name:    "test 3",
			text:    "ABBB",
			replace: 2,
			want:    4,
		},
		{
			name:    "test 4",
			text:    "AABABBB",
			replace: 1,
			want:    5,
		},
		{
			name:    "test 5",
			text:    "ABCDE",
			replace: 5,
			want:    5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := characterReplacement(tt.text, tt.replace)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
