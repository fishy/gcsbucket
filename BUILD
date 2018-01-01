load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")
load("@bazel_gazelle//:def.bzl", "gazelle")

gazelle(
    name = "gazelle",
    prefix = "github.com/fishy/gcsbucket",
)

go_library(
    name = "go_default_library",
    srcs = ["gcs.go"],
    importpath = "github.com/fishy/gcsbucket",
    visibility = ["//visibility:public"],
    deps = [
        "@com_github_fishy_fsdb//bucket:go_default_library",
        "@com_google_cloud_go//storage:go_default_library",
    ],
)

go_test(
    name = "go_default_xtest",
    srcs = ["gcs_test.go"],
    importpath = "github.com/fishy/gcsbucket_test",
    deps = [":go_default_library"],
)
