package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_test",
    "go_library",
)

go_test(
    name = "go_default_test",
    srcs = ["codel_test.go"],
    embed = [":go_default_library"],
    rundir = ".",
    tags = ["automanaged"],
    deps = ["//library/ecode:go_default_library"],
)

go_library(
    name = "go_default_library",
    srcs = ["codel.go"],
    importpath = "wwqdrh/handbook/library/container/queue/aqm",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = ["//library/ecode:go_default_library"],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
