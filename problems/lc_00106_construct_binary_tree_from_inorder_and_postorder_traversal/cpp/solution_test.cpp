#include "solution.h"
#include "../../kit/cpp/tree_node.h"
#include <cassert>
#include <iostream>

void printTree(TreeNode* root) {
    if (!root) return;
    cout << root->val << " ";
    printTree(root->left);
    printTree(root->right);
}

bool treesEqual(TreeNode* t1, TreeNode* t2) {
    if (!t1 && !t2) return true;
    if (!t1 || !t2) return false;
    if (t1->val != t2->val) return false;
    return treesEqual(t1->left, t2->left) && treesEqual(t1->right, t2->right);
}

void deleteTree(TreeNode* root) {
    if (!root) return;
    deleteTree(root->left);
    deleteTree(root->right);
    delete root;
}

int main() {
    Solution solution;
    
    // Test case 1: Normal case
    vector<int> inorder1 = {9, 3, 15, 20, 7};
    vector<int> postorder1 = {9, 15, 7, 20, 3};
    
    TreeNode* expected1 = new TreeNode(3);
    expected1->left = new TreeNode(9);
    expected1->right = new TreeNode(20);
    expected1->right->left = new TreeNode(15);
    expected1->right->right = new TreeNode(7);
    
    TreeNode* result1 = solution.buildTree(inorder1, postorder1);
    assert(treesEqual(result1, expected1));
    deleteTree(result1);
    deleteTree(expected1);
    
    // Test case 2: Single node
    vector<int> inorder2 = {1};
    vector<int> postorder2 = {1};
    
    TreeNode* expected2 = new TreeNode(1);
    
    TreeNode* result2 = solution.buildTree(inorder2, postorder2);
    assert(treesEqual(result2, expected2));
    deleteTree(result2);
    deleteTree(expected2);
    
    // Test case 3: Empty tree
    vector<int> inorder3 = {};
    vector<int> postorder3 = {};
    
    TreeNode* result3 = solution.buildTree(inorder3, postorder3);
    assert(result3 == nullptr);
    
    // Test case 4: Left-skewed tree
    vector<int> inorder4 = {4, 3, 2, 1};
    vector<int> postorder4 = {4, 3, 2, 1};
    
    TreeNode* expected4 = new TreeNode(1);
    expected4->left = new TreeNode(2);
    expected4->left->left = new TreeNode(3);
    expected4->left->left->left = new TreeNode(4);
    
    TreeNode* result4 = solution.buildTree(inorder4, postorder4);
    assert(treesEqual(result4, expected4));
    deleteTree(result4);
    deleteTree(expected4);
    
    // Test case 5: Right-skewed tree
    vector<int> inorder5 = {1, 2, 3, 4};
    vector<int> postorder5 = {4, 3, 2, 1};
    
    TreeNode* expected5 = new TreeNode(1);
    expected5->right = new TreeNode(2);
    expected5->right->right = new TreeNode(3);
    expected5->right->right->right = new TreeNode(4);
    
    TreeNode* result5 = solution.buildTree(inorder5, postorder5);
    assert(treesEqual(result5, expected5));
    deleteTree(result5);
    deleteTree(expected5);
    
    cout << "All tests passed!" << endl;
    
    return 0;
}