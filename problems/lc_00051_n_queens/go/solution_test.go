package lc_00051_n_queens

import (
	"reflect"
	"testing"
)

func Test_solveNQueens(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "4 queens",
			args: args{n: 4},
			want: [][]string{
				{".Q..", "...Q", "Q...", "..Q."},
				{"..Q.", "Q...", "...Q", ".Q.."},
			},
		},
		{
			name: "1 queen",
			args: args{n: 1},
			want: [][]string{{"Q"}},
		},
		{
			name: "2 queens - no solution",
			args: args{n: 2},
			want: [][]string{},
		},
		{
			name: "3 queens - no solution",
			args: args{n: 3},
			want: [][]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := solveNQueens(tt.args.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solveNQueens() = %v, want %v", got, tt.want)
			}
		})
	}
}