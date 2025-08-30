#ifndef LC_00226_INVERT_BINARY_TREE_SOLUTION_H
#define LC_00226_INVERT_BINARY_TREE_SOLUTION_H

#include <vector>

// Definition for a binary tree node.
struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode() : val(0), left(nullptr), right(nullptr) {}
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
    TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
};

class Solution {
public:
    TreeNode* invertTree(TreeNode* root);
};

#endif //LC_00226_INVERT_BINARY_TREE_SOLUTION_H