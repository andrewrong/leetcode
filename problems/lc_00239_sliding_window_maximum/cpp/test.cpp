// leetcode/problems/lc_00239_sliding_window_maximum/cpp/test.cpp
#include "gtest/gtest.h"
#include "solution.h"

TEST(Solution, MaxSlidingWindow) {
  Solution s;

  // Example 1
  std::vector<int> nums1 = {1, 3, -1, -3, 5, 3, 6, 7};
  int k1 = 3;
  std::vector<int> expected1 = {3, 3, 5, 5, 6, 7};
  EXPECT_EQ(s.maxSlidingWindow(nums1, k1), expected1);

  // Test case with k = 1
  std::vector<int> nums2 = {1, 2, 3, 4, 5};
  int k2 = 1;
  std::vector<int> expected2 = {1, 2, 3, 4, 5};
  EXPECT_EQ(s.maxSlidingWindow(nums2, k2), expected2);

  // Test case with k = size of nums
  std::vector<int> nums3 = {5, 4, 3, 2, 1};
  int k3 = 5;
  std::vector<int> expected3 = {5};
  EXPECT_EQ(s.maxSlidingWindow(nums3, k3), expected3);

  // Test case with decreasing numbers
  std::vector<int> nums4 = {9, 8, 7, 6, 5};
  int k4 = 2;
  std::vector<int> expected4 = {9, 8, 7, 6};
  EXPECT_EQ(s.maxSlidingWindow(nums4, k4), expected4);

  // Test case with all same numbers
  std::vector<int> nums5 = {4, 4, 4, 4, 4};
  int k5 = 4;
  std::vector<int> expected5 = {4, 4};
  EXPECT_EQ(s.maxSlidingWindow(nums5, k5), expected5);
}
