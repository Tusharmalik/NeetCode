package main

import "testing"

func TestTrapWater(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "test 1",
			input: []int{0, 2, 0, 3, 1, 0, 1, 3, 2, 1},
			want:  9,
		},
		{
			name:  "test 2",
			input: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			want:  6,
		},
		{
			name:  "test 3",
			input: []int{4, 2, 0, 3, 2, 5},
			want:  9,
		},
		{
			name:  "test 4",
			input: []int{0},
			want:  0,
		},
		{
			name:  "test 5",
			input: []int{5, 4, 1, 2},
			want:  1,
		},
		{
			name:  "test 6",
			input: []int{9, 6, 8, 8, 5, 6, 3},
			want:  3,
		},
		{
			name:  "test 7",
			input: []int{5},
			want:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name+" Brute Force", func(t *testing.T) {
			got := trapBruteForce(tt.input)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})

		t.Run(tt.name+" Prefix Suffix", func(t *testing.T) {
			got := trapPrefixSuffix(tt.input)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})

		t.Run(tt.name+" Two Pointer", func(t *testing.T) {
			got := trapTwoPointer(tt.input)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkTrapWater(b *testing.B) {
	input := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}

	b.Run("Brute Force", func(b *testing.B) {
		b.ResetTimer()
		for b.Loop() {
			trapBruteForce(input)
		}
	})

	b.Run("Prefix Suffix", func(b *testing.B) {
		b.ResetTimer()
		for b.Loop() {
			trapPrefixSuffix(input)
		}
	})

	b.Run("Two Pointer", func(b *testing.B) {
		b.ResetTimer()
		for b.Loop() {
			trapTwoPointer(input)
		}
	})
}
