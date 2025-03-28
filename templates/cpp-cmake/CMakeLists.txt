cmake_minimum_required(VERSION 3.15)
project({{.ProjectName}} VERSION 1.0.0 LANGUAGES CXX)

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_CXX_EXTENSIONS OFF)

# Options
option(BUILD_TESTS "Build test suite" ON)

# Find packages
find_package(GTest REQUIRED)

# Add subdirectories
add_subdirectory(src)
if(BUILD_TESTS)
    add_subdirectory(test)
endif()

# Install targets
install(TARGETS {{.ProjectName}}
    RUNTIME DESTINATION bin
    LIBRARY DESTINATION lib
    ARCHIVE DESTINATION lib
) 