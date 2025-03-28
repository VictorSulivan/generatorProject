add_executable({{.ProjectName}}
    main.cpp
)

target_include_directories({{.ProjectName}}
    PUBLIC
        ${CMAKE_SOURCE_DIR}/include
)

target_link_libraries({{.ProjectName}}
    PRIVATE
        GTest::GTest
        GTest::Main
) 