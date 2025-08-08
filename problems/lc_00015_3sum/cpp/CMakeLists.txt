# CMakeLists.txt for lc_00015_3sum problem

add_library(lc_00015_3sum solution.cpp)

target_include_directories(lc_00015_3sum PUBLIC 
    ${CMAKE_CURRENT_SOURCE_DIR}
    ${PROJECT_SOURCE_DIR}/kit/cpp
)

add_executable(lc_00015_3sum_test solution_test.cpp)

target_link_libraries(lc_00015_3sum_test 
    lc_00015_3sum
    gtest 
    gtest_main
)

target_include_directories(lc_00015_3sum_test PUBLIC 
    ${CMAKE_CURRENT_SOURCE_DIR}
    ${PROJECT_SOURCE_DIR}/kit/cpp
)

add_test(NAME lc_00015_3sum_test COMMAND lc_00015_3sum_test)