#include "solution.h"
#include <gtest/gtest.h>

TEST(TrimBSTTest, Test1) {
    // Test case 1: Normal case
    // Tree:     1
    //          / \
    //         0   2
    //       low=1, high=2
    // Expected:  1
    //             \
    //              2
    TreeNode* root = new TreeNode(1);
    root->left = new TreeNode(0);
    root->right = new TreeNode(2);
    
    Solution solution;
    TreeNode* result = solution.trimBST(root, 1, 2);
    
    EXPECT_EQ(result->val, 1);
    EXPECT_EQ(result->left, nullptr);
    EXPECT_NE(result->right, nullptr);
    EXPECT_EQ(result->right->val, 2);
    EXPECT_EQ(result->right->left, nullptr);
    EXPECT_EQ(result->right->right, nullptr);
}

TEST(TrimBSTTest, Test2) {
    // Test case 2: More complex case
    // Tree:       3
    //            / \
    //           0   4
    //            \
    //             2
    //            /
    //           1
    //       low=1, high=3
    // Expected:   3
    //            /
    //           1
    //            \
    //             2
    TreeNode* root = new TreeNode(3);
    root->left = new TreeNode(0);
    root->left->right = new TreeNode(2);
    root->left->right->left = new TreeNode(1);
    root->right = new TreeNode(4);
    
    Solution solution;
    TreeNode* result = solution.trimBST(root, 1, 3);
    
    EXPECT_EQ(result->val, 3);
    EXPECT_NE(result->left, nullptr);
    EXPECT_EQ(result->left->val, 1);
    EXPECT_EQ(result->left->left, nullptr);
    EXPECT_NE(result->left->right, nullptr);
    EXPECT_EQ(result->left->right->val, 2);
    EXPECT_EQ(result->left->right->left, nullptr);
    EXPECT_EQ(result->left->right->right, nullptr);
    EXPECT_EQ(result->right, nullptr);
}

TEST(TrimBSTTest, Test3) {
    // Test case 3: All nodes within range
    // Tree:   2
    //        / \
    //       1   3
    //     low=0, high=4
    // Expected:  2
    //           / \
    //          1   3
    TreeNode* root = new TreeNode(2);
    root->left = new TreeNode(1);
    root->right = new TreeNode(3);
    
    Solution solution;
    TreeNode* result = solution.trimBST(root, 0, 4);
    
    EXPECT_EQ(result->val, 2);
    EXPECT_NE(result->left, nullptr);
    EXPECT_EQ(result->left->val, 1);
    EXPECT_NE(result->right, nullptr);
    EXPECT_EQ(result->right->val, 3);
}

TEST(TrimBSTTest, Test4) {
    // Test case 4: Empty tree
    Solution solution;
    TreeNode* result = solution.trimBST(nullptr, 1, 2);
    
    EXPECT_EQ(result, nullptr);
}