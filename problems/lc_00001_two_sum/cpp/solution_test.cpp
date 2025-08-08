#include "solution.h"
#include <gtest/gtest.h>
#include <vector>

TEST(TwoSumTest, Example1) {
    Solution solution;
    std::vector<int> nums = {2, 7, 11, 15};
    int target = 9;
    std::vector<int> expected = {0, 1};
    EXPECT_EQ(solution.twoSum(nums, target), expected);
}

TEST(TwoSumTest, Example2) {
    Solution solution;
    std::vector<int> nums = {3, 2, 4};
    int target = 6;
    std::vector<int> expected = {1, 2};
    EXPECT_EQ(solution.twoSum(nums, target), expected);
}

TEST(TwoSumTest, Example3) {
    Solution solution;
    std::vector<int> nums = {3, 3};
    int target = 6;
    std::vector<int> expected = {0, 1};
    EXPECT_EQ(solution.twoSum(nums, target), expected);
}

TEST(TwoSumTest, EmptyArray) {
    Solution solution;
    std::vector<int> nums = {};
    int target = 0;
    std::vector<int> expected = {};
    EXPECT_EQ(solution.twoSum(nums, target), expected);
}