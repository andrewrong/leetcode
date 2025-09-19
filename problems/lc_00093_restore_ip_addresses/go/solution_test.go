package lc_00093_restore_ip_addresses

import (
	"reflect"
	"testing"
)

func Test_restoreIpAddresses(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []string
	}{
		{
			name: "示例1",
			s:    "25525511135",
			want: []string{"255.255.11.135", "255.255.111.35"},
		},
		{
			name: "示例2",
			s:    "0000",
			want: []string{"0.0.0.0"},
		},
		{
			name: "示例3",
			s:    "101023",
			want: []string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"},
		},
		{
			name: "边界情况 - 空字符串",
			s:    "",
			want: []string{},
		},
		{
			name: "边界情况 - 长度不足",
			s:    "111",
			want: []string{},
		},
		{
			name: "边界情况 - 长度超出",
			s:    "111111111111",
			want: []string{},
		},
		{
			name: "边界情况 - 全是0但长度超过4",
			s:    "00000",
			want: []string{},
		},
		{
			name: "特殊IP地址",
			s:    "1111",
			want: []string{"1.1.1.1"},
		},
		{
			name: "包含255",
			s:    "111255",
			want: []string{"1.1.1.255", "11.1.1.255", "111.1.1.255"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := restoreIpAddresses(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restoreIpAddresses() = %v, want %v", got, tt.want)
			}
		})
	}
}
