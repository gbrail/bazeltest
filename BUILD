load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_prefix")

go_prefix("github.com/gbrail/bazeltest")

go_binary(
  name = "bazeltest",
  srcs = [ "main.go" ],
  deps = [
    "@com_github_spf13_pflag//:go_default_library",
    "@com_github_spf13_viper//:go_default_library",
  ],
)
