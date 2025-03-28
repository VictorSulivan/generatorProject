add_executable({{.ProjectName}}_tests
    main_test.cpp
)

target_include_directories({{.ProjectName}}_tests
    PUBLIC
        ${CMAKE_SOURCE_DIR}/include
)

target_link_libraries({{.ProjectName}}_tests
    PRIVATE
        GTest::GTest
        GTest::Main
)

add_test(NAME {{.ProjectName}}_tests COMMAND {{.ProjectName}}_tests) 