package kit

type Search interface {
	BinarySearch1(nums []int, target int) int
	BinarySearch2(nums []int, target int) int
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
