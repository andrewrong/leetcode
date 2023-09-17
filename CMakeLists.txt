cmake_minimum_required(VERSION 3.12)
project(select)

set(CMAKE_CXX_STANDARD 14)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

include(FetchContent)
FetchContent_Declare(
  googletest
  # Specify the commit you depend on and update it regularly.
  URL https://github.com/google/googletest/archive/5376968f6948923e2411081fd9372e71a59d8e77.zip
)
FetchContent_MakeAvailable(googletest)

# 添加子目录，每个子目录对应一个要编译的模块
add_subdirectory(common)
add_subdirectory(test)

include_directories(${PROJECT_SOURCE_DIR})
include_directories(${PROJECT_SOURCE_DIR}/common)
include_directories(${PROJECT_SOURCE_DIR}/index)
include_directories(${PROJECT_SOURCE_DIR}/_deps/googletest-src/googletest/include)

set(SOURCES
    main.cpp
    common/utils.cpp 
)

# 编译目标
add_executable(select ${SOURCES})
# target_link_libraries(select gtest_main)
# add_test(NAME select_test COMMAND select)

# 添加测试文件的源文件
set(TEST_SOURCES
    test/test.cpp
    test/utils.cpp
    common/utils.cpp
)

# 创建测试目标
add_executable(${PROJECT_NAME}_test ${TEST_SOURCES})
target_link_libraries(${PROJECT_NAME}_test gtest_main)

# 添加测试
enable_testing()
add_test(NAME MyTest COMMAND ${PROJECT_NAME}_test)

