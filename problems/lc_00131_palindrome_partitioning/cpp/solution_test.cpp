#include "solution.h"
#include <gtest/gtest.h>
#include <vector>
#include <string>

TEST(PalindromePartitioningTest, Example1) {
    Solution solution;
    std::string s = "aab";
    std::vector<std::vector<std::string>> expected = {{"a", "a", "b"}, {"aa", "b"}};
    EXPECT_EQ(solution.partition(s), expected);
}

TEST(PalindromePartitioningTest, SingleChar) {
    Solution solution;
    std::string s = "a";
    std::vector<std::vector<std::string>> expected = {{"a"}};
    EXPECT_EQ(solution.partition(s), expected);
}

TEST(PalindromePartitioningTest, AllPalindromes) {
    Solution solution;
    std::string s = "aba";
    std::vector<std::vector<std::string>> expected = {{"a", "b", "a"}, {"aba"}};
    EXPECT_EQ(solution.partition(s), expected);
}

TEST(PalindromePartitioningTest, ComplexCase) {
    Solution solution;
    std::string s = "abcba";
    std::vector<std::vector<std::string>> expected = {
        {"a", "b", "c", "b", "a"},
        {"a", "bcb", "a"},
        {"abcba"}
    };
    EXPECT_EQ(solution.partition(s), expected);
}

TEST(PalindromePartitioningTest, EmptyString) {
    Solution solution;
    std::string s = "";
    std::vector<std::vector<std::string>> expected = {};
    EXPECT_EQ(solution.partition(s), expected);
}