#
# leetcode/problems/lc_00028_find_the_index_of_the_first_occurrence_in_a_string/CMakeLists.txt
#

# 1. test target
add_executable(test_lc_00028 test.cpp solution.cpp)

# 2. include directories
target_include_directories(
  test_lc_00028
  PRIVATE ${PROJECT_SOURCE_DIR}/kit/cpp)

# 3. link libraries
target_link_libraries(test_lc_00028 gtest_main)

# 4. add gtest
add_test(NAME test_lc_00028 COMMAND test_lc_00028)
