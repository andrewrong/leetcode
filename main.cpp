#include <algorithm>
#include <atomic>
#include <chrono>
#include <fstream>
#include <iostream>
#include <limits>
#include <random>
#include <vector>
#include "common/data_type.h"
#include "common/utils.h"

using namespace std;

TaskResult task1(const Row* rows, int nrows) {
  if (nrows <= 0) {
    return TaskResult{};
  }

  vector<vector<int>> results;
  atomic<int> validCount(0);
  auto threadNum = std::thread::hardware_concurrency() - 1;

  for (int i = 0; i < threadNum; ++i) {
    results.emplace_back(vector<int>());
  }

  measureFuncTime("task1-phase-1", [&]() {
    ParallelFor(threadNum, 0, nrows, [&](int threadIdx, int start, int end) {
      for (int i = start; i < end; ++i) {
        if (rows[i].Valid()) {
          validCount++;
          results[threadIdx].push_back(i);
        }
      }
    });
  });

  // print result
  TaskResult taskResult;
  measureFuncTime("task1-phase-2", [&]() {
    std::cout << "task1 count: " << validCount << endl;
    for (auto& item : results) {
      for (auto& idx : item) {
        // std::cout << rows[idx].ToString() << endl;
        taskResult.insert(rows[idx]);
      }
    }
  });
  return taskResult;
}

TaskResult task2(const Row* rows, int nrows) {
  if (nrows <= 0) {
    return TaskResult{};
  }

  vector<vector<const Row*>> results;
  atomic<int> validCount(0);
  auto threadNum = std::thread::hardware_concurrency() - 1;

  for (int i = 0; i < threadNum; ++i) {
    results.emplace_back(vector<const Row*>());
  }

  measureFuncTime("task2-phase-1", [&]() {
    ParallelFor(threadNum, 0, nrows, [&](int threadIdx, int start, int end) {
      getRowsInSortArray(rows, start, end, LessAB(), [&](const Row* row) {
        validCount++;
        results[threadIdx].push_back(row);
      });
    });
  });

  // print result
  TaskResult taskResult;
  measureFuncTime("task2-phase-2", [&]() {
    std::cout << "task2 count: " << validCount << endl;
    for (auto& item : results) {
      for (auto& idx : item) {
        // std::cout << idx->ToString() << endl;
        taskResult.insert(*idx);
      }
    }
  });

  return taskResult;
}

// 希望按照 b 的顺序来进行排序输出
TaskResult task3(const Row* rows, int nrows) {
  if (nrows <= 0) {
    return TaskResult{};
  }

  vector<shared_ptr<BContainer>> results;
  atomic<int> validCount(0);
  auto threadNum = std::thread::hardware_concurrency() - 1;

  for (int i = 0; i < threadNum; ++i) {
    results.emplace_back(make_shared<BContainer>());
  }

  // 根据 B 的值进行分组
  measureFuncTime("task3-phase-1", [&]() {
    ParallelFor(threadNum, 0, nrows, [&](int threadIdx, int start, int end) {
      getRowsInSortArray(rows, start, end, LessAB(), [&](const Row* row) {
        validCount++;
        results[threadIdx]->insert(row);
      });
    });
  });

  auto merge = results[0];
  measureFuncTime("task3-phase-2", [&]() {
    for (int i = 1; i < threadNum; ++i) {
      merge = BContainer::merge(merge, results[i]);
    }
  });

  TaskResult taskResult;
  // print result
  measureFuncTime("task3-phase-3", [&]() {
    std::cout << "task3 count: " << validCount << endl;
    merge->iterator([&](const Row* idx) {
      // std::cout << idx->ToString() << endl;
      taskResult.insert(*idx);
    });
  });
  return taskResult;
}

shared_ptr<vector<Row>> readRows(const string& filename) {
  auto data = make_shared<vector<Row>>();
  string line;
  ifstream file(filename);
  while (getline(file, line)) {
    istringstream iss(line);
    string token;
    getline(iss, token, ',');
    auto a = stoi(token);
    getline(iss, token, ',');
    auto b = stoi(token);
    data->push_back(Row(a, b));
  }
  return data;
}

void geneTestData(const string& filename) {
  ofstream file(filename);
  random_device rd;
  mt19937 gen(rd());
  uniform_int_distribution<int> a_dist(3001, 10000);
  uniform_int_distribution<int> b_dist(0, 60);
  for (int i = 0; i < 5000; i++) {
    file << "1000," << b_dist(gen) << std::endl;
  }
  for (int i = 0; i < 5000; i++) {
    file << "2000," << b_dist(gen) << std::endl;
  }
  for (int i = 0; i < 5000; i++) {
    file << "3000," << b_dist(gen) << std::endl;
  }
  for (int i = 0; i < 5000; i++) {
    file << a_dist(gen) << "," << b_dist(gen) << std::endl;
  }
  file.close();
}

int main() {
  geneTestData("./data.txt");
  TaskResult t1, t2, t3;
  auto rows = readRows("./data.txt");
  measureFuncTime("task1", [&]() { t1 = task1(rows->data(), rows->size()); });
  std::sort(rows->begin(), rows->end(), LessAB());
  measureFuncTime("task2", [&]() { t2 = task2(rows->data(), rows->size()); });
  measureFuncTime("task3", [&]() { t3 = task3(rows->data(), rows->size()); });

  if (t1 == t2 && t2 == t3) {
    cout << "this is equal" << endl;
  }
  return 0;
}
