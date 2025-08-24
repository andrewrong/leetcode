// leetcode/problems/lc_00459_repeated_substring_pattern/cpp/test.cpp
#include "gtest/gtest.h"
#include "solution.h"

TEST(Solution, RepeatedSubstringPattern) {
  Solution s;

  // Test case 1: Basic true case
  EXPECT_TRUE(s.repeatedSubstringPattern("abab"));

  // Test case 2: Basic false case
  EXPECT_FALSE(s.repeatedSubstringPattern("aba"));

  // Test case 3: Longer true case
  EXPECT_TRUE(s.repeatedSubstringPattern("abcabcabcabc"));

  // Test case 4: Single character repeated
  EXPECT_TRUE(s.repeatedSubstringPattern("aaaa"));

  // Test case 5: Single character not repeated
  EXPECT_FALSE(s.repeatedSubstringPattern("a"));

  // Test case 6: Complex false case
  EXPECT_FALSE(s.repeatedSubstringPattern("abac"));

  // Test case 7: True case with longer substring
  EXPECT_TRUE(s.repeatedSubstringPattern("ababab"));

  // Test case 8: False case with similar pattern
  EXPECT_FALSE(s.repeatedSubstringPattern("ababac"));
}
