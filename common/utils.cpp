#include "common/utils.h"
#include <thread>

void ParallelFor(int threadCnt,
                 int start,
                 int end,
                 std::function<void(int, int, int)> fn) {
  if (threadCnt <= 0) {
    threadCnt = std::thread::hardware_concurrency();
  }
  std::cout << "thread: " << threadCnt << std::endl;
  int nthreads = threadCnt;
  int n = end - start;
  int block_size = (n + nthreads - 1) / nthreads;
  std::vector<std::thread> threads;
  for (int i = 0; i < nthreads; ++i) {
    int s = start + i * block_size;
    int e = std::min(s + block_size, end);
    threads.emplace_back([s, e, i, &fn]() {
      if (s >= e) {
        return;
      }
      fn(i, s, e);
    });
  }
  for (auto& t : threads) {
    t.join();
  }
}
