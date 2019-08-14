load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "posev.go",
    ],
    importpath = "github.com/antonovvk/posev",
    visibility = ["//visibility:public"],
    deps = [
        "@com_github_mitsuse_matrix_go//dense:go_default_library",
        "@com_github_mitsuse_matrix_go//:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["posev_test.go"],
    importpath = "github.com/antonovvk/posev",
    embed = [":go_default_library"],
)
