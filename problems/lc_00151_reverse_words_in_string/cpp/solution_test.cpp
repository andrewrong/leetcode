#include "solution.h"
#include <gtest/gtest.h>
#include <string>

TEST(ReverseWordsInStringTest, BasicTest) {
    Solution solution;
    std::string s = "the sky is blue";
    std::string expected = "blue is sky the";
    std::string result = solution.reverseWords(s);
    EXPECT_EQ(result, expected);
}

TEST(ReverseWordsInStringTest, ExtraSpaces) {
    Solution solution;
    std::string s = "  hello world  ";
    std::string expected = "world hello";
    std::string result = solution.reverseWords(s);
    EXPECT_EQ(result, expected);
}

TEST(ReverseWordsInStringTest, MultipleSpaces) {
    Solution solution;
    std::string s = "a good   example";
    std::string expected = "example good a";
    std::string result = solution.reverseWords(s);
    EXPECT_EQ(result, expected);
}

TEST(ReverseWordsInStringTest, SingleWord) {
    Solution solution;
    std::string s = "word";
    std::string expected = "word";
    std::string result = solution.reverseWords(s);
    EXPECT_EQ(result, expected);
}

TEST(ReverseWordsInStringTest, EmptyString) {
    Solution solution;
    std::string s = "";
    std::string expected = "";
    std::string result = solution.reverseWords(s);
    EXPECT_EQ(result, expected);
}