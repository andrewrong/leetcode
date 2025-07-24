#include "search.h"

class BinarySearch: public Search {
    public:
    int BinarySearch1(const vector<int>& nums, int val) override {
        //[) 左闭右开
        if (nums.size() == 0) {
            return -1;
        }         
        
        auto left = 0;
        auto right = nums.size();
        
        while(left < right) {
            auto mid = (left + right) / 2;
            
            if (nums.at(mid) > val) {
                right = mid;     
            } else if (nums.at(mid) < val) {
                left = mid+1;
            } else {
                return mid;
            }
        }
        
        return -1;
    } 
    
    int BinarySearch2(const vector<int>& nums, int val) override {
        // [] 闭区间
        if (nums.size() == 0) {
            return -1;
        }
        
        auto left = 0;
        auto right = nums.size() - 1;
        
        while(left <= right) {
            auto mid = (left + right) /2;
            if (nums.at(mid) > val) {
                right = mid-1;     
            } else if (nums.at(mid) < val) {
                left = mid+1;
            } else {
                return mid;
            }
        }
        
        return -1;
    }
};