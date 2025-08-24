// leetcode/problems/lc_00232_implement_queue_using_stacks/cpp/solution.h
#pragma once

#include <stack>

class MyQueue {
private:
  std::stack<int> stackIn;
  std::stack<int> stackOut;
  void transfer();

public:
  MyQueue();
  void push(int x);
  int pop();
  int peek();
  bool empty();
};
