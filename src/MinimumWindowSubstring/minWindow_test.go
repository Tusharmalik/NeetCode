package main

import "testing"

func TestMinWindow(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want string
	}{
		{
			name: "Test 1",
			s:    "OUZODYXAZV",
			t:    "XYZ",
			want: "YXAZ",
		},
		{
			name: "Test 2",
			s:    "xyz",
			t:    "xyz",
			want: "xyz",
		},
		{
			name: "Test 3",
			s:    "x",
			t:    "xy",
			want: "",
		},
		{
			name: "Test 4",
			s:    "ADOBECODEBANC",
			t:    "ABC",
			want: "BANC",
		},
		{
			name: "Test 5",
			s:    "a",
			t:    "a",
			want: "a",
		},
		{
			name: "Test 6",
			s:    "a",
			t:    "aa",
			want: "",
		},
		{
			name: "Test 7",
			s:    "aaaaaaaaaaaabbbbbcdd",
			t:    "abcdd",
			want: "abbbbbcdd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name+" Brute Force", func(t *testing.T) {
			got := minWindowBruteForce(tt.s, tt.t)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name+" Sliding Window", func(t *testing.T) {
			got := minWindowSlidingWindow(tt.s, tt.t)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkMinWindow(b *testing.B) {
	inputS := "OUZODYXAZV"
	inputT := "XYZ"

	b.Run("Brute Force", func(b *testing.B) {
		b.ResetTimer()
		for b.Loop() {
			minWindowBruteForce(inputS, inputT)
		}
	})
	b.Run("Sliding Window", func(b *testing.B) {
		b.ResetTimer()
		for b.Loop() {
			minWindowSlidingWindow(inputS, inputT)
		}
	})
}
