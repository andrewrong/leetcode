#include "binary_search.h"
#include <gtest/gtest.h>
#include <vector>

using namespace std;

TEST(BinarySearchTest, BinarySearch1Found) {
    BinarySearch bs;
    vector<int> nums = {1, 2, 3, 4, 5, 6, 7, 8, 9};
    int val = 5;
    int expected = 4;  // Index of 5 in the array
    EXPECT_EQ(bs.BinarySearch1(nums, val), expected);
}

TEST(BinarySearchTest, BinarySearch1NotFound) {
    BinarySearch bs;
    vector<int> nums = {1, 2, 3, 4, 5, 6, 7, 8, 9};
    int val = 10;  // Not in the array
    int expected = -1;
    EXPECT_EQ(bs.BinarySearch1(nums, val), expected);
}

TEST(BinarySearchTest, BinarySearch1EmptyArray) {
    BinarySearch bs;
    vector<int> nums = {};
    int val = 5;
    int expected = -1;
    EXPECT_EQ(bs.BinarySearch1(nums, val), expected);
}

TEST(BinarySearchTest, BinarySearch1SingleElementFound) {
    BinarySearch bs;
    vector<int> nums = {5};
    int val = 5;
    int expected = 0;
    EXPECT_EQ(bs.BinarySearch1(nums, val), expected);
}

TEST(BinarySearchTest, BinarySearch1SingleElementNotFound) {
    BinarySearch bs;
    vector<int> nums = {5};
    int val = 3;
    int expected = -1;
    EXPECT_EQ(bs.BinarySearch1(nums, val), expected);
}

TEST(BinarySearchTest, BinarySearch2Found) {
    BinarySearch bs;
    vector<int> nums = {1, 2, 3, 4, 5, 6, 7, 8, 9};
    int val = 5;
    int expected = 4;  // Index of 5 in the array
    EXPECT_EQ(bs.BinarySearch2(nums, val), expected);
}

TEST(BinarySearchTest, BinarySearch2NotFound) {
    BinarySearch bs;
    vector<int> nums = {1, 2, 3, 4, 5, 6, 7, 8, 9};
    int val = 10;  // Not in the array
    int expected = -1;
    EXPECT_EQ(bs.BinarySearch2(nums, val), expected);
}

TEST(BinarySearchTest, BinarySearch2EmptyArray) {
    BinarySearch bs;
    vector<int> nums = {};
    int val = 5;
    int expected = -1;
    EXPECT_EQ(bs.BinarySearch2(nums, val), expected);
}

TEST(BinarySearchTest, BinarySearch2SingleElementFound) {
    BinarySearch bs;
    vector<int> nums = {5};
    int val = 5;
    int expected = 0;
    EXPECT_EQ(bs.BinarySearch2(nums, val), expected);
}

TEST(BinarySearchTest, BinarySearch2SingleElementNotFound) {
    BinarySearch bs;
    vector<int> nums = {5};
    int val = 3;
    int expected = -1;
    EXPECT_EQ(bs.BinarySearch2(nums, val), expected);
}