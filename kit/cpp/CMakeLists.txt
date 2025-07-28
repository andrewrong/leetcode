# kit/cpp/CMakeLists.txt

# 创建一个库，包含所有的工具类头文件
add_library(kit_cpp INTERFACE)

# 设置库的包含目录
target_include_directories(kit_cpp INTERFACE ${CMAKE_CURRENT_SOURCE_DIR} ${CMAKE_CURRENT_SOURCE_DIR}/print/pfr/include)


# 如果需要创建测试可执行文件
if(BUILD_TESTING)
  # 添加测试可执行文件
  add_executable(kit_cpp_test
    binary_search_test.cpp
  )
  
  # 链接 GoogleTest 和我们的库
  target_link_libraries(kit_cpp_test
    kit_cpp
    gtest_main
  )
  
  # 将测试添加到 CTest
  include(GoogleTest)
  gtest_add_tests(TARGET kit_cpp_test)
endif()