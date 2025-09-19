#include <vector>
#include <algorithm>

class Solution {
public:
    std::vector<std::vector<int>> subsetsWithDup(std::vector<int>& nums) {
        // Sort the array to handle duplicates properly
        std::sort(nums.begin(), nums.end());
        
        std::vector<std::vector<int>> result;
        std::vector<int> current;
        
        // Define a helper function for backtracking
        std::function<void(int)> backtrack = [&](int start) {
            // Add a copy of the current subset to the result
            result.push_back(current);
            
            // Iterate through the remaining elements
            for (int i = start; i < nums.size(); i++) {
                // Skip duplicates
                if (i > start && nums[i] == nums[i-1]) {
                    continue;
                }
                
                // Include the current element
                current.push_back(nums[i]);
                
                // Recurse with the next index
                backtrack(i + 1);
                
                // Backtrack by removing the current element
                current.pop_back();
            }
        };
        
        // Start the backtracking process
        backtrack(0);
        
        return result;
    }
};