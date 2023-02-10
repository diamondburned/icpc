package main

import "testing"

func TestSummer(t *testing.T) {
	type testCase struct {
		i, j int
		sum  int
	}

	type test struct {
		elems []int
		cases []testCase
	}

	tests := []test{
		{
			elems: []int{1, 2, 3, 4, 5},
			cases: []testCase{
				{0, 5, 15},
				{0, 4, 10},
				{1, 5, 14},
			},
		},
	}

	for _, tt := range tests {
		s := NewSummer(tt.elems)
		for _, tc := range tt.cases {
			if got := s.Sum(tc.i, tc.j); got != tc.sum {
				t.Errorf("sum(%d, %d) = %d, want %d", tc.i, tc.j, got, tc.sum)
			}
		}
	}
}
