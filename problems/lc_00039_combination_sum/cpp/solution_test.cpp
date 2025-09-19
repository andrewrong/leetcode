#include <vector>
#include <algorithm>
#include <gtest/gtest.h>

using namespace std;

// 函数声明
vector<vector<int>> combinationSum(vector<int>& candidates, int target);

TEST(CombinationSumTest, TestCases) {
    // 测试用例1: 示例1
    vector<int> candidates1 = {2, 3, 6, 7};
    vector<vector<int>> expected1 = {{2, 2, 3}, {7}};
    auto result1 = combinationSum(candidates1, 7);
    EXPECT_EQ(result1, expected1);
    
    // 测试用例2: 示例2
    vector<int> candidates2 = {2, 3, 5};
    vector<vector<int>> expected2 = {{2, 2, 2, 2}, {2, 3, 3}, {3, 5}};
    auto result2 = combinationSum(candidates2, 8);
    EXPECT_EQ(result2, expected2);
    
    // 测试用例3: 无解情况
    vector<int> candidates3 = {2};
    vector<vector<int>> expected3 = {};
    auto result3 = combinationSum(candidates3, 1);
    EXPECT_EQ(result3, expected3);
    
    // 测试用例4: 空数组
    vector<int> candidates4 = {};
    vector<vector<int>> expected4 = {};
    auto result4 = combinationSum(candidates4, 3);
    EXPECT_EQ(result4, expected4);
    
    // 测试用例5: 包含0的情况
    vector<int> candidates5 = {0, 1};
    vector<vector<int>> expected5 = {{0, 1}};
    auto result5 = combinationSum(candidates5, 1);
    EXPECT_EQ(result5, expected5);
    
    // 测试用例6: 目标值为0
    vector<int> candidates6 = {1, 2, 3};
    vector<vector<int>> expected6 = {};
    auto result6 = combinationSum(candidates6, 0);
    EXPECT_EQ(result6, expected6);
    
    // 测试用例7: 完全匹配
    vector<int> candidates7 = {1, 2, 3};
    vector<vector<int>> expected7 = {{1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 2}, {1, 1, 2, 2}, {2, 2, 2}, {1, 1, 1, 3}, {1, 2, 3}, {3, 3}};
    auto result7 = combinationSum(candidates7, 6);
    EXPECT_EQ(result7, expected7);
    
    // 测试用例8: 重复元素
    vector<int> candidates8 = {2, 2, 3, 7};
    vector<vector<int>> expected8 = {{2, 2, 3}, {7}};
    auto result8 = combinationSum(candidates8, 7);
    EXPECT_EQ(result8, expected8);
}