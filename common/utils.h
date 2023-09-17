#pragma once

#include <chrono>
#include <cstdio>
#include <iostream>
#include <map>
#include <set>
#include <sstream>
#include <string>
#include <thread>
#include <vector>
#include "common/data_type.h"

struct TaskResult {
  int count{0};
  std::map<std::string, int> datas;

  void insert(const Row& row) {
    count++;
    std::string key = row.getKey();
    if (datas.find(key) == datas.end()) {
      datas[key] = 1;
    } else {
      datas[key]++;
    }
  }

  bool operator==(const TaskResult& lhs) {
    if (this->count != lhs.count) {
      return false;
    }
    return this->datas == lhs.datas;
  }

  void print() {
    std::cout << "count:" << this->count << std::endl;
    for (auto& item : this->datas) {
      std::cout << "[" << item.first << "," << item.second << "]" << std::endl;
    }
  }
};

template <typename CMP>
void getRowsInSortArray(const Row* datas,
                        int start,
                        int end,
                        CMP cmp,
                        std::function<void(const Row* data)> fn) {
  static Row cond1Lower = {1000, 10};
  static Row cond1Upper = {1000, 49};
  static Row cond2Lower = {2000, 10};
  static Row cond2Upper = {2000, 49};
  static Row cond3Lower = {3000, 10};
  static Row cond3Upper = {3000, 49};

  do {
    auto lowIdx = std::lower_bound(datas + start, datas + end, cond1Lower, cmp);
    if (lowIdx == datas + end) {
      break;
    }
    auto upperIdx =
        std::upper_bound(datas + start, datas + end, cond1Upper, cmp);
    for (auto it = lowIdx; it < upperIdx; ++it) {
      fn(it);
    }
  } while (false);

  do {
    auto lowIdx = std::lower_bound(datas + start, datas + end, cond2Lower, cmp);
    if (lowIdx == datas + end) {
      break;
    }
    auto upperIdx =
        std::upper_bound(datas + start, datas + end, cond2Upper, cmp);
    for (auto it = lowIdx; it < upperIdx; ++it) {
      fn(it);
    }
  } while (false);

  do {
    auto lowIdx = std::lower_bound(datas + start, datas + end, cond3Lower, cmp);
    if (lowIdx == datas + end) {
      break;
    }
    auto upperIdx =
        std::upper_bound(datas + start, datas + end, cond3Upper, cmp);
    for (auto it = lowIdx; it < upperIdx; ++it) {
      fn(it);
    }
  } while (false);
}
// 并行计算
void ParallelFor(int threadCnt,
                 int start,
                 int end,
                 std::function<void(int, int, int)> fn);

template <typename CMP>
std::shared_ptr<std::multiset<const Row*, CMP>> ParallelForMerge(
    std::vector<std::shared_ptr<std::multiset<const Row*, CMP>>>& v) {
  if (v.empty()) {
    return std::make_shared<std::multiset<const Row*, CMP>>();
  }

  std::vector<std::shared_ptr<std::multiset<const Row*, CMP>>> result = v;
  while (true) {
    if (result.size() == 1) {
      break;
    }

    int n = result.size();
    int mid = n / 2;
    std::vector<std::thread> threads;
    std::vector<std::shared_ptr<std::multiset<const Row*, CMP>>> midResult;
    for (int i = 0; i < mid; i++) {
      midResult.emplace_back(
          std::make_shared<std::multiset<const Row*, CMP>>());
    }

    for (int i = 0; i < mid; ++i) {
      std::shared_ptr<std::multiset<const Row*, CMP>> first = result.back();
      result.pop_back();
      std::shared_ptr<std::multiset<const Row*, CMP>> second = result.back();
      result.pop_back();

      threads.emplace_back([&result, &midResult, first, second, i]() {
        std::merge(first->begin(), first->end(), second->begin(), second->end(),
                   std::inserter(*(midResult[i]), midResult[i]->begin()));
      });
    }

    for (auto& t : threads) {
      t.join();
    }

    result.insert(result.end(), midResult.begin(), midResult.end());
  }
  return result[0];
}

template <typename Func>
void measureFuncTime(const std::string& name, Func fn) {
  auto start = std::chrono::high_resolution_clock::now();
  fn();
  auto end = std::chrono::high_resolution_clock::now();
  auto duration = std::chrono::duration_cast<std::chrono::milliseconds>(
      end - start);  // 计算耗时
  std::cout << "[" << name << "] Time taken:" << duration.count()
            << " milliseconds" << std::endl;
}
