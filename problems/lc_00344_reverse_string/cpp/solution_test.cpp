#include "solution.h"
#include <gtest/gtest.h>
#include <vector>
#include <string>

TEST(ReverseStringTest, BasicTest) {
    Solution solution;
    std::vector<char> s = {'h', 'e', 'l', 'l', 'o'};
    std::vector<char> expected = {'o', 'l', 'l', 'e', 'h'};
    solution.reverseString(s);
    EXPECT_EQ(s, expected);
}

TEST(ReverseStringTest, SingleChar) {
    Solution solution;
    std::vector<char> s = {'H'};
    std::vector<char> expected = {'H'};
    solution.reverseString(s);
    EXPECT_EQ(s, expected);
}

TEST(ReverseStringTest, EmptyString) {
    Solution solution;
    std::vector<char> s = {};
    std::vector<char> expected = {};
    solution.reverseString(s);
    EXPECT_EQ(s, expected);
}