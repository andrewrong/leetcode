# CMakeLists.txt for lc_00102_binary_tree_level_order_traversal
set(PROJECT_NAME lc_00102_binary_tree_level_order_traversal)

project(${PROJECT_NAME})

# Add solution.cpp and solution.h
add_library(solution_cpp solution.cpp solution.h)

# Link with kit libraries if needed
target_link_libraries(solution_cpp PRIVATE kit_cpp)