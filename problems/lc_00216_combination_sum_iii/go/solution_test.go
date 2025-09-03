package solution

import (
	"reflect"
	"testing"
)

func TestCombinationSum3(t *testing.T) {
	tests := []struct {
		k      int
		n      int
		expect [][]int
	}{
		{
			k:      3,
			n:      7,
			expect: [][]int{{1, 2, 4}},
		},
		{
			k:      3,
			n:      9,
			expect: [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}},
		},
		{
			k:      4,
			n:      1,
			expect: [][]int{},
		},
		{
			k:      3,
			n:      2,
			expect: [][]int{},
		},
		{
			k:      1,
			n:      1,
			expect: [][]int{{1}},
		},
	}

	for _, test := range tests {
		result := combinationSum3(test.k, test.n)
		if !reflect.DeepEqual(result, test.expect) {
			t.Errorf("combinationSum3(%d, %d) = %v, expect %v", test.k, test.n, result, test.expect)
		}
	}
}