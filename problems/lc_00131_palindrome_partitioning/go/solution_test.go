package solution

import (
	"reflect"
	"testing"
)

func TestPartition(t *testing.T) {
	tests := []struct {
		s        string
		expected [][]string
	}{
		{
			"aab",
			[][]string{{"a", "a", "b"}, {"aa", "b"}},
		},
		{
			"a",
			[][]string{{"a"}},
		},
		{
			"aba",
			[][]string{{"a", "b", "a"}, {"aba"}},
		},
		{
			"abcba",
			[][]string{
				{"a", "b", "c", "b", "a"},
				{"a", "bcb", "a"},
				{"abcba"},
			},
		},
		{
			"",
			[][]string{},
		},
	}

	for _, test := range tests {
		result := Partition(test.s)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Partition(%q) = %v; expected %v", test.s, result, test.expected)
		}
	}
}