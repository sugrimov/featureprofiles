# TODO(rahulsy): describe this package.
load("//ops/netops/lab/wbb/testbeds:testbeds.bzl", "featureprofiles_test_suite")

package(
    default_applicable_licenses = ["//third_party/openconfig/featureprofiles:license"],
    default_visibility = ["//third_party/openconfig/featureprofiles:internal"],
)

featureprofiles_test_suite(
    name = "aggregate_all_not_forwarding_viable_test",
    srcs = ["aggregate_all_not_forwarding_viable_test.go"],
    data = [
        "metadata.textproto",
        "//ops/netops/lab/kne/guitar/functional_tests:topologies",
        "//third_party/openconfig/featureprofiles/topologies/kne:topologies",
    ],
    glaze_kind = "go_test",
    tags = ["external"],
    testbed = "//ops/netops/lab/wbb/testbeds:testbed_dut_ate_8links_100g_fr.textproto",
    deps = [
        "//third_party/golang/ygot/ygot",
        "//third_party/open_traffic_generator/gosnappi",
        "//third_party/openconfig/featureprofiles/internal/attrs",
        "//third_party/openconfig/featureprofiles/internal/deviations",
        "//third_party/openconfig/featureprofiles/internal/fptest",
        "//third_party/openconfig/featureprofiles/internal/gribi",
        "//third_party/openconfig/featureprofiles/internal/otgutils",
        "//third_party/openconfig/gribigo/fluent",
        "//third_party/openconfig/ondatra",
        "//third_party/openconfig/ondatra/gnmi",
        "//third_party/openconfig/ondatra/gnmi/oc",
        "//third_party/openconfig/ondatra/gnmi/oc:ocpath",
        "//third_party/openconfig/ondatra/netutil",
        "//third_party/openconfig/ygnmi/ygnmi",
    ],
)
