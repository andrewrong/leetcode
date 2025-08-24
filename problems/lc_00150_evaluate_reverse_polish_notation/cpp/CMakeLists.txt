#
# leetcode/problems/lc_00150_evaluate_reverse_polish_notation/cpp/CMakeLists.txt
#

# 1. test target
add_executable(test_lc_00150 test.cpp solution.cpp)

# 2. include directories
target_include_directories(
  test_lc_00150
  PRIVATE ${PROJECT_SOURCE_DIR}/kit/cpp)

# 3. link libraries
target_link_libraries(test_lc_00150 gtest_main)

# 4. add gtest
add_test(NAME test_lc_00150 COMMAND test_lc_00150)
