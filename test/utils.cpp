#include "common/utils.h"
#include <atomic>
#include <iostream>
#include "common/data_type.h"
#include "gtest/gtest.h"
using namespace std;

TEST(DataTypeTest, RowToString) {
  Row row(1, 2);
  EXPECT_EQ(row.a, 1);
  EXPECT_EQ(row.b, 2);
  EXPECT_EQ(row.ToString(), "{ 1, 2 }");
}

TEST(DataTypeTest, RowValid) {
  Row row(1000, 10);
  EXPECT_TRUE(row.Valid());
  Row row2(1000, 9);
  EXPECT_FALSE(row2.Valid());
  Row row3(2000, 10);
  EXPECT_TRUE(row3.Valid());
  Row row4(2000, 60);
  EXPECT_FALSE(row4.Valid());
  Row row5(3000, 10);
  EXPECT_TRUE(row5.Valid());
  Row row6(3000, 50);
  EXPECT_FALSE(row6.Valid());
  Row row7(3000, 100);
  EXPECT_FALSE(row7.Valid());
}

TEST(LessABTest, LessThan) {
  Row row1 = {1, 2};
  Row row2 = {2, 3};
  LessAB less_ab;
  EXPECT_TRUE(less_ab(row1, row2));
}

TEST(LessABTest, EqualA) {
  Row row1 = {1, 2};
  Row row2 = {1, 3};
  LessAB less_ab;
  EXPECT_TRUE(less_ab(row1, row2));
}

TEST(LessABTest, EqualAB) {
  Row row1 = {1, 2};
  Row row2 = {1, 2};
  LessAB less_ab;
  EXPECT_FALSE(less_ab(row1, row2));
}

TEST(LessABTest, GreaterThan) {
  Row row1 = {
      2,
      3,
  };
  Row row2 = {
      1,
      2,
  };
  LessAB less_ab;
  EXPECT_FALSE(less_ab(row1, row2));
}

TEST(LessABTest, GreaterThan2) {
  Row row1 = {
      2,
      3,
  };
  Row row2 = {
      2,
      2,
  };
  LessAB less_ab;
  EXPECT_FALSE(less_ab(row1, row2));
}

TEST(LessBATest, LessThan) {
  Row row1 = {1, 1};
  Row row2 = {2, 2};
  LessB less_b;
  EXPECT_TRUE(less_b(row1, row2));
}

TEST(LessBATest, LessThan2) {
  Row row1 = {2, 1};
  Row row2 = {1, 2};
  LessB less_b;
  EXPECT_TRUE(less_b(row1, row2));
}

TEST(LessBATest, EqualB) {
  Row row1 = {1, 2};
  Row row2 = {1, 2};
  LessB less_ba;
  EXPECT_FALSE(less_ba(row1, row2));
}

TEST(LessBATest, EqualB2) {
  Row row1 = {2, 2};
  Row row2 = {1, 2};
  LessB less_b;
  EXPECT_FALSE(less_b(row1, row2));
}

TEST(LessBATest, GreaterThan) {
  Row row1 = {2, 2};
  Row row2 = {2, 1};
  LessB less_b;
  EXPECT_FALSE(less_b(row1, row2));
}

TEST(LessBATest, GreaterThan2) {
  Row row1 = {2, 2};
  Row row2 = {1, 1};
  LessB less_b;
  EXPECT_FALSE(less_b(row1, row2));
}

TEST(ParallelForTest, SingleThread) {
  int sum = 0;
  ParallelFor(1, 0, 10, [&sum](int thread_id, int start, int end) {
    for (int i = start; i < end; ++i) {
      sum += 1;
    }
  });
  EXPECT_EQ(sum, 10);
}

TEST(ParallelForTest, MultipleThreads) {
  atomic<int> sum(0);
  ParallelFor(4, 0, 10, [&sum](int thread_id, int start, int end) {
    for (int i = start; i < end; ++i) {
      sum += 1;
    }
  });
  EXPECT_EQ(sum, 10);
}

