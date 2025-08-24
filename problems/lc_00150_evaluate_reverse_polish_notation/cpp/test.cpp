// leetcode/problems/lc_00150_evaluate_reverse_polish_notation/cpp/test.cpp
#include "gtest/gtest.h"
#include "solution.h"

TEST(Solution, EvalRPN) {
  Solution s;

  // Example 1
  std::vector<std::string> tokens1 = {"2", "1", "+", "3", "*"};
  EXPECT_EQ(s.evalRPN(tokens1), 9);

  // Example 2
  std::vector<std::string> tokens2 = {"4", "13", "5", "/", "+"};
  EXPECT_EQ(s.evalRPN(tokens2), 6);

  // Example 3
  std::vector<std::string> tokens3 = {"10", "6", "9", "3", "+", "-11", "*",
                                      "/",  "*", "17", "+", "5",   "+"};
  EXPECT_EQ(s.evalRPN(tokens3), 22);

  // Test with negative numbers
  std::vector<std::string> tokens4 = {"-1", "1", "*", "-1", "+"};
  EXPECT_EQ(s.evalRPN(tokens4), -2);

  // Test with subtraction
  std::vector<std::string> tokens5 = {"5", "2", "-"};
  EXPECT_EQ(s.evalRPN(tokens5), 3);

  // Test with single number
  std::vector<std::string> tokens6 = {"42"};
  EXPECT_EQ(s.evalRPN(tokens6), 42);
}
