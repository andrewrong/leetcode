// leetcode/problems/lc_00020_valid_parentheses/cpp/test.cpp
#include "gtest/gtest.h"
#include "solution.h"

TEST(Solution, IsValid) {
  Solution s;

  // Test case 1: Basic valid case
  EXPECT_TRUE(s.isValid("()"));

  // Test case 2: All types valid
  EXPECT_TRUE(s.isValid("()[]{}"));

  // Test case 3: Basic invalid case
  EXPECT_FALSE(s.isValid("(]"));

  // Test case 4: Nested valid case
  EXPECT_TRUE(s.isValid("{[]}"));

  // Test case 5: Nested invalid case
  EXPECT_FALSE(s.isValid("([)]"));

  // Test case 6: Starts with closing bracket
  EXPECT_FALSE(s.isValid("]"));

  // Test case 7: Odd number of brackets
  EXPECT_FALSE(s.isValid("([)"));

  // Test case 8: Empty string is valid
  EXPECT_TRUE(s.isValid(""));

  // Test case 9: Long valid string
  EXPECT_TRUE(s.isValid("(([]){})"));

  // Test case 10: Long invalid string
  EXPECT_FALSE(s.isValid("(([]){})("));
}
