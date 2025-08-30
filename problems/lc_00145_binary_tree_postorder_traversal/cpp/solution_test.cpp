#include "solution.h"
#include <vector>
#include <cassert>
#include "../../kit/cpp/tree_node.h"

int main() {
    Solution solution;
    
    // Test case 1: [1,null,2,3] => [3,2,1]
    TreeNode* root1 = new TreeNode(1);
    root1->right = new TreeNode(2);
    root1->right->left = new TreeNode(3);
    
    std::vector<int> expected1 = {3, 2, 1};
    std::vector<int> result1 = solution.postorderTraversal(root1);
    assert(result1 == expected1);
    
    // Test case 2: [] => []
    TreeNode* root2 = nullptr;
    std::vector<int> expected2 = {};
    std::vector<int> result2 = solution.postorderTraversal(root2);
    assert(result2 == expected2);
    
    // Test case 3: [1] => [1]
    TreeNode* root3 = new TreeNode(1);
    std::vector<int> expected3 = {1};
    std::vector<int> result3 = solution.postorderTraversal(root3);
    assert(result3 == expected3);
    
    // Clean up memory
    delete root1->right->left;
    delete root1->right;
    delete root1;
    delete root3;
    
    return 0;
}