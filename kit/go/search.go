package kit

import "log"

type Search interface {
	BinarySearch1(nums []int, target int) int
	BinarySearch2(nums []int, target int) int
	StringSearch(a, substring string) int
	GetNext(str string) []int
}

type search struct{}

func NewSearch() Search {
	return &search{}
}

/**
 * 二分搜索的本质在于区间的开闭的控制，有两种[)和[]这两种，主要区别在于边界问题的处理
 * 如果找到就返回index，如果找不到就返回-1,
 * 下面binarysearch1是左闭右开
 */
func (s *search) BinarySearch1(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left := 0
	right := len(nums)

	for left < right {
		mid := (left + right) / 2

		if nums[mid] > target {
			right = mid
			continue
		} else if nums[mid] < target {
			left = mid + 1
			continue
		} else {
			return mid
		}
	}

	return -1
}

// 这次的区间是[]这种;
func (s *search) BinarySearch2(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}

	return -1
}

func (s *search) StringSearch(a, substring string) int {
	if len(a) < len(substring) {
		return -1
	}

	// step 1: 获得substring的next函数,所谓next其实是找到目前substring的最大的公共的前后substring
	next := s.GetNext(substring)
	log.Printf("substring:%v,next: %v", substring, next)

	i := 0
	j := 0

	for i < len(a) && j < len(substring) {
		if j > 0 && substring[j] != a[i] {
			j = next[j-1]
		}

		if substring[j] == a[i] {
			i++
			j++
			continue
		}

		if j == 0 {
			i++
		}
	}

	if j == len(substring) && i <= len(a) {
		return i - len(substring)
	}
	return -1
}

func (s *search) GetNext(str string) []int {
	next := make([]int, len(str))

	for i := 1; i < len(str); i += 1 {
		// 前一个字符串的LPS
		j := next[i-1]
		for j > 0 && str[j] != str[i] {
			j = next[j-1]
		}

		if str[j] == str[i] {
			next[i] = j + 1
		} else {
			next[i] = 0
		}
	}

	return next
}
