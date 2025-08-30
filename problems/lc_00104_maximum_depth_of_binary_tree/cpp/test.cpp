#include "solution.h"
#include <cassert>
#include <iostream>

int main() {
    Solution solution;

    // Test case 1: Empty tree
    TreeNode* root1 = nullptr;
    assert(solution.maxDepth(root1) == 0);

    // Test case 2: Single node
    TreeNode root2(1);
    assert(solution.maxDepth(&root2) == 1);

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
    assert(solution.maxDepth(&root3) == 3);

    // Test case 4: Right-skewed tree
    //   1
    //    \
    //     2
    //      \
    //       3
    //        \
    //         4
    TreeNode node4(4);
    TreeNode node3(3, nullptr, &node4);
    TreeNode node2(2, nullptr, &node3);
    TreeNode root4(1, nullptr, &node2);
    assert(solution.maxDepth(&root4) == 4);

    // Test case 5: Left-skewed tree
    //       1
    //      /
    //     2
    //    /
    //   3
    TreeNode node3left(3);
    TreeNode node2left(2, &node3left, nullptr);
    TreeNode root5(1, &node2left, nullptr);
    assert(solution.maxDepth(&root5) == 3);

    std::cout << "All tests passed!" << std::endl;
    return 0;
}