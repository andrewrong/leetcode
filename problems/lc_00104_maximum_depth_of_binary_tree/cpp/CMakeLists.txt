add_executable(lc_00104_maximum_depth_of_binary_tree_test
    solution.cpp
    test.cpp
)

target_link_libraries(lc_00104_maximum_depth_of_binary_tree_test
    PRIVATE
    leetcode_kit_cpp
)