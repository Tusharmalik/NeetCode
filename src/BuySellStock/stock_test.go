package main

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "Test 1",
			input: []int{10, 1, 5, 6, 7, 1},
			want:  6,
		},
		{
			name:  "Test 2",
			input: []int{10, 8, 7, 5, 2},
			want:  0,
		},
		{
			name:  "Test 3",
			input: []int{7, 1, 5, 3, 6, 4},
			want:  5,
		},
		{
			name:  "Test 4",
			input: []int{5, 10, 2, 8, 7, 5, 2},
			want:  6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name+" Brute Force", func(t *testing.T) {
			got := maxProfitBruteForce(tt.input)
			if got != tt.want {
				t.Errorf("maxProfitBruteForce(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
		t.Run(tt.name+" Two Pointer", func(t *testing.T) {
			got := maxProfitTwoPointer(tt.input)
			if got != tt.want {
				t.Errorf("maxProfitTwoPointer(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
		t.Run(tt.name+" Dynamic Programming", func(t *testing.T) {
			got := maxProfitDynamicProgramming(tt.input)
			if got != tt.want {
				t.Errorf("maxProfitDynamicProgramming(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func BenchmarkMaxProfit(b *testing.B) {
	input := []int{7, 1, 5, 3, 6, 4}

	b.Run("Brute Force", func(b *testing.B) {
		b.ResetTimer()
		for b.Loop() {
			maxProfitBruteForce(input)
		}
	})
	b.Run("Two Pointer", func(b *testing.B) {
		b.ResetTimer()
		for b.Loop() {
			maxProfitTwoPointer(input)
		}
	})
	b.Run("Dynamic Programming", func(b *testing.B) {
		b.ResetTimer()
		for b.Loop() {
			maxProfitDynamicProgramming(input)
		}
	})
}
