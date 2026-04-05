package main

import (
	"slices"
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name:     "Test 1",
			nums:     []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{{-1, 0, 1}, {-1, -1, 2}},
		},
		{
			name:     "Test 2",
			nums:     []int{0, 1, 1},
			expected: [][]int{},
		},
		{
			name:     "Test 3",
			nums:     []int{0, 0, 0},
			expected: [][]int{{0, 0, 0}},
		},
		{
			name:     "Test 4",
			nums:     []int{-1, 0, 1, 2, -1, -4, -1, 0, 1, 2, -1, -4, -1, 0, 1, 2, -1, -4},
			expected: [][]int{{-4, 2, 2}, {-1, -1, 2}, {-1, 0, 1}, {0, 0, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name+" Brute Force", func(t *testing.T) {
			got := threeSumBruteForce(tt.nums)
			AssertEqual(t, got, tt.expected)
		})

		t.Run(tt.name+" Hash Map", func(t *testing.T) {
			got := threeSumHashMap(tt.nums)
			AssertEqual(t, got, tt.expected)
		})

		t.Run(tt.name+" Pointer", func(t *testing.T) {
			got := threeSumTwoPointer(tt.nums)
			AssertEqual(t, got, tt.expected)
		})
	}
}

// sortOuter sorts a 2D Integer slice in lexicographic order.
//
// Sorting strategy:
// 1. Compares slices element by element (lexicographically)
// 2. If all compared elements are equal, the shorter slice comes first
func sortOuter(s [][]int) {
	sort.Slice(s, func(i, j int) bool {
		// Get the two slices to compare
		sa := s[i]
		sb := s[j]

		// Find the minimum length between the two slices
		// We'll compare elements up to this length
		n := len(sa)
		if n > len(sb) {
			n = len(sb)
		}

		// Compare each element lexicographically
		for k := 0; k < n; k++ {
			if sa[k] != sb[k] {
				// If elements differ, return true if sa[k] is less than sb[k]
				return sa[k] < sb[k]
			}
		}

		// If all compared elements are equal, the shorter slice comes first
		return len(sa) < len(sb)
	})
}

func AssertEqual(t *testing.T, got, want [][]int) {
	t.Helper()
	if len(got) != len(want) {
		t.Errorf("got: %v, want: %v", got, want)
	}

	for _, slice := range got {
		sort.Ints(slice)
	}

	for _, slice := range want {
		sort.Ints(slice)
	}

	sortOuter(got)
	sortOuter(want)

	if !slices.EqualFunc(got, want, func(s1, s2 []int) bool {
		return slices.Equal(s1, s2)
	}) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func BenchmarkThreeSumBruteForce(b *testing.B) {
	input := []int{-1, 0, 1, 2, -1, -4, -1, 0, 1, 2, -1, -4, -1, 0, 1, 2, -1, -4}
	for b.Loop() {
		threeSumBruteForce(input)
	}
}

func BenchmarkThreeSumHashMap(b *testing.B) {
	input := []int{-1, 0, 1, 2, -1, -4, -1, 0, 1, 2, -1, -4, -1, 0, 1, 2, -1, -4}
	for b.Loop() {
		threeSumHashMap(input)
	}
}

func BenchmarkThreeSumTwoPointers(b *testing.B) {
	input := []int{-1, 0, 1, 2, -1, -4, -1, 0, 1, 2, -1, -4, -1, 0, 1, 2, -1, -4}
	for b.Loop() {
		threeSumTwoPointer(input)
	}
}
