# problems/CMakeLists.txt

# 遍历 problems 目录下的所有子目录
file(GLOB problem_dirs LIST_DIRECTORIES true *)

foreach(problem_dir ${problem_dirs})
  # 检查每个子目录中是否存在 CMakeLists.txt 文件
  if(EXISTS "${problem_dir}/cpp/CMakeLists.txt")
    # 如果存在，则将其添加到构建系统中
    add_subdirectory(${problem_dir}/cpp)
  endif()
endforeach()