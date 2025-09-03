package lc_00077_combinations

import (
	"reflect"
	"testing"
)

func TestCombine(t *testing.T) {
	tests := []struct {
		n      int
		k      int
		expect [][]int
	}{
		{
			n:      1,
			k:      1,
			expect: [][]int{{1}},
		},
		{
			n:      1,
			k:      2,
			expect: [][]int{},
		},
		{
			n:      2,
			k:      1,
			expect: [][]int{{1}, {2}},
		},
		{
			n:      4,
			k:      2,
			expect: [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}},
		},
		{
			n:      1,
			k:      1,
			expect: [][]int{{1}},
		},
	}

	for _, test := range tests {
		result := combine(test.n, test.k)
		if !reflect.DeepEqual(result, test.expect) {
			t.Errorf("combine(%d, %d) = %v, expect %v", test.n, test.k, result, test.expect)
		}
	}
}