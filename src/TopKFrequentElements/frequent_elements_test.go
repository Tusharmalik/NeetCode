package main

import (
	"sort"
	"testing"
)

func TestFrequentElements(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{
			name: "Test Case 1",
			nums: []int{1, 1, 1, 2, 2, 3},
			k:    2,
			want: []int{1, 2},
		},
		{
			name: "Test Case 2",
			nums: []int{7, 7},
			k:    1,
			want: []int{7},
		},
		{
			name: "Test Case 3",
			nums: []int{1, 1, 2, 2, 3, 3, 4, 5},
			k:    3,
			want: []int{3, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := topKFrequent(tt.nums, tt.k)
			AssertCorrectMessage(t, got, tt.want)
		})
	}
}

func AssertCorrectMessage(t testing.TB, got, want []int) {
	t.Helper()
	if !equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
