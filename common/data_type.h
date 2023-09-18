#pragma once

#include <iostream>
#include <list>
#include <sstream>
#include <string>
#include <vector>

// 表示一行数据
typedef struct Row {
  int a;
  int b;

  Row(int a, int b) : a(a), b(b) {}
  std::string ToString() const {
    std::ostringstream oss;
    oss << "{ " << a << ", " << b << " }";
    return oss.str();
  }

  // for test
  std::string getKey() const {
    return std::to_string(a) + "_" + std::to_string(b);
  }

  bool Valid() const {
    if (!(a == 1000 || a == 2000 || a == 3000)) {
      return false;
    }
    if (b >= 10 && b < 50) {
      return true;
    }
    return false;
  }

} Row;

// a,b sort compare
struct LessAB {
  bool operator()(const Row& lhs, const Row& rhs) const {
    return (lhs.a < rhs.a) || (lhs.a == rhs.a && lhs.b < rhs.b);
  }
};
struct LessABPointer {
  bool operator()(const Row* lhs, const Row* rhs) const {
    return (lhs->a < rhs->a) || (lhs->a == rhs->a && lhs->b < rhs->b);
  }
};

struct LessB {
  bool operator()(const Row& lhs, const Row& rhs) const {
    return (lhs.b < rhs.b);
  }
};

struct LessBPointer {
  bool operator()(const Row* lhs, const Row* rhs) const {
    return (lhs->b < rhs->b);
  }
};
