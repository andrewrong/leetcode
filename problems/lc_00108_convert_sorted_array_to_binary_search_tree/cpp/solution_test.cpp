#include "solution.h"
#include <gtest/gtest.h>
#include <vector>

TEST(ConvertSortedArrayToBSTTest, Example1) {
    Solution solution;
    std::vector<int> nums = {-10, -3, 0, 5, 9};
    TreeNode* result = solution.sortedArrayToBST(nums);
    // Add assertions when implementation is available
}

TEST(ConvertSortedArrayToBSTTest, Example2) {
    Solution solution;
    std::vector<int> nums = {1, 3};
    TreeNode* result = solution.sortedArrayToBST(nums);
    // Add assertions when implementation is available
}

TEST(ConvertSortedArrayToBSTTest, SingleElement) {
    Solution solution;
    std::vector<int> nums = {1};
    TreeNode* result = solution.sortedArrayToBST(nums);
    // Add assertions when implementation is available
}

TEST(ConvertSortedArrayToBSTTest, EmptyArray) {
    Solution solution;
    std::vector<int> nums = {};
    TreeNode* result = solution.sortedArrayToBST(nums);
    // Add assertions when implementation is available
}