load(
    "@io_bazel_rules_go//proto:def.bzl",
    "go_proto_library",
)

package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

proto_library(
    name = "v2_proto",
    srcs = ["Room.proto"],
    tags = ["automanaged"],
)

go_proto_library(
    name = "v2_go_proto",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    importpath = "wwqdrh/handbook/library/net/rpc/liverpc/testdata/v2",
    proto = ":v2_proto",
    tags = ["automanaged"],
)

go_library(
    name = "go_default_library",
    srcs = ["Room.liverpc.go"],
    embed = [":v2_go_proto"],
    importpath = "wwqdrh/handbook/library/net/rpc/liverpc/testdata/v2",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/net/rpc/liverpc:go_default_library",
        "@com_github_gogo_protobuf//proto:go_default_library",
        "@com_github_golang_protobuf//proto:go_default_library",
    ],
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
