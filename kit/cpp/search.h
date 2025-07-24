#include <vector>
using std::vector;

class Search {
   public:
   virtual int BinarySearch1(const vector<int>& num, int val) = 0; 
   virtual int BinarySearch2(const vector<int>& num, int val) = 0; 
   
   virtual ~Search(){};
};