TEST(ParallelForTest, EmptyRange) {
  atomic<int> sum(0);
  ParallelFor(4, 10, 10, [&sum](int thread_id, int start, int end) {
    for (int i = start; i < end; ++i) {
      sum += 1;
    }
  });
  EXPECT_EQ(sum, 0);
}

TEST(ParallelForTest, LargeRange) {
  atomic<int64_t> sum(0);
  atomic<int> count(0);
  ParallelFor(4, 0, 1000000, [&sum, &count](int idx, int start, int end) {
    std::cout << "threadIdx:" << idx << std::endl;
    for (int i = start; i < end; ++i) {
      sum.fetch_add(i);
      count.fetch_add(1);
    }
  });

  EXPECT_EQ(sum, 499999500000);
  EXPECT_EQ(count, 1000000);
}

TEST(ParallelForTest, SmallRange) {
  atomic<int64_t> sum(0);
  ParallelFor(8, 0, 3, [&sum](int idx, int start, int end) {
    for (int i = start; i < end; ++i) {
      sum.fetch_add(i);
    }
  });
  EXPECT_EQ(sum, 3);
}

TEST(ParallelForTest, ZeroThread) {
  atomic<int64_t> sum(0);
  atomic<int> count(0);
  ParallelFor(0, 0, 1000000, [&sum, &count](int idx, int start, int end) {
    for (int i = start; i < end; ++i) {
      sum.fetch_add(i);
      count.fetch_add(1);
    }
  });

  EXPECT_EQ(sum, 499999500000);
  EXPECT_EQ(count, 1000000);
}

TEST(ParallelForMergeTest, MergeTwoSets) {
  std::shared_ptr<std::multiset<const Row*, LessABPointer>> set1 =
      std::make_shared<std::multiset<const Row*, LessABPointer>>();
  std::shared_ptr<std::multiset<const Row*, LessABPointer>> set2 =
      std::make_shared<std::multiset<const Row*, LessABPointer>>();
  Row row1 = {1, 2};
  Row row2 = {2, 4};
  Row row3 = {1, 3};
  Row row4 = {2, 5};
  set1->insert(&row1);
  set1->insert(&row2);
  set2->insert(&row3);
  set2->insert(&row4);

  std::vector<std::shared_ptr<std::multiset<const Row*, LessABPointer>>> input =
      {set1, set2};
  std::shared_ptr<std::multiset<const Row*, LessABPointer>> output =
      ParallelForMerge(input);

  EXPECT_EQ(output->size(), 4);
  auto iter = output->begin();
  EXPECT_EQ((*iter)->a, 1);
  EXPECT_EQ((*iter)->b, 2);
  ++iter;
  EXPECT_EQ((*iter)->a, 1);
  EXPECT_EQ((*iter)->b, 3);
  ++iter;
  EXPECT_EQ((*iter)->a, 2);
  EXPECT_EQ((*iter)->b, 4);
  ++iter;
  EXPECT_EQ((*iter)->a, 2);
  EXPECT_EQ((*iter)->b, 5);
}

