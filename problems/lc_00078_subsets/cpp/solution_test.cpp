#include "solution.h"
#include <gtest/gtest.h>
#include <algorithm>

TEST(SubsetsTest, Example1) {
    Solution solution;
    std::vector<int> nums = {1, 2, 3};
    std::vector<std::vector<int>> result = solution.subsets(nums);
    
    // Expected subsets
    std::vector<std::vector<int>> expected = {
        {}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}
    };
    
    // Sort both result and expected for comparison
    for (auto& subset : result) {
        std::sort(subset.begin(), subset.end());
    }
    for (auto& subset : expected) {
        std::sort(subset.begin(), subset.end());
    }
    
    std::sort(result.begin(), result.end());
    std::sort(expected.begin(), expected.end());
    
    EXPECT_EQ(result, expected);
}

TEST(SubsetsTest, SingleElement) {
    Solution solution;
    std::vector<int> nums = {0};
    std::vector<std::vector<int>> result = solution.subsets(nums);
    
    std::vector<std::vector<int>> expected = {{}, {0}};
    
    // Sort both result and expected for comparison
    for (auto& subset : result) {
        std::sort(subset.begin(), subset.end());
    }
    for (auto& subset : expected) {
        std::sort(subset.begin(), subset.end());
    }
    
    std::sort(result.begin(), result.end());
    std::sort(expected.begin(), expected.end());
    
    EXPECT_EQ(result, expected);
}

TEST(SubsetsTest, EmptyArray) {
    Solution solution;
    std::vector<int> nums = {};
    std::vector<std::vector<int>> result = solution.subsets(nums);
    
    std::vector<std::vector<int>> expected = {{}};
    
    EXPECT_EQ(result, expected);
}