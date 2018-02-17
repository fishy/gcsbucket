workspace(name = "com_github_fishy_gcsbucket")

http_archive(
    name = "io_bazel_rules_go",
    url = "https://github.com/bazelbuild/rules_go/releases/download/0.9.0/rules_go-0.9.0.tar.gz",
    sha256 = "4d8d6244320dd751590f9100cf39fd7a4b75cd901e1f3ffdfd6f048328883695",
)
http_archive(
    name = "bazel_gazelle",
    url = "https://github.com/bazelbuild/bazel-gazelle/releases/download/0.9/bazel-gazelle-0.9.tar.gz",
    sha256 = "0103991d994db55b3b5d7b06336f8ae355739635e0c2379dea16b8213ea5a223",
)

load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains", "go_repository")
go_rules_dependencies()
go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
gazelle_dependencies()

go_repository(
    name = "com_github_fishy_fsdb",
    importpath = "github.com/fishy/fsdb",
    commit = "5527ded0137100d26d3b85b6473aab591f98959b",
)

go_repository(
    name = "com_google_cloud_go",
    importpath = "cloud.google.com/go",
    commit = "767c40d6a2e058483c25fa193e963a22da17236d",
)

go_repository(
    name = "org_golang_google_api",
    importpath = "google.golang.org/api",
    commit = "c7a403bb5fe12fe4644bb956edd5cd677626120a",
)

go_repository(
    name = "com_github_googleapis_gax_go",
    importpath = "github.com/googleapis/gax-go",
    commit = "317e0006254c44a0ac427cc52a0e083ff0b9622f",
)

go_repository(
    name = "org_golang_x_oauth2",
    importpath = "golang.org/x/oauth2",
    commit = "543e37812f10c46c622c9575afd7ad22f22a12ba",
)
