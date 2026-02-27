package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want [][]string
	}{
		{
			name: "Test Case 1",
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},
		{
			name: "Test Case 2",
			strs: []string{"act", "pots", "tops", "cat", "stop", "hat"},
			want: [][]string{{"act", "cat"}, {"pots", "stop", "tops"}, {"hat"}},
		},
		{
			name: "Test Case 3",
			strs: []string{"abct", "cat"},
			want: [][]string{{"cat"}, {"abct"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := groupAnagrams(tt.strs)
			AssertCorrectMessage(t, got, tt.want)
		})
	}
}

// sortOuter sorts a 2D string slice in lexicographic order.
//
// Sorting strategy:
// 1. Compares slices element by element (lexicographically)
// 2. If all compared elements are equal, the shorter slice comes first
func sortOuter(s [][]string) {
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

func AssertCorrectMessage(t testing.TB, got, want [][]string) {
	t.Helper()

	for _, s := range got {
		sort.Strings(s)
	}

	for _, s := range want {
		sort.Strings(s)
	}

	sortOuter(got)
	sortOuter(want)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}
