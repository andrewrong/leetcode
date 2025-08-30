// leetcode/problems/lc_00347_top_k_frequent_elements/cpp/solution_test.cpp
#include "solution.cpp"
#include <gtest/gtest.h>
#include <algorithm>
#include <vector>

TEST(TopKFrequentTest, Example1) {
    Solution s;
    std::vector<int> nums = {1, 1, 1, 2, 2, 3};
    int k = 2;
    std::vector<int> expected = {1, 2};
    std::vector<int> result = s.topKFrequent(nums, k);
    std::sort(result.begin(), result.end());
    std::sort(expected.begin(), expected.end());
    ASSERT_EQ(expected, result);
}

TEST(TopKFrequentTest, Example2) {
    Solution s;
    std::vector<int> nums = {1};
    int k = 1;
    std::vector<int> expected = {1};
    std::vector<int> result = s.topKFrequent(nums, k);
    ASSERT_EQ(expected, result);
}

TEST(TopKFrequentTest, Example3) {
    Solution s;
    std::vector<int> nums = {4, 1, -1, 2, -1, 2, 3};
    int k = 2;
    std::vector<int> expected = {-1, 2};
    std::vector<int> result = s.topKFrequent(nums, k);
    std::sort(result.begin(), result.end());
    std::sort(expected.begin(), expected.end());
    ASSERT_EQ(expected, result);
}
