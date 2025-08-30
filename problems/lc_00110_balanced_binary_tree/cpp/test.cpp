#include "solution.h"
#include <cassert>
#include <iostream>

int main() {
    Solution solution;

    // Test case 1: Empty tree
    TreeNode* root1 = nullptr;
    assert(solution.isBalanced(root1) == true);

    // Test case 2: Single node
    TreeNode root2(1);
    assert(solution.isBalanced(&root2) == true);

    // Test case 3: Balanced tree
    //       3
    //      / \
    //     9   20
    //        /  \
    //       15   7
    TreeNode node15(15);
    TreeNode node7(7);
    TreeNode node20(20, &node15, &node7);
    TreeNode node9(9);
    TreeNode root3(3, &node9, &node20);
    assert(solution.isBalanced(&root3) == true);

    // Test case 4: Unbalanced tree (right subtree deeper)
    //       1
    //      / \
    //     2   2
    //    / \
    //   3   3
    //  / \
    // 4   4
    TreeNode node4Left(4);
    TreeNode node4Right(4);
    TreeNode node3Left(3, &node4Left, &node4Right);
    TreeNode node3Right(3);
    TreeNode node2Left(2, &node3Left, &node3Right);
    TreeNode node2Right(2);
    TreeNode root4(1, &node2Left, &node2Right);
    assert(solution.isBalanced(&root4) == false);

    // Test case 5: Another balanced tree
    //       1
    //      / \
    //     2   3
    //    / \
    //   4   5
    TreeNode node4(4);
    TreeNode node5(5);
    TreeNode node2(2, &node4, &node5);
    TreeNode node3(3);
    TreeNode root5(1, &node2, &node3);
    assert(solution.isBalanced(&root5) == true);

    // Test case 6: Unbalanced tree from problem example
    //       1
    //      / \
    //     2   2
    //    /   /
    //   3   3
    //  /   /
    // 4   4
    // This represents [1,2,2,3,null,null,3,4,null,null,4]
    TreeNode node4Left(4);
    TreeNode node3Left(3, &node4Left, nullptr);
    TreeNode node2Left(2, &node3Left, nullptr);
    
    TreeNode node4Right(4);
    TreeNode node3Right(3, &node4Right, nullptr);
    TreeNode node2Right(2, nullptr, &node3Right);
    
    TreeNode root6(1, &node2Left, &node2Right);
    assert(solution.isBalanced(&root6) == false);

    std::cout << "All tests passed!" << std::endl;
    return 0;
}