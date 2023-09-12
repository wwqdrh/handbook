set_languages("c++17")

add_requires("gtest")

target("test_os")
    set_kind("binary")
    add_files("tests/test_os/file_test.cpp", "src/os/file.cpp")
    add_includedirs("include")
    add_packages("gtest")
