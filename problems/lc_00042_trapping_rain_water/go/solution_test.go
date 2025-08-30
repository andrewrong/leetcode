package lc_00042_trapping_rain_water

import "testing"

func TestTrap(t *testing.T) {
	testCases := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "Example 1",
			height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			want:   6,
		},
		{
			name:   "Example 2",
			height: []int{4, 2, 0, 3, 2, 5},
			want:   9,
		},
		{
			name:   "Empty height",
			height: []int{},
			want:   0,
		},
		{
			name:   "No trapping",
			height: []int{1, 2, 3, 4, 5},
			want:   0,
		},
		{
			name:   "Single element",
			height: []int{5},
			want:   0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := trapDymanic(tc.height)
			if got != tc.want {
				t.Errorf("trap(%v) = %v; want %v", tc.height, got, tc.want)
			}
		})
	}
}
