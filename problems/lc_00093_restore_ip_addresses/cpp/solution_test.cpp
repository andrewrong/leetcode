#include "solution.h"
#include <vector>
#include <string>
#include <algorithm>
#include <gtest/gtest.h>

class RestoreIpAddressesTest : public ::testing::Test {};

TEST_F(RestoreIpAddressesTest, Example1) {
    Solution solution;
    std::string s = "25525511135";
    std::vector<std::string> expected = {"255.255.11.135", "255.255.111.35"};
    std::vector<std::string> result = solution.restoreIpAddresses(s);
    
    // 排序后比较，因为顺序可能不同
    std::sort(result.begin(), result.end());
    std::sort(expected.begin(), expected.end());
    
    EXPECT_EQ(result, expected);
}

TEST_F(RestoreIpAddressesTest, Example2) {
    Solution solution;
    std::string s = "0000";
    std::vector<std::string> expected = {"0.0.0.0"};
    std::vector<std::string> result = solution.restoreIpAddresses(s);
    
    EXPECT_EQ(result, expected);
}

TEST_F(RestoreIpAddressesTest, Example3) {
    Solution solution;
    std::string s = "101023";
    std::vector<std::string> expected = {"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"};
    std::vector<std::string> result = solution.restoreIpAddresses(s);
    
    // 排序后比较，因为顺序可能不同
    std::sort(result.begin(), result.end());
    std::sort(expected.begin(), expected.end());
    
    EXPECT_EQ(result, expected);
}

TEST_F(RestoreIpAddressesTest, EmptyString) {
    Solution solution;
    std::string s = "";
    std::vector<std::string> expected = {};
    std::vector<std::string> result = solution.restoreIpAddresses(s);
    
    EXPECT_EQ(result, expected);
}

TEST_F(RestoreIpAddressesTest, TooShort) {
    Solution solution;
    std::string s = "111";
    std::vector<std::string> expected = {};
    std::vector<std::string> result = solution.restoreIpAddresses(s);
    
    EXPECT_EQ(result, expected);
}

TEST_F(RestoreIpAddressesTest, TooLong) {
    Solution solution;
    std::string s = "111111111111";
    std::vector<std::string> expected = {};
    std::vector<std::string> result = solution.restoreIpAddresses(s);
    
    EXPECT_EQ(result, expected);
}

TEST_F(RestoreIpAddressesTest, SpecialCase) {
    Solution solution;
    std::string s = "1111";
    std::vector<std::string> expected = {"1.1.1.1"};
    std::vector<std::string> result = solution.restoreIpAddresses(s);
    
    EXPECT_EQ(result, expected);
}

TEST_F(RestoreIpAddressesTest, Contains255) {
    Solution solution;
    std::string s = "111255";
    std::vector<std::string> expected = {"1.1.1.255", "11.1.1.255", "111.1.1.255"};
    std::vector<std::string> result = solution.restoreIpAddresses(s);
    
    // 排序后比较，因为顺序可能不同
    std::sort(result.begin(), result.end());
    std::sort(expected.begin(), expected.end());
    
    EXPECT_EQ(result, expected);
}