#include "solution.h"
#include <gtest/gtest.h>
#include <algorithm>

// Helper function to sort vectors for comparison
void sortResult(std::vector<std::vector<int>>& result) {
    for (auto& subset : result) {
        std::sort(subset.begin(), subset.end());
    }
    std::sort(result.begin(), result.end());
}

// Helper function to compare two vectors of vectors
bool compareResults(std::vector<std::vector<int>> a, std::vector<std::vector<int>> b) {
    sortResult(a);
    sortResult(b);
    return a == b;
}

TEST(SubsetsWithDupTest, Example1) {
    Solution solution;
    std::vector<int> nums = {1, 2, 2};
    std::vector<std::vector<int>> expected = {{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}};
    std::vector<std::vector<int>> result = solution.subsetsWithDup(nums);
    
    EXPECT_TRUE(compareResults(result, expected));
}

TEST(SubsetsWithDupTest, Example2) {
    Solution solution;
    std::vector<int> nums = {0};
    std::vector<std::vector<int>> expected = {{}, {0}};
    std::vector<std::vector<int>> result = solution.subsetsWithDup(nums);
    
    EXPECT_TRUE(compareResults(result, expected));
}

TEST(SubsetsWithDupTest, EmptyArray) {
    Solution solution;
    std::vector<int> nums = {};
    std::vector<std::vector<int>> expected = {{}};
    std::vector<std::vector<int>> result = solution.subsetsWithDup(nums);
    
    EXPECT_TRUE(compareResults(result, expected));
}

TEST(SubsetsWithDupTest, AllDuplicates) {
    Solution solution;
    std::vector<int> nums = {1, 1, 1};
    std::vector<std::vector<int>> expected = {{}, {1}, {1, 1}, {1, 1, 1}};
    std::vector<std::vector<int>> result = solution.subsetsWithDup(nums);
    
    EXPECT_TRUE(compareResults(result, expected));
}

TEST(SubsetsWithDupTest, NoDuplicates) {
    Solution solution;
    std::vector<int> nums = {1, 2, 3};
    std::vector<std::vector<int>> expected = {{}, {1}, {1, 2}, {1, 2, 3}, {1, 3}, {2}, {2, 3}, {3}};
    std::vector<std::vector<int>> result = solution.subsetsWithDup(nums);
    
    EXPECT_TRUE(compareResults(result, expected));
}

TEST(SubsetsWithDupTest, NegativeNumbersWithDuplicates) {
    Solution solution;
    std::vector<int> nums = {-1, 1, 2, 2};
    std::vector<std::vector<int>> expected = {
        {}, {-1}, {-1, 1}, {-1, 1, 2}, {-1, 1, 2, 2}, 
        {-1, 2}, {-1, 2, 2}, {1}, {1, 2}, {1, 2, 2}, 
        {2}, {2, 2}
    };
    std::vector<std::vector<int>> result = solution.subsetsWithDup(nums);
    
    EXPECT_TRUE(compareResults(result, expected));
}