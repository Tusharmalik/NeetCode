package main

import "testing"

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "Test Case 1",
			nums: []int{1, 2, 3, 4},
			want: []int{24, 12, 8, 6},
		},
		{
			name: "Test Case 2",
			nums: []int{-1, 1, 0, -3, 3},
			want: []int{0, 0, 9, 0, 0},
		},
		{
			name: "Test Case 3",
			nums: []int{1, 2, 4, 6},
			want: []int{48, 24, 12, 8},
		},
		{
			name: "Test Case 4",
			nums: []int{-1, 0, 1, 2, 3},
			want: []int{0, -6, 0, 0, 0},
		},
		{
			name: "Test Case 5",
			nums: []int{0, 0, 1, 2, 3},
			want: []int{0, 0, 0, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name+" PrefixSuffixOptimised", func(t *testing.T) {
			got := productExceptSelfPrefixSuffixOptimised(tt.nums)
			AssertCorrectMessage(t, got, tt.want)
		})
		t.Run(tt.name+" PrefixSuffixUnoptimised", func(t *testing.T) {
			got := productExceptSelfPrefixSuffixUnoptimised(tt.nums)
			AssertCorrectMessage(t, got, tt.want)
		})
		t.Run(tt.name+" Unoptimised", func(t *testing.T) {
			got := productExceptSelfUnoptimized(tt.nums)
			AssertCorrectMessage(t, got, tt.want)
		})
		t.Run(tt.name+" Divide", func(t *testing.T) {
			got := productExceptSelfDivide(tt.nums)
			AssertCorrectMessage(t, got, tt.want)
		})
	}
}

func AssertCorrectMessage(t *testing.T, got, want []int) {
	t.Helper()
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("got %v, want %v", got, want)
			return
		}
	}
}

func BenchmarkProductExceptSelfPrefixSuffixOptimised(b *testing.B) {
	totalElements := 10000
	input := make([]int, 0, totalElements)
	for i := 1; i <= totalElements; i++ {
		input = append(input, i)
	}
	for b.Loop() {
		productExceptSelfPrefixSuffixOptimised(input)
	}
}

func BenchmarkProductExceptSelfPrefixSuffixUnoptimised(b *testing.B) {
	totalElements := 10000
	input := make([]int, 0, totalElements)
	for i := 1; i <= totalElements; i++ {
		input = append(input, i)
	}
	for b.Loop() {
		productExceptSelfPrefixSuffixUnoptimised(input)
	}
}

func BenchmarkProductExceptSelfUnoptimized(b *testing.B) {
	totalElements := 10000
	input := make([]int, 0, totalElements)
	for i := 1; i <= totalElements; i++ {
		input = append(input, i)
	}
	for b.Loop() {
		productExceptSelfUnoptimized(input)
	}
}

func BenchmarkProductExceptSelfDivide(b *testing.B) {
	totalElements := 10000
	input := make([]int, 0, totalElements)
	for i := 1; i <= totalElements; i++ {
		input = append(input, i)
	}
	for b.Loop() {
		productExceptSelfDivide(input)
	}
}