TEST(ParallelForMergeTest, MergeMultipleSets) {
  std::shared_ptr<std::multiset<const Row*, LessABPointer>> set1 =
      std::make_shared<std::multiset<const Row*, LessABPointer>>();
  std::shared_ptr<std::multiset<const Row*, LessABPointer>> set2 =
      std::make_shared<std::multiset<const Row*, LessABPointer>>();
  std::shared_ptr<std::multiset<const Row*, LessABPointer>> set3 =
      std::make_shared<std::multiset<const Row*, LessABPointer>>();
  std::shared_ptr<std::multiset<const Row*, LessABPointer>> set4 =
      std::make_shared<std::multiset<const Row*, LessABPointer>>();

  Row row1 = {1, 2};
  Row row2 = {2, 4};
  Row row3 = {1, 3};
  Row row4 = {2, 5};
  Row row5 = {3, 6};
  Row row6 = {4, 7};
  Row row7 = {5, 8};
  Row row8 = {6, 9};

  set1->insert(&row1);
  set1->insert(&row2);
  set2->insert(&row3);
  set2->insert(&row4);
  set3->insert(&row5);
  set3->insert(&row6);
  set4->insert(&row7);
  set4->insert(&row8);

  std::vector<std::shared_ptr<std::multiset<const Row*, LessABPointer>>> input =
      {set1, set2, set3, set4};
  std::shared_ptr<std::multiset<const Row*, LessABPointer>> output =
      ParallelForMerge(input);

  EXPECT_EQ(output->size(), 8);
  auto iter = output->begin();
  EXPECT_EQ((*iter)->a, 1);
  EXPECT_EQ((*iter)->b, 2);
  ++iter;
  EXPECT_EQ((*iter)->a, 1);
  EXPECT_EQ((*iter)->b, 3);
  ++iter;
  EXPECT_EQ((*iter)->a, 2);
  EXPECT_EQ((*iter)->b, 4);
  ++iter;
  EXPECT_EQ((*iter)->a, 2);
  EXPECT_EQ((*iter)->b, 5);
  ++iter;
  EXPECT_EQ((*iter)->a, 3);
  EXPECT_EQ((*iter)->b, 6);
  ++iter;
  EXPECT_EQ((*iter)->a, 4);
  EXPECT_EQ((*iter)->b, 7);
  ++iter;
  EXPECT_EQ((*iter)->a, 5);
  EXPECT_EQ((*iter)->b, 8);
}

TEST(ParallelForMergeTest, MergeEmptySet) {
  std::shared_ptr<std::multiset<const Row*, LessABPointer>> set1 =
      std::make_shared<std::multiset<const Row*, LessABPointer>>(
          std::initializer_list<const Row*>{});
  std::shared_ptr<std::multiset<const Row*, LessABPointer>> set2 =
      std::make_shared<std::multiset<const Row*, LessABPointer>>();

  Row row1 = {1, 3};
  Row row2 = {2, 5};

  set2->insert(&row1);
  set2->insert(&row2);

  std::vector<std::shared_ptr<std::multiset<const Row*, LessABPointer>>> input =
      {set1, set2};
  std::shared_ptr<std::multiset<const Row*, LessABPointer>> output =
      ParallelForMerge(input);

  EXPECT_EQ(output->size(), 2);
  auto iter = output->begin();
  EXPECT_EQ((*iter)->a, 1);
  EXPECT_EQ((*iter)->b, 3);
  ++iter;
  EXPECT_EQ((*iter)->a, 2);
  EXPECT_EQ((*iter)->b, 5);
}

TEST(ParallelForMergeTest, MergeSingleSet) {
  std::shared_ptr<std::multiset<const Row*, LessABPointer>> set1 =
      std::make_shared<std::multiset<const Row*, LessABPointer>>();

  Row row1 = {1, 2};
  Row row2 = {2, 4};
  set1->insert(&row1);
  set1->insert(&row2);
  std::vector<std::shared_ptr<std::multiset<const Row*, LessABPointer>>> input =
      {set1};
  std::shared_ptr<std::multiset<const Row*, LessABPointer>> output =
      ParallelForMerge(input);

  EXPECT_EQ(output->size(), 2);
  auto iter = output->begin();
  EXPECT_EQ((*iter)->a, 1);
  EXPECT_EQ((*iter)->b, 2);
  ++iter;
  EXPECT_EQ((*iter)->a, 2);
  EXPECT_EQ((*iter)->b, 4);
}

TEST(ParallelForMergeTest, MergeEmptyInput) {
  std::vector<std::shared_ptr<std::multiset<const Row*, LessABPointer>>> input =
      {};
  std::shared_ptr<std::multiset<const Row*, LessABPointer>> output =
      ParallelForMerge(input);

  EXPECT_EQ(output->size(), 0);
}

