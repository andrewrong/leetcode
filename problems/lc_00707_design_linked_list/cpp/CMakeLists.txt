add_executable(lc_00707_design_linked_list_test
    list_node_test.cpp
)

target_link_libraries(lc_00707_design_linked_list_test
    gtest_main
)

include(GoogleTest)
gtest_discover_tests(lc_00707_design_linked_list_test)