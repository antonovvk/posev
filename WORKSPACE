git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go",
    commit = "63cf58b28feca3690ebc6565d4c4e26b0c80df25",
)
load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains", "go_repository")
go_rules_dependencies()
go_register_toolchains()

go_repository(
    name = "com_github_mitsuse_serial_go",
    tag = "v0.2.0",
    importpath = "github.com/mitsuse/serial-go",
)

go_repository(
    name = "com_github_mitsuse_matrix_go",
    tag = "v0.1.5",
    importpath = "github.com/mitsuse/matrix-go",
)