TEST(GetRowsInSortArrayTest, SingleCondition) {
  std::vector<Row> rows = {
      {1000, 5},  {1000, 10}, {1000, 15}, {1000, 20}, {1000, 25}, {1000, 30},
      {1000, 35}, {1000, 40}, {1000, 45}, {1000, 50}, {2000, 5},  {2000, 10},
      {2000, 15}, {2000, 20}, {2000, 25}, {2000, 30}, {2000, 35}, {2000, 40},
      {2000, 45}, {2000, 50}, {3000, 5},  {3000, 10}, {3000, 15}, {3000, 20},
      {3000, 25}, {3000, 30}, {3000, 35}, {3000, 40}, {3000, 45}, {3000, 50},
  };
  std::vector<const Row*> results;
  auto fn = [&results](const Row* data) { results.push_back(data); };
  getRowsInSortArray(rows.data(), 0, rows.size(), LessAB(), fn);

  EXPECT_EQ(results.size(), 24);
}

TEST(GetRowsInSortArrayTest, MultipleConditions) {
  std::vector<Row> rows = {
      {500, 5},   {500, 10},   {500, 15},   {500, 20},   {500, 25},
      {1000, 5},  {1000, 10},  {1000, 15},  {1000, 20},  {1000, 25},
      {1500, 5},  {1500, 10},  {1500, 15},  {1500, 20},  {1500, 25},
      {2000, 5},  {2000, 10},  {2000, 15},  {2000, 20},  {2000, 25},
      {2500, 5},  {2500, 10},  {2500, 15},  {2500, 20},  {2500, 25},
      {3000, 5},  {3000, 10},  {3000, 15},  {3000, 20},  {3000, 25},
      {3500, 5},  {3500, 10},  {3500, 15},  {3500, 20},  {3500, 25},
      {4000, 5},  {4000, 10},  {4000, 15},  {4000, 20},  {4000, 25},
      {4500, 5},  {4500, 10},  {4500, 15},  {4500, 20},  {4500, 25},
      {5000, 5},  {5000, 10},  {5000, 15},  {5000, 20},  {5000, 25},
      {5500, 5},  {5500, 10},  {5500, 15},  {5500, 20},  {5500, 25},
      {6000, 5},  {6000, 10},  {6000, 15},  {6000, 20},  {6000, 25},
      {6500, 5},  {6500, 10},  {6500, 15},  {6500, 20},  {6500, 25},
      {7000, 5},  {7000, 10},  {7000, 15},  {7000, 20},  {7000, 25},
      {7500, 5},  {7500, 10},  {7500, 15},  {7500, 20},  {7500, 25},
      {8000, 5},  {8000, 10},  {8000, 15},  {8000, 20},  {8000, 25},
      {8500, 5},  {8500, 10},  {8500, 15},  {8500, 20},  {8500, 25},
      {9000, 5},  {9000, 10},  {9000, 15},  {9000, 20},  {9000, 25},
      {9500, 5},  {9500, 10},  {9500, 15},  {9500, 20},  {9500, 25},
      {10000, 5}, {10000, 10}, {10000, 15}, {10000, 20}, {10000, 25},
  };
  std::vector<const Row*> results;
  auto fn = [&results](const Row* data) { results.push_back(data); };
  getRowsInSortArray(rows.data(), 0, rows.size(), LessAB(), fn);

  EXPECT_EQ(results.size(), 12);
  for (int i = 0; i < 4; ++i) {
    EXPECT_EQ(results[i]->a, 1000);
    EXPECT_EQ(results[i]->b, (i + 2) * 5);
  }
  for (int i = 4; i < 8; ++i) {
    EXPECT_EQ(results[i]->a, 2000);
    EXPECT_EQ(results[i]->b, (i - 4 + 2) * 5);
  }
  for (int i = 8; i < 12; ++i) {
    EXPECT_EQ(results[i]->a, 3000);
    EXPECT_EQ(results[i]->b, (i - 8 + 2) * 5);
  }
}

