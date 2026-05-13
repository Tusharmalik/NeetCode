package main

import "testing"

func TestMaximumSlidingWindow(t *testing.T) {
	tests := []struct {
		name       string
		inputArray []int
		inputSize  int
		want       []int
	}{
		{
			name:       "test 1",
			inputArray: []int{1, 2, 1, 0, 4, 2, 6},
			inputSize:  3,
			want:       []int{2, 2, 4, 4, 6},
		},
		{
			name:       "test 2",
			inputArray: []int{1, -1},
			inputSize:  1,
			want:       []int{1, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name+" Brute Force", func(t *testing.T) {
			got := maxSlidingWindowBruteForce(tt.inputArray, tt.inputSize)
			if !equal(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})

		t.Run(tt.name+" Dynamic Programming", func(t *testing.T) {
			got := maxSlidingWindowDynamicProgramming(tt.inputArray, tt.inputSize)
			if !equal(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})

		t.Run(tt.name+" Hash Segment Tree", func(t *testing.T) {
			got := maxSlidingWindowHashSegmentTree(tt.inputArray, tt.inputSize)
			if !equal(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})

		t.Run(tt.name+" Segment Tree", func(t *testing.T) {
			got := maxSlidingWindowSegmentTree(tt.inputArray, tt.inputSize)
			if !equal(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func BenchmarkMaxSlidingWindow(b *testing.B) {
	inputArray := []int{1, 2, 1, 0, 4, 2, 6, 1, 2, 1, 0, 4, 2, 6, 1, 2, 1, 0, 4, 2, 6, 1, 2, 1, 0, 4, 2, 6}
	inputSize := 3

	b.Run("Brute Force", func(b *testing.B) {
		b.ResetTimer()
		for b.Loop() {
			maxSlidingWindowBruteForce(inputArray, inputSize)
		}
	})

	b.Run("Dynamic Programming", func(b *testing.B) {
		b.ResetTimer()
		for b.Loop() {
			maxSlidingWindowDynamicProgramming(inputArray, inputSize)
		}
	})

	b.Run("Hash Segment Tree", func(b *testing.B) {
		b.ResetTimer()
		for b.Loop() {
			maxSlidingWindowHashSegmentTree(inputArray, inputSize)
		}
	})

	b.Run("Segment Tree", func(b *testing.B) {
		b.ResetTimer()
		for b.Loop() {
			maxSlidingWindowSegmentTree(inputArray, inputSize)
		}
	})
}
