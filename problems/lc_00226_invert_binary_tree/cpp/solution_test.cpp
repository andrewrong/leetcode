#include "solution.h"
#include <gtest/gtest.h>

// Helper function to create a tree from a vector representation
TreeNode* createTree(const std::vector<int>& values) {
    if (values.empty() || values[0] == NULL) return nullptr;
    
    TreeNode* root = new TreeNode(values[0]);
    std::queue<TreeNode*> q;
    q.push(root);
    
    int i = 1;
    while (!q.empty() && i < values.size()) {
        TreeNode* node = q.front();
        q.pop();
        
        if (i < values.size() && values[i] != NULL) {
            node->left = new TreeNode(values[i]);
            q.push(node->left);
        }
        i++;
        
        if (i < values.size() && values[i] != NULL) {
            node->right = new TreeNode(values[i]);
            q.push(node->right);
        }
        i++;
    }
    
    return root;
}

// Helper function to convert tree to vector representation (level-order traversal)
std::vector<int> treeToVector(TreeNode* root) {
    std::vector<int> result;
    if (!root) return result;
    
    std::queue<TreeNode*> q;
    q.push(root);
    
    while (!q.empty()) {
        TreeNode* node = q.front();
        q.pop();
        
        if (node) {
            result.push_back(node->val);
            q.push(node->left);
            q.push(node->right);
        } else {
            result.push_back(NULL); // Representing null nodes
        }
    }
    
    // Remove trailing nulls
    while (!result.empty() && result.back() == NULL) {
        result.pop_back();
    }
    
    return result;
}

TEST(InvertBinaryTreeTest, Example1) {
    Solution solution;
    std::vector<int> input = {4, 2, 7, 1, 3, 6, 9};
    std::vector<int> expected = {4, 7, 2, 9, 6, 3, 1};
    
    TreeNode* root = createTree(input);
    TreeNode* result = solution.invertTree(root);
    std::vector<int> resultVec = treeToVector(result);
    
    EXPECT_EQ(resultVec, expected);
}

TEST(InvertBinaryTreeTest, Example2) {
    Solution solution;
    std::vector<int> input = {2, 1, 3};
    std::vector<int> expected = {2, 3, 1};
    
    TreeNode* root = createTree(input);
    TreeNode* result = solution.invertTree(root);
    std::vector<int> resultVec = treeToVector(result);
    
    EXPECT_EQ(resultVec, expected);
}

TEST(InvertBinaryTreeTest, EmptyTree) {
    Solution solution;
    TreeNode* root = nullptr;
    TreeNode* result = solution.invertTree(root);
    
    EXPECT_EQ(result, nullptr);
}