#include "solution.h"
#include <gtest/gtest.h>

TEST(NQueensTest, BasicCases) {
    Solution solution;
    
    // Test case 1: n = 4
    auto result1 = solution.solveNQueens(4);
    std::vector<std::vector<std::string>> expected1 = {
        {".Q..", "...Q", "Q...", "..Q."},
        {"..Q.", "Q...", "...Q", ".Q.."}
    };
    EXPECT_EQ(result1, expected1);
    
    // Test case 2: n = 1
    auto result2 = solution.solveNQueens(1);
    std::vector<std::vector<std::string>> expected2 = {{"Q"}};
    EXPECT_EQ(result2, expected2);
    
    // Test case 3: n = 2 (no solution)
    auto result3 = solution.solveNQueens(2);
    std::vector<std::vector<std::string>> expected3 = {};
    EXPECT_EQ(result3, expected3);
    
    // Test case 4: n = 3 (no solution)
    auto result4 = solution.solveNQueens(3);
    std::vector<std::vector<std::string>> expected4 = {};
    EXPECT_EQ(result4, expected4);
}