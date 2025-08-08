#include "solution.h"
#include <gtest/gtest.h>
#include <vector>

TEST(IntersectionOfTwoArraysTest, Example1) {
    Solution solution;
    std::vector<int> nums1 = {1, 2, 2, 1};
    std::vector<int> nums2 = {2, 2};
    std::vector<int> expected = {2};
    EXPECT_EQ(solution.intersection(nums1, nums2), expected);
}

TEST(IntersectionOfTwoArraysTest, Example2) {
    Solution solution;
    std::vector<int> nums1 = {4, 9, 5};
    std::vector<int> nums2 = {9, 4, 9, 8, 4};
    std::vector<int> expected = {9, 4}; // Order doesn't matter
    auto result = solution.intersection(nums1, nums2);
    
    // Since order doesn't matter, we need to check if both elements are present
    EXPECT_EQ(result.size(), expected.size());
    for (int val : expected) {
        bool found = false;
        for (int res : result) {
            if (res == val) {
                found = true;
                break;
            }
        }
        EXPECT_TRUE(found);
    }
}

TEST(IntersectionOfTwoArraysTest, EmptyArrays) {
    Solution solution;
    std::vector<int> nums1 = {};
    std::vector<int> nums2 = {};
    std::vector<int> expected = {};
    EXPECT_EQ(solution.intersection(nums1, nums2), expected);
}

TEST(IntersectionOfTwoArraysTest, NoIntersection) {
    Solution solution;
    std::vector<int> nums1 = {1, 2, 3};
    std::vector<int> nums2 = {4, 5, 6};
    std::vector<int> expected = {};
    EXPECT_EQ(solution.intersection(nums1, nums2), expected);
}