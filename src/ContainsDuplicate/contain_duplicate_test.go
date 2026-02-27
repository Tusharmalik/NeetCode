package main

import "testing"

func TestContainDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "Test Case 1",
			nums: []int{1, 2, 3, 1},
			want: true,
		},
		{
			name: "Test Case 2",
			nums: []int{1, 2, 3, 4},
			want: false,
		},
		{
			name: "Test Case 3",
			nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			want: true,
		},
		{
			name: "Test Case 4",
			nums: []int{1, 2, 3, 3},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hasDuplicate(tt.nums)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
