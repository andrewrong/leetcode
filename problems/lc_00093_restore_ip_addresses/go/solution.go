package lc_00093_restore_ip_addresses

import (
	"strconv"
	"strings"
)

var steps []int = []int{
	1, 2, 3,
}

// restoreIpAddresses 复原IP地址
func restoreIpAddresses(s string) []string {
	all := make([][]string, 0)
	one := make([]string, 0)

	combination(s, &one, &all)

	last := make([]string, 0)
	for _, item := range all {
		last = append(last, strings.Join(item, "."))
	}
	return last
}

func combination(s string, one *[]string, all *[][]string) {
	if len(s) == 0 {
		if len(*one) == 4 {
			cp := make([]string, len(*one))
			copy(cp, *one)
			*all = append(*all, cp)
		}
		return
	}
	size := len(s)

	if size != 0 && len(*one) == 4 {
		return
	}

	for _, step := range steps {
		if step > size {
			break
		}

		part := s[:step]
		// 如果当前的切分是不合法的话，那不需要再迭代了，因为再变大也没意义
		if !isValidInt(part) {
			break
		}

		if (size - step) < (4 - len(*one) - 1) {
			break
		}

		*one = append(*one, part)
		combination(s[step:], one, all)
		*one = (*one)[:len(*one)-1]
	}
}

func isValidInt(s string) bool {
	if (s[0] - '0') == 0 {
		if len(s) > 1 {
			return false
		}

		return true
	}

	tmp, err := strconv.Atoi(s)
	if err != nil {
		return false
	}

	if tmp < 0 || tmp > 255 {
		return false
	}

	return true
}
