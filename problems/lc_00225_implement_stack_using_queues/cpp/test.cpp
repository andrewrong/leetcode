// leetcode/problems/lc_00225_implement_stack_using_queues/cpp/test.cpp
#include "gtest/gtest.h"
#include "solution.h"

TEST(MyStack, ComprehensiveTest) {
  MyStack s;

  // Initially, the stack should be empty
  EXPECT_TRUE(s.empty());

  // Push 1
  s.push(1);
  EXPECT_FALSE(s.empty());
  EXPECT_EQ(s.top(), 1);

  // Push 2
  s.push(2);
  EXPECT_FALSE(s.empty());
  EXPECT_EQ(s.top(), 2); // Top should now be 2

  // Pop
  int val = s.pop();
  EXPECT_EQ(val, 2);
  EXPECT_FALSE(s.empty());
  EXPECT_EQ(s.top(), 1); // Top should be 1 again

  // Push 3
  s.push(3);
  EXPECT_EQ(s.top(), 3);

  // Pop remaining elements
  EXPECT_EQ(s.pop(), 3);
  EXPECT_EQ(s.top(), 1);
  EXPECT_EQ(s.pop(), 1);
  EXPECT_TRUE(s.empty());
}

TEST(MyStack, PushPopSequence) {
  MyStack s;
  s.push(1);
  s.push(2);
  EXPECT_EQ(s.pop(), 2);
  s.push(3);
  s.push(4);
  EXPECT_EQ(s.pop(), 4);
  EXPECT_EQ(s.pop(), 3);
  EXPECT_EQ(s.pop(), 1);
  EXPECT_TRUE(s.empty());
}
