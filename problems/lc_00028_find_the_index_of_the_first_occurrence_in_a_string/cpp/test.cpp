// leetcode/problems/lc_00028_find_the_index_of_the_first_occurrence_in_a_string/test.cpp
#include "gtest/gtest.h"
#include "solution.h"

TEST(Solution, strStr) {
  Solution s;

  // Test case 1: Basic case, needle is in the middle
  EXPECT_EQ(s.strStr("sadbutsad", "sad"), 0);

  // Test case 2: Needle is not in haystack
  EXPECT_EQ(s.strStr("leetcode", "leeto"), -1);

  // Test case 3: Needle is at the beginning
  EXPECT_EQ(s.strStr("hello", "he"), 0);

  // Test case 4: Needle is at the end
  EXPECT_EQ(s.strStr("hello", "lo"), 3);

  // Test case 5: Haystack and needle are the same
  EXPECT_EQ(s.strStr("a", "a"), 0);

  // Test case 6: Needle is an empty string
  EXPECT_EQ(s.strStr("abc", ""), 0);

  // Test case 7: Haystack is an empty string and needle is not
  EXPECT_EQ(s.strStr("", "a"), -1);

  // Test case 8: Both are empty strings
  EXPECT_EQ(s.strStr("", ""), 0);

  // Test case 9: Needle is longer than haystack
  EXPECT_EQ(s.strStr("a", "aa"), -1);

  // Test case 10: Complex case with repeated characters
  EXPECT_EQ(s.strStr("mississippi", "issip"), 4);

  // Test case 11: Another complex case
  EXPECT_EQ(s.strStr("ababcabcabababd", "ababd"), 10);
}
