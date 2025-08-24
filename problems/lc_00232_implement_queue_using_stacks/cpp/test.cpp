// leetcode/problems/lc_00232_implement_queue_using_stacks/cpp/test.cpp
#include "gtest/gtest.h"
#include "solution.h"

TEST(MyQueue, ComprehensiveTest) {
  MyQueue q;

  // Initially, the queue should be empty
  EXPECT_TRUE(q.empty());

  // Push 1
  q.push(1);
  EXPECT_FALSE(q.empty());
  EXPECT_EQ(q.peek(), 1);

  // Push 2
  q.push(2);
  EXPECT_FALSE(q.empty());
  EXPECT_EQ(q.peek(), 1); // Front should still be 1

  // Pop
  int val = q.pop();
  EXPECT_EQ(val, 1);
  EXPECT_FALSE(q.empty());
  EXPECT_EQ(q.peek(), 2); // Now front should be 2

  // Push 3
  q.push(3);
  EXPECT_EQ(q.peek(), 2);

  // Pop remaining elements
  EXPECT_EQ(q.pop(), 2);
  EXPECT_EQ(q.peek(), 3);
  EXPECT_EQ(q.pop(), 3);
  EXPECT_TRUE(q.empty());
}

TEST(MyQueue, PushPopSequence) {
  MyQueue q;
  q.push(1);
  q.push(2);
  EXPECT_EQ(q.pop(), 1);
  q.push(3);
  q.push(4);
  EXPECT_EQ(q.pop(), 2);
  EXPECT_EQ(q.pop(), 3);
  EXPECT_EQ(q.pop(), 4);
  EXPECT_TRUE(q.empty());
}
