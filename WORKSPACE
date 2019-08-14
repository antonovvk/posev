load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "f04d2373bcaf8aa09bccb08a98a57e721306c8f6043a2a0ee610fd6853dcde3d",
    urls = [
        "https://storage.googleapis.com/bazel-mirror/github.com/bazelbuild/rules_go/releases/download/0.18.6/rules_go-0.18.6.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/0.18.6/rules_go-0.18.6.tar.gz",
    ],
)

git_repository(
    name = "bazel_gazelle",
    commit = "e443c54b396a236e0d3823f46c6a931e1c9939f2",
    remote = "https://github.com/bazelbuild/bazel-gazelle",
)

load("@bazel_gazelle//:deps.bzl", "go_repository")

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()
go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

gazelle_dependencies()

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
