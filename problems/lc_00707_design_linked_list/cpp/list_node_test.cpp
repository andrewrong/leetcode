#include "list_node.h"
#include <gtest/gtest.h>

class MyLinkedListTest : public ::testing::Test {
protected:
    void SetUp() override {
        // 在每个测试用例之前运行
    }

    void TearDown() override {
        // 在每个测试用例之后运行
    }
};

// 测试基本功能
TEST_F(MyLinkedListTest, BasicOperations) {
    MyLinkedList linkedList;
    
    // 测试空链表获取元素
    EXPECT_EQ(linkedList.get(0), -1);
    
    // 测试在头部添加元素
    linkedList.addAtHead(1);
    EXPECT_EQ(linkedList.get(0), 1);
    
    // 测试在尾部添加元素
    linkedList.addAtTail(3);
    EXPECT_EQ(linkedList.get(1), 3);
    
    // 测试在指定索引添加元素
    linkedList.addAtIndex(1, 2);
    EXPECT_EQ(linkedList.get(0), 1);
    EXPECT_EQ(linkedList.get(1), 2);
    EXPECT_EQ(linkedList.get(2), 3);
    
    // 测试删除元素
    linkedList.deleteAtIndex(1);
    EXPECT_EQ(linkedList.get(0), 1);
    EXPECT_EQ(linkedList.get(1), 3);
    EXPECT_EQ(linkedList.get(2), -1); // 超出范围
}

// 测试边界情况
TEST_F(MyLinkedListTest, EdgeCases) {
    MyLinkedList linkedList;
    
    // 测试在空链表的索引 0 处添加元素
    linkedList.addAtIndex(0, 1);
    EXPECT_EQ(linkedList.get(0), 1);
    
    // 测试在超出范围的索引处添加元素（应该无效）
    linkedList.addAtIndex(5, 5);
    EXPECT_EQ(linkedList.get(1), -1); // 应该仍然是 -1
    
    // 测试删除超出范围的索引
    linkedList.deleteAtIndex(5); // 应该无效
    EXPECT_EQ(linkedList.get(0), 1); // 链表应保持不变
    
    // 测试删除唯一的元素
    linkedList.deleteAtIndex(0);
    EXPECT_EQ(linkedList.get(0), -1); // 链表应为空
    
    // 测试添加多个元素后删除尾部元素
    linkedList.addAtTail(10);
    linkedList.addAtTail(20);
    linkedList.addAtTail(30);
    linkedList.deleteAtIndex(2); // 删除最后一个元素
    EXPECT_EQ(linkedList.get(2), -1);
    EXPECT_EQ(linkedList.get(1), 20);
}

// 测试大量元素操作
TEST_F(MyLinkedListTest, LargeDataSet) {
    MyLinkedList linkedList;
    
    // 添加大量元素到头部
    for (int i = 0; i < 100; i++) {
        linkedList.addAtHead(i);
    }
    
    // 验证元素顺序
    EXPECT_EQ(linkedList.get(0), 99);
    EXPECT_EQ(linkedList.get(99), 0);
    
    // 在中间添加元素
    linkedList.addAtIndex(50, 500);
    EXPECT_EQ(linkedList.get(50), 500);
    
    // 删除一些元素
    linkedList.deleteAtIndex(0);
    EXPECT_EQ(linkedList.get(0), 98);
    
    linkedList.deleteAtIndex(49);
    EXPECT_EQ(linkedList.get(49), 49);
}