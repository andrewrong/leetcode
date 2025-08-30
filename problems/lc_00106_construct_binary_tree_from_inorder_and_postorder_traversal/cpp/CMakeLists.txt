# CMakeLists.txt for lc_00106_construct_binary_tree_from_inorder_and_postorder_traversal

# Get the directory name
get_filename_component(DIR_NAME ${CMAKE_CURRENT_SOURCE_DIR} NAME)

# Add the solution library
add_library(${DIR_NAME}_solution solution.cpp)

# Link with the kit library
target_link_libraries(${DIR_NAME}_solution kit)

# Create the test executable
add_executable(${DIR_NAME}_test solution_test.cpp)

# Link with the solution library and kit library
target_link_libraries(${DIR_NAME}_test ${DIR_NAME}_solution kit)

# Add the test to CTest
add_test(NAME ${DIR_NAME}_test COMMAND ${DIR_NAME}_test)