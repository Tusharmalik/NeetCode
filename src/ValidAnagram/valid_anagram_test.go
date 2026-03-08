package main

import "testing"

func TestValidAnagram(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{
			name: "Test Case 1",
			s:    "anagram",
			t:    "nagaram",
			want: true,
		},
		{
			name: "Test Case 2",
			s:    "rat",
			t:    "car",
			want: false,
		},
		{
			name: "Test Case 3",
			s:    "racecar",
			t:    "carrace",
			want: true,
		},
		{
			name: "Test Case 4",
			s:    "jar",
			t:    "jam",
			want: false,
		},
		{
			name: "Test Case 5",
			s:    "",
			t:    "",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isAnagram(tt.s, tt.t)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
