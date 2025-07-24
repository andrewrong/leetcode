#include <vector>
using std::vector;

// LeetCode Problem 27: Remove Element
class Solution {
public:
    int Solve(vector<int>& nums, int val) {
        int size = 0;
        int prev = 0;    
        for (auto i = 0; i < nums.size(); i++) {
            if (nums.at(i) != val) {
                size+=1;
                nums[prev++] = nums.at(i); 
            }
        }
        return size;   
    }
};