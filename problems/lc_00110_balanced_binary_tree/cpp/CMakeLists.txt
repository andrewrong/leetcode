# CMakeLists.txt for lc_00110_balanced_binary_tree

# Add the solution library
add_library(lc_00110_balanced_binary_tree_solution solution.cpp)

# Add the test executable
add_executable(lc_00110_balanced_binary_tree_test test.cpp)

# Link the solution library to the test executable
target_link_libraries(lc_00110_balanced_binary_tree_test lc_00110_balanced_binary_tree_solution)

# Include directories for the test executable
target_include_directories(lc_00110_balanced_binary_tree_test PRIVATE ${CMAKE_CURRENT_SOURCE_DIR})