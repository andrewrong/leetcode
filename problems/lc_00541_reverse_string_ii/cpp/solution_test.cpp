#include "solution.h"
#include <gtest/gtest.h>
#include <string>

TEST(ReverseStringIITest, BasicTest) {
    Solution solution;
    std::string s = "abcdefg";
    int k = 2;
    std::string expected = "bacdfeg";
    std::string result = solution.reverseStr(s, k);
    EXPECT_EQ(result, expected);
}

TEST(ReverseStringIITest, SingleGroup) {
    Solution solution;
    std::string s = "abcdefgh";
    int k = 9; // k greater than string length
    std::string expected = "hgfedcba";
    std::string result = solution.reverseStr(s, k);
    EXPECT_EQ(result, expected);
}

TEST(ReverseStringIITest, ExactMultiple) {
    Solution solution;
    std::string s = "abcdefghijk";
    int k = 2;
    std::string expected = "bacdfeghjki";
    std::string result = solution.reverseStr(s, k);
    EXPECT_EQ(result, expected);
}

TEST(ReverseStringIITest, EmptyString) {
    Solution solution;
    std::string s = "";
    int k = 2;
    std::string expected = "";
    std::string result = solution.reverseStr(s, k);
    EXPECT_EQ(result, expected);
}