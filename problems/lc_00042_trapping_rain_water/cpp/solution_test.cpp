// problems/lc_00042_trapping_rain_water/cpp/solution_test.cpp
#include "solution.h" // 包含解法文件
#include <gtest/gtest.h>
#include <vector>

// 使用 TEST 宏定义一个测试用例
TEST(TrappingRainWaterTest, Example1) {
    Solution sol;
    std::vector<int> height = {0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1};
    int expected = 6;
    int actual = sol.trap(height);
    ASSERT_EQ(expected, actual);
}

TEST(TrappingRainWaterTest, Example2) {
    Solution sol;
    std::vector<int> height = {4, 2, 0, 3, 2, 5};
    int expected = 9;
    int actual = sol.trap(height);
    ASSERT_EQ(expected, actual);
}
