package main

import "testing"

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "Test Case 1",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			name:   "Test Case 2",
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
		{
			name:   "Test Case 3",
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
		{
			name:   "Test Case 4",
			nums:   []int{1, 2, 3, 4, 5},
			target: 10,
			want:   nil,
		},
		{
			name:   "Test Case 5",
			nums:   []int{1},
			target: 1,
			want:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.nums, tt.target)
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
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
