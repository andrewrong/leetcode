# Root CMakeLists.txt

cmake_minimum_required(VERSION 3.14)
project(LeetCode_Solutions_CPP)

# 设置 C++ 标准为 C++17
set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

# --- GoogleTest ---
# 包含 FetchContent 模块，用于从 Git 仓库拉取依赖
include(FetchContent)

# 声明 GoogleTest 的内容
FetchContent_Declare(
  googletest
  GIT_REPOSITORY https://github.com/google/googletest.git
  GIT_TAG         release-1.12.1
)

# 获取 GoogleTest 并使其可用，避免手动添加子目录
FetchContent_MakeAvailable(googletest)
# --- End GoogleTest ---

# 开启测试，并添加 problems 和 kit 目录
enable_testing()
add_subdirectory(problems)
add_subdirectory(kit/cpp)