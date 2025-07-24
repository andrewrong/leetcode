// problems/lc0027_remove_element/cpp/solution_test.cpp
#include "solution.h" // 包含解法文件
#include <gtest/gtest.h>
#include <vector>
#include <algorithm>

// 为了方便，可以直接在测试文件中使用命名空间
using namespace std;

// 自定义测试宏，用于验证结果
void ASSERT_REMOVE_ELEMENT(vector<int> nums, int val, int expected_k, vector<int> expected_nums) {
    Solution sol;
    int k = sol.Solve(nums, val);
    
    // 1. 检查返回的 k 值
    ASSERT_EQ(k, expected_k);
    
    // 2. 检查数组内容（顺序无关）
    // 截取前 k 个元素并排序
    vector<int> result_nums(nums.begin(), nums.begin() + k);
    sort(result_nums.begin(), result_nums.end());
    sort(expected_nums.begin(), expected_nums.end());
    
    ASSERT_EQ(result_nums, expected_nums);
}


TEST(RemoveElementTest, Example1) {
    ASSERT_REMOVE_ELEMENT({3, 2, 2, 3}, 3, 2, {2, 2});
}

TEST(RemoveElementTest, Example2) {
    ASSERT_REMOVE_ELEMENT({0, 1, 2, 2, 3, 0, 4, 2}, 2, 5, {0, 1, 4, 0, 3});
}

TEST(RemoveElementTest, NoElementToRemove) {
    ASSERT_REMOVE_ELEMENT({1, 2, 3, 4, 5}, 6, 5, {1, 2, 3, 4, 5});
}

TEST(RemoveElementTest, AllElementsToRemove) {
    ASSERT_REMOVE_ELEMENT({4, 4, 4, 4}, 4, 0, {});
}

TEST(RemoveElementTest, EmptyInputSlice) {
    ASSERT_REMOVE_ELEMENT({}, 1, 0, {});
}