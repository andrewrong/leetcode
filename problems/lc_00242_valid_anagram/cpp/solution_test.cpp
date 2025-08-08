#include "solution.h"
#include <gtest/gtest.h>

TEST(ValidAnagramTest, Example1) {
    Solution solution;
    std::string s = "anagram";
    std::string t = "nagaram";
    EXPECT_TRUE(solution.isAnagram(s, t));
}

TEST(ValidAnagramTest, Example2) {
    Solution solution;
    std::string s = "rat";
    std::string t = "car";
    EXPECT_FALSE(solution.isAnagram(s, t));
}

TEST(ValidAnagramTest, EmptyStrings) {
    Solution solution;
    std::string s = "";
    std::string t = "";
    EXPECT_TRUE(solution.isAnagram(s, t));
}

TEST(ValidAnagramTest, DifferentLengths) {
    Solution solution;
    std::string s = "abc";
    std::string t = "abcd";
    EXPECT_FALSE(solution.isAnagram(s, t));
}