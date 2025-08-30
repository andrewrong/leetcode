// leetcode/problems/lc_00239_sliding_window_maximum/go/solution.go
package main

// maxSlidingWindow finds the maximum value in each sliding window of size k.
// This optimized solution uses a slice as a monotonic decreasing deque.
func maxSlidingWindow(nums []int, k int) []int {
	// Handle edge cases where a valid window is not possible.
	if len(nums) < k || k <= 0 {
		return []int{}
	}

	// The deque stores indices of elements from `nums`.
	// It is "monotonic decreasing" because the values at these indices in `nums`
	// are strictly decreasing from front to back.
	// e.g., if deque is [i, j, l], then nums[i] > nums[j] > nums[l].
	deque := make([]int, 0, k)
	result := make([]int, 0, len(nums)-k+1)

	for i, num := range nums {
		// 1. Remove indices from the front of the deque that are no longer
		//    in the current sliding window [i-k+1, i].
		//    An index `j` is out of the window if j <= i-k.
		if len(deque) > 0 && deque[0] <= i-k {
			deque = deque[1:]
		}

		// 2. Maintain the monotonic decreasing property of the deque.
		//    Remove indices from the back of the deque if their corresponding
		//    values in `nums` are less than the current number.
		//    This is because a smaller number that appears earlier will never be
		//    the maximum in any future window that also includes the current, larger number.
		for len(deque) > 0 && nums[deque[len(deque)-1]] < num {
			deque = deque[:len(deque)-1]
		}

		// 3. Add the current index to the back of the deque.
		deque = append(deque, i)

		// 4. The front of the deque always holds the index of the maximum element
		//    for the current window. Once the window is full (i.e., has k elements),
		//    we can start adding the maximums to our result.
		if i >= k-1 {
			result = append(result, nums[deque[0]])
		}
	}

	return result
}