TEST(GetRowsInSortArrayTest, NoMatches) {
  std::vector<Row> rows = {
      {500, 5}, {500, 10}, {500, 15}, {500, 20}, {500, 25},
      {600, 5}, {600, 10}, {600, 15}, {600, 20}, {600, 25},
  };
  std::vector<const Row*> results;
  auto fn = [&results](const Row* data) { results.push_back(data); };
  getRowsInSortArray(rows.data(), 0, rows.size(), LessAB(), fn);

  EXPECT_EQ(results.size(), 0);
}

TEST(GetRowsInSortArrayTest, EmptyInput) {
  std::vector<Row> rows = {};
  std::vector<const Row*> results;
  auto fn = [&results](const Row* data) { results.push_back(data); };
  getRowsInSortArray(rows.data(), 0, rows.size(), LessAB(), fn);

  EXPECT_EQ(results.size(), 0);
}

TEST(BContainerTest, InsertAndIterator) {
  std::vector<Row> rows = {
      {1, 1},
      {1, 3},
      {2, 4},
      {2, 2},
  };
  std::shared_ptr<BContainer> container = std::make_shared<BContainer>();
  for (auto& row : rows) {
    container->insert(&row);
  }

  std::vector<const Row*> results;
  auto fn = [&results](const Row* data) { results.push_back(data); };
  container->iterator(fn);

  EXPECT_EQ(results.size(), 4);
  EXPECT_EQ(results[0]->a, 1);
  EXPECT_EQ(results[0]->b, 1);
  EXPECT_EQ(results[1]->a, 2);
  EXPECT_EQ(results[1]->b, 2);
  EXPECT_EQ(results[2]->a, 1);
  EXPECT_EQ(results[2]->b, 3);
  EXPECT_EQ(results[3]->a, 2);
  EXPECT_EQ(results[3]->b, 4);
}

TEST(BContainerTest, Merge) {
  std::vector<Row> rows1 = {
      {1, 1},
      {1, 3},
  };
  std::vector<Row> rows2 = {
      {2, 2},
      {2, 4},
  };
  std::shared_ptr<BContainer> container1 = std::make_shared<BContainer>();
  std::shared_ptr<BContainer> container2 = std::make_shared<BContainer>();
  for (auto& row : rows1) {
    container1->insert(&row);
  }
  for (auto& row : rows2) {
    container2->insert(&row);
  }

  std::shared_ptr<BContainer> merged =
      BContainer::merge(container1, container2);

  std::vector<const Row*> results;
  auto fn = [&results](const Row* data) { results.push_back(data); };
  merged->iterator(fn);

  EXPECT_EQ(results.size(), 4);
  EXPECT_EQ(results[0]->a, 1);
  EXPECT_EQ(results[0]->b, 1);
  EXPECT_EQ(results[1]->a, 2);
  EXPECT_EQ(results[1]->b, 2);
  EXPECT_EQ(results[2]->a, 1);
  EXPECT_EQ(results[2]->b, 3);
  EXPECT_EQ(results[3]->a, 2);
  EXPECT_EQ(results[3]->b, 4);
}

TEST(BContainerTest, MergeEmpty) {
  std::vector<Row> rows = {
      {1, 1},
      {1, 3},
  };
  std::shared_ptr<BContainer> container1 = std::make_shared<BContainer>();
  std::shared_ptr<BContainer> container2 = std::make_shared<BContainer>();
  for (auto& row : rows) {
    container1->insert(&row);
  }

  std::shared_ptr<BContainer> merged1 =
      BContainer::merge(container1, container2);

  std::vector<const Row*> results1;
  auto fn1 = [&results1](const Row* data) { results1.push_back(data); };
  merged1->iterator(fn1);

  EXPECT_EQ(results1.size(), 2);
  EXPECT_EQ(results1[0]->a, 1);
  EXPECT_EQ(results1[0]->b, 1);
  EXPECT_EQ(results1[1]->a, 1);
  EXPECT_EQ(results1[1]->b, 3);
}