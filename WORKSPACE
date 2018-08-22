load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    urls = ["https://github.com/bazelbuild/rules_go/releases/download/0.15.0/rules_go-0.15.0.tar.gz"],
    sha256 = "56d946edecb9879aed8dff411eb7a901f687e242da4fa95c81ca08938dd23bb4",
)

http_archive(
    name = "bazel_gazelle",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.14.0/bazel-gazelle-0.14.0.tar.gz"],
    sha256 = "c0a5739d12c6d05b6c1ad56f2200cb0b57c5a70e03ebd2f7b87ce88cabf09c7b",
)

load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

go_repository(
    name = "com_github_influxdata_influxdb",
    commit = "6150bc1eea14f9c9bd7f5b6f2f4d88597576cb32",
    importpath = "github.com/influxdata/influxdb",
)

go_repository(
    name = "com_github_spf13_cobra",
    commit = "6fd8e29b07d8242ebe2888060fede5766e240c25",
    importpath = "github.com/spf13/cobra",
)

go_repository(
    name = "com_github_spf13_viper",
    commit = "907c19d40d9a6c9bb55f040ff4ae45271a4754b9",
    importpath = "github.com/spf13/viper",
)

go_repository(
    name = "com_github_stretchr_testify",
    commit = "f35b8ab0b5a2cef36673838d662e249dd9c94686",
    importpath = "github.com/stretchr/testify",
)

go_repository(
    name = "com_github_magiconair_properties",
    commit = "c2353362d570a7bfa228149c62842019201cfb71",
    importpath = "github.com/magiconair/properties",
)

go_repository(
    name = "com_github_hashicorp_hcl",
    commit = "ef8a98b0bbce4a65b5aa4c368430a80ddc533168",
    importpath = "github.com/hashicorp/hcl",
)

go_repository(
    name = "com_github_spf13_pflag",
    commit = "d929dcbb10863323c436af3cf76cb16a6dfc9b29",
    importpath = "github.com/spf13/pflag",
)

go_repository(
    name = "in_gopkg_yaml_v2",
    commit = "5420a8b6744d3b0345ab293f6fcba19c978f1183",
    importpath = "gopkg.in/yaml.v2",
)

go_repository(
    name = "com_github_mitchellh_mapstructure",
    commit = "f15292f7a699fcc1a38a80977f80a046874ba8ac",
    importpath = "github.com/mitchellh/mapstructure",
)

go_repository(
    name = "com_github_spf13_afero",
    commit = "787d034dfe70e44075ccc060d346146ef53270ad",
    importpath = "github.com/spf13/afero",
)

go_repository(
    name = "com_github_spf13_jwalterweatherman",
    commit = "14d3d4c518341bea657dd8a226f5121c0ff8c9f2",
    importpath = "github.com/spf13/jwalterweatherman",
)

go_repository(
    name = "com_github_pelletier_go_toml",
    commit = "c2dbbc24a97911339e01bda0b8cabdbd8f13b602",
    importpath = "github.com/pelletier/go-toml",
)

go_repository(
    name = "com_github_spf13_cast",
    commit = "8965335b8c7107321228e3e3702cab9832751bac",
    importpath = "github.com/spf13/cast",
)

go_repository(
    name = "com_github_fsnotify_fsnotify",
    commit = "c2828203cd70a50dcccfb2761f8b1f8ceef9a8e9",
    importpath = "github.com/fsnotify/fsnotify",
)
