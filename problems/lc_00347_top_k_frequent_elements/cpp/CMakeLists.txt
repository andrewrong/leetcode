# leetcode/problems/lc_00347_top_k_frequent_elements/cpp/CMakeLists.txt

# Add the solution file to the library
add_library(lc_00347_top_k_frequent_elements solution.cpp)

# Enable C++11
set_target_properties(lc_00347_top_k_frequent_elements PROPERTIES CXX_STANDARD 11)

# Add include directories
target_include_directories(lc_00347_top_k_frequent_elements PUBLIC ${CMAKE_CURRENT_SOURCE_DIR})

# Add the test executable
add_executable(test_lc_00347_top_k_frequent_elements solution_test.cpp)

# Link the test executable against Google Test
target_link_libraries(test_lc_00347_top_k_frequent_elements gtest_main)

# Add the test to CTest
add_test(NAME test_lc_00347_top_k_frequent_elements COMMAND test_lc_00347_top_k_frequent_elements)
