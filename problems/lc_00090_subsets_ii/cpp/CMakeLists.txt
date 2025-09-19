# CMakeLists.txt for lc_00090_subsets_ii problem
add_library(lc_00090_subsets_ii_solution solution.cpp)

target_include_directories(lc_00090_subsets_ii_solution PUBLIC ${CMAKE_CURRENT_SOURCE_DIR})

# Create the main executable
add_executable(lc_00090_subsets_ii_test solution_test.cpp)

# Link the necessary libraries
target_link_libraries(lc_00090_subsets_ii_test lc_00090_subsets_ii_solution gtest gtest_main)

# Add the test
add_test(NAME lc_00090_subsets_ii_test COMMAND lc_00090_subsets_ii_test)