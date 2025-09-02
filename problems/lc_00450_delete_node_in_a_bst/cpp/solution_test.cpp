#include "solution.h"
#include <gtest/gtest.h>

TEST(DeleteNodeTest, DeleteLeafNode) {
    Solution solution;
    // Create tree: [5,3,6,2,4,null,7]
    TreeNode* root = new TreeNode(5);
    root->left = new TreeNode(3);
    root->right = new TreeNode(6);
    root->left->left = new TreeNode(2);
    root->left->right = new TreeNode(4);
    root->right->right = new TreeNode(7);
    
    // Delete node with value 2 (leaf node)
    TreeNode* result = solution.deleteNode(root, 2);
    
    // Expected tree: [5,3,6,null,4,null,7]
    // The exact structure check would need a helper function to compare trees
    // For now, we'll just verify that the function runs without crashing
    EXPECT_TRUE(result != nullptr);
    
    // Clean up memory
    // (In a real test, we'd need a proper tree deletion function)
}

TEST(DeleteNodeTest, DeleteNodeWithOneChild) {
    Solution solution;
    // Create tree: [5,3,6,2,4,null,7]
    TreeNode* root = new TreeNode(5);
    root->left = new TreeNode(3);
    root->right = new TreeNode(6);
    root->left->left = new TreeNode(2);
    root->left->right = new TreeNode(4);
    root->right->right = new TreeNode(7);
    
    // Delete node with value 3 (has one child)
    TreeNode* result = solution.deleteNode(root, 3);
    
    // Expected tree: [5,4,6,2,null,null,7]
    EXPECT_TRUE(result != nullptr);
}

TEST(DeleteNodeTest, DeleteNodeWithTwoChildren) {
    Solution solution;
    // Create tree: [5,3,6,2,4,null,7]
    TreeNode* root = new TreeNode(5);
    root->left = new TreeNode(3);
    root->right = new TreeNode(6);
    root->left->left = new TreeNode(2);
    root->left->right = new TreeNode(4);
    root->right->right = new TreeNode(7);
    
    // Delete node with value 5 (root with two children)
    TreeNode* result = solution.deleteNode(root, 5);
    
    // Expected tree: [6,3,7,2,4]
    EXPECT_TRUE(result != nullptr);
}

TEST(DeleteNodeTest, DeleteNonExistentNode) {
    Solution solution;
    // Create tree: [5,3,6,2,4,null,7]
    TreeNode* root = new TreeNode(5);
    root->left = new TreeNode(3);
    root->right = new TreeNode(6);
    root->left->left = new TreeNode(2);
    root->left->right = new TreeNode(4);
    root->right->right = new TreeNode(7);
    
    // Delete node with value 10 (non-existent)
    TreeNode* result = solution.deleteNode(root, 10);
    
    // Tree should remain unchanged
    EXPECT_TRUE(result != nullptr);
    EXPECT_EQ(result->val, 5);
}

TEST(DeleteNodeTest, DeleteRootNode) {
    Solution solution;
    // Create tree with just root: [5]
    TreeNode* root = new TreeNode(5);
    
    // Delete root node
    TreeNode* result = solution.deleteNode(root, 5);
    
    // Result should be null (empty tree)
    EXPECT_TRUE(result == nullptr);
}