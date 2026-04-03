package main

import (
	"sort"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		target int
		want   []int
	}{
		{
			name:   "Test 1",
			input:  []int{2, 7, 11, 15},
			target: 9,
			want:   []int{1, 2},
		},
		{
			name:   "Test 2",
			input:  []int{1, 2, 3, 4},
			target: 6,
			want:   []int{2, 4},
		},
		{
			name:   "Test 3",
			input:  []int{1, 2, 3, 4, 5},
			target: 10,
			want:   []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name+" Brute Force", func(t *testing.T) {
			got := twoSumBruteForce(tt.input, tt.target)
			AssertEqual(t, got, tt.want)
		})

		t.Run(tt.name+" Binary Search", func(t *testing.T) {
			got := twoSumBinarySearch(tt.input, tt.target)
			AssertEqual(t, got, tt.want)
		})

		t.Run(tt.name+" Hash Map", func(t *testing.T) {
			got := twoSumHashMap(tt.input, tt.target)
			AssertEqual(t, got, tt.want)
		})

		t.Run(tt.name+" Two Pointer", func(t *testing.T) {
			got := twoSumTwoPointer(tt.input, tt.target)
			AssertEqual(t, got, tt.want)
		})
	}
}

func AssertEqual(t *testing.T, got, want []int) {
	t.Helper()
	if !equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	sort.Ints(a)
	sort.Ints(b)

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func BenchmarkTwoSumBruteForce(b *testing.B) {
	input := []int{1, 2, 3, 4, 5}
	target := 6

	for b.Loop() {
		twoSumBruteForce(input, target)
	}
}

func BenchmarkTwoSumBinarySearch(b *testing.B) {
	input := []int{1, 2, 3, 4, 5}
	target := 6

	for b.Loop() {
		twoSumBinarySearch(input, target)
	}
}

func BenchmarkTwoSumHashMap(b *testing.B) {
	input := []int{1, 2, 3, 4, 5}
	target := 6

	for b.Loop() {
		twoSumHashMap(input, target)
	}
}

func BenchmarkTwoSumTwoPointer(b *testing.B) {
	input := []int{1, 2, 3, 4, 5}
	target := 6

	for b.Loop() {
		twoSumTwoPointer(input, target)
	}
}
