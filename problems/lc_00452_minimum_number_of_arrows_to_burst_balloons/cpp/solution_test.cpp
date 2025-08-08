#include "solution.h"
#include <gtest/gtest.h>
#include <vector>

TEST(MinimumArrowsToBurstBalloonsTest, Example1) {
    Solution solution;
    std::vector<std::vector<int>> points = {{10,16},{2,8},{1,6},{7,12}};
    int expected = 2;
    EXPECT_EQ(solution.findMinArrowShots(points), expected);
}

TEST(MinimumArrowsToBurstBalloonsTest, Example2) {
    Solution solution;
    std::vector<std::vector<int>> points = {{1,2},{3,4},{5,6},{7,8}};
    int expected = 4;
    EXPECT_EQ(solution.findMinArrowShots(points), expected);
}

TEST(MinimumArrowsToBurstBalloonsTest, Example3) {
    Solution solution;
    std::vector<std::vector<int>> points = {{1,2},{2,3},{3,4},{4,5}};
    int expected = 2;
    EXPECT_EQ(solution.findMinArrowShots(points), expected);
}

TEST(MinimumArrowsToBurstBalloonsTest, EmptyArray) {
    Solution solution;
    std::vector<std::vector<int>> points = {};
    int expected = 0;
    EXPECT_EQ(solution.findMinArrowShots(points), expected);
}