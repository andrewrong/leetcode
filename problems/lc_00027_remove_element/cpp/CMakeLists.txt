# problems/lc_00027_remove_element/cpp/CMakeLists.txt

# 项目名称，例如 lc_00027_remove_element
project(lc_00027_remove_element)

# 添加可执行文件，源文件包括解法和测试
add_executable(
    ${PROJECT_NAME}_test
    solution_test.cpp
)

# 链接 GoogleTest 库
target_link_libraries(
    ${PROJECT_NAME}_test
    gtest_main # 使用 gtest_main 会自动生成 main 函数
)

# 将可执行文件添加到 CTest 测试集中
include(GoogleTest)
gtest_add_tests(TARGET ${PROJECT_NAME}_test)