#ifndef LC_00106_CONSTRUCT_BINARY_TREE_FROM_INORDER_AND_POSTORDER_TRAVERSAL_SOLUTION_H
#define LC_00106_CONSTRUCT_BINARY_TREE_FROM_INORDER_AND_POSTORDER_TRAVERSAL_SOLUTION_H

#include <vector>
#include "../../kit/cpp/tree_node.h"

using namespace std;

class Solution {
public:
    TreeNode* buildTree(vector<int>& inorder, vector<int>& postorder);
};

#endif // LC_00106_CONSTRUCT_BINARY_TREE_FROM_INORDER_AND_POSTORDER_TRAVERSAL_SOLUTION_H