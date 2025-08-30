#include "solution.h"
#include <unordered_map>

TreeNode* buildTreeHelper(vector<int>& inorder, vector<int>& postorder, 
                          unordered_map<int, int>& inorderMap, 
                          int inStart, int inEnd, int postStart, int postEnd) {
    if (inStart > inEnd || postStart > postEnd) {
        return nullptr;
    }
    
    int rootVal = postorder[postEnd];
    TreeNode* root = new TreeNode(rootVal);
    
    int rootIndex = inorderMap[rootVal];
    int leftSize = rootIndex - inStart;
    
    root->left = buildTreeHelper(inorder, postorder, inorderMap,
                                inStart, rootIndex - 1, 
                                postStart, postStart + leftSize - 1);
    
    root->right = buildTreeHelper(inorder, postorder, inorderMap,
                                 rootIndex + 1, inEnd,
                                 postStart + leftSize, postEnd - 1);
    
    return root;
}

TreeNode* Solution::buildTree(vector<int>& inorder, vector<int>& postorder) {
    if (inorder.empty() || postorder.empty()) {
        return nullptr;
    }
    
    unordered_map<int, int> inorderMap;
    for (int i = 0; i < inorder.size(); i++) {
        inorderMap[inorder[i]] = i;
    }
    
    return buildTreeHelper(inorder, postorder, inorderMap, 
                          0, inorder.size() - 1, 
                          0, postorder.size() - 1);
}