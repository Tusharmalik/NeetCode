package main

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "Test 1",
			input: "abcabcbb",
			want:  3,
		},
		{
			name:  "Test 2",
			input: "bbbbb",
			want:  1,
		},
		{
			name:  "Test 3",
			input: "pwwkew",
			want:  3,
		},
		{
			name:  "Test 4",
			input: " ",
			want:  1,
		},
		{
			name:  "Test 5",
			input: "dvdf",
			want:  3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name+" Brute Force", func(t *testing.T) {
			got := lengthOfLongestSubstringBruteForce(tt.input)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name+" Sliding Window", func(t *testing.T) {
			got := lengthOfLongestSubstringSlidingWindow(tt.input)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name+" Sliding Window Optimal", func(t *testing.T) {
			got := lengthOfLongestSubstringSlidingWindowOptimal(tt.input)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkLengthOfLongestSubstring(b *testing.B) {
	input := "abcabcbb"

	b.Run("Brute Force", func(b *testing.B) {
		b.ResetTimer()
		for b.Loop() {
			lengthOfLongestSubstringBruteForce(input)
		}
	})
	b.Run("Sliding Window", func(b *testing.B) {
		b.ResetTimer()
		for b.Loop() {
			lengthOfLongestSubstringSlidingWindow(input)
		}
	})
	b.Run("Sliding Window Optimal", func(b *testing.B) {
		b.ResetTimer()
		for b.Loop() {
			lengthOfLongestSubstringSlidingWindowOptimal(input)
		}
	})
}
