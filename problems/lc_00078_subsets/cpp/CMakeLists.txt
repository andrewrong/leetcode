# CMakeLists.txt for lc_00078_subsets problem

# Add the executable for the solution
add_executable(lc_00078_subsets_solution solution.cpp)

# Link with standard libraries if needed
target_link_libraries(lc_00078_subsets_solution)

# Add the executable for the test
add_executable(lc_00078_subsets_test solution_test.cpp solution.cpp)

# Link with GoogleTest
target_link_libraries(lc_00078_subsets_test gtest gtest_main)

# Add the test to CTest
add_test(NAME lc_00078_subsets_test COMMAND lc_00078_subsets_test)