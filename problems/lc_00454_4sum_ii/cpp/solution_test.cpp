#include "solution.h"
#include <gtest/gtest.h>
#include <vector>

TEST(FourSumIITest, Example1) {
    Solution solution;
    std::vector<int> nums1 = {1, 2};
    std::vector<int> nums2 = {-2, -1};
    std::vector<int> nums3 = {-1, 2};
    std::vector<int> nums4 = {0, 2};
    int expected = 2;
    EXPECT_EQ(solution.fourSumCount(nums1, nums2, nums3, nums4), expected);
}

TEST(FourSumIITest, Example2) {
    Solution solution;
    std::vector<int> nums1 = {0};
    std::vector<int> nums2 = {0};
    std::vector<int> nums3 = {0};
    std::vector<int> nums4 = {0};
    int expected = 1;
    EXPECT_EQ(solution.fourSumCount(nums1, nums2, nums3, nums4), expected);
}

TEST(FourSumIITest, EmptyArrays) {
    Solution solution;
    std::vector<int> nums1 = {};
    std::vector<int> nums2 = {};
    std::vector<int> nums3 = {};
    std::vector<int> nums4 = {};
    int expected = 0;
    EXPECT_EQ(solution.fourSumCount(nums1, nums2, nums3, nums4), expected);
}