package solution

// 思路: 因为是字符串，用双指针的方式的从尾部进行scan,大概的step应该是这样的
// 1. 设置三个指针，可以被写入的pointer，非空的pointer tail，word start的pointer
// 2. 找到第一个非空的tail pointer，然后开始通过迭代的方式写入到当前的指针
func reverseWords(s string) string {
	if len(s) == 0 {
		return s
	}

	modifyPointer := len(s) - 1
	nonSpacePointer := -1
	movePointer := len(s) - 1

	for movePointer >= 0 {
		if s[movePointer] != ' ' && nonSpacePointer == -1 {

		}
	}
	return s
}
