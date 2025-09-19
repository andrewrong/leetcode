#include <vector>
#include <algorithm>
#include <gtest/gtest.h>

using namespace std;

// 函数声明
vector<vector<int>> combinationSum2(vector<int>& candidates, int target);

TEST(CombinationSum2Test, TestCases) {
    // 测试用例1: 示例1
    vector<int> candidates1 = {10, 1, 2, 7, 6, 1, 5};
    vector<vector<int>> expected1 = {{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}};
    auto result1 = combinationSum2(candidates1, 8);
    EXPECT_EQ(result1, expected1);
    
    // 测试用例2: 示例2
    vector<int> candidates2 = {2, 5, 2, 1, 2};
    vector<vector<int>> expected2 = {{1, 2, 2}, {5}};
    auto result2 = combinationSum2(candidates2, 5);
    EXPECT_EQ(result2, expected2);
    
    // 测试用例3: 无解情况
    vector<int> candidates3 = {1, 2, 3};
    vector<vector<int>> expected3 = {};
    auto result3 = combinationSum2(candidates3, 7);
    EXPECT_EQ(result3, expected3);
    
    // 测试用例4: 空数组
    vector<int> candidates4 = {};
    vector<vector<int>> expected4 = {};
    auto result4 = combinationSum2(candidates4, 3);
    EXPECT_EQ(result4, expected4);
    
    // 测试用例5: 单元素数组
    vector<int> candidates5 = {1};
    vector<vector<int>> expected5 = {{1}};
    auto result5 = combinationSum2(candidates5, 1);
    EXPECT_EQ(result5, expected5);
    
    // 测试用例6: 重复元素
    vector<int> candidates6 = {1, 1, 1, 1};
    vector<vector<int>> expected6 = {{1, 1, 1}};
    auto result6 = combinationSum2(candidates6, 3);
    EXPECT_EQ(result6, expected6);
    
    // 测试用例7: 目标值为0
    vector<int> candidates7 = {1, 2, 3};
    vector<vector<int>> expected7 = {};
    auto result7 = combinationSum2(candidates7, 0);
    EXPECT_EQ(result7, expected7);
    
    // 测试用例8: 包含重复元素
    vector<int> candidates8 = {1, 1, 2, 2, 3, 3};
    vector<vector<int>> expected8 = {{1, 1, 2, 2}, {1, 2, 3}, {3, 3}};
    auto result8 = combinationSum2(candidates8, 6);
    EXPECT_EQ(result8, expected8);
    
    // 测试用例9: 大量重复元素
    vector<int> candidates9 = {1, 1, 1, 1, 1, 1, 1, 1, 1, 1};
    vector<vector<int>> expected9 = {{1, 1, 1}};
    auto result9 = combinationSum2(candidates9, 3);
    EXPECT_EQ(result9, expected9);
    
    // 测试用例10: 目标值为1
    vector<int> candidates10 = {1, 2, 3};
    vector<vector<int>> expected10 = {{1}};
    auto result10 = combinationSum2(candidates10, 1);
    EXPECT_EQ(result10, expected10);
}