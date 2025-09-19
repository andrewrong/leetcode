package solution

import (
	"fmt"
)

var digitsMap map[byte]string = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	result := []string{}
	if len(digits) == 0 {
		return result
	}

	size := len(digits)
	for _, v := range digitsMap[digits[0]] {
		if size == 1 {
			result = append(result, string(v))
		} else {
			tmp := letterCombinations(digits[1:])
			for _, res := range tmp {
				result = append(result, fmt.Sprintf("%s%s", string(v), res))
			}
		}
	}
	return result
}
