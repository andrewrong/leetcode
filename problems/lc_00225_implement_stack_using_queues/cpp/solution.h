// leetcode/problems/lc_00225_implement_stack_using_queues/cpp/solution.h
#pragma once

#include <queue>

class MyStack {
private:
  std::queue<int> q;

public:
  MyStack();
  void push(int x);
  int pop();
  int top();
  bool empty();
};
