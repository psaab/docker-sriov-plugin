module github.com/psaab/docker-sriov-plugin

go 1.26.2

require (
	github.com/Mellanox/rdmamap v1.2.0
	github.com/docker/docker v25.0.6+incompatible
	github.com/docker/go-plugins-helpers v0.0.0-20240701071450-45e2431495c8
	github.com/k8snetworkplumbingwg/sriovnet v1.2.0
	github.com/urfave/cli v1.22.17
	github.com/vishvananda/netlink v1.3.1
)

require (
	github.com/Microsoft/go-winio v0.6.2 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/containerd/log v0.1.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.7 // indirect
	github.com/distribution/reference v0.6.0 // indirect
	github.com/docker/go-connections v0.5.0 // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/go-logr/logr v1.4.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/moby/term v0.5.2 // indirect
	github.com/morikuni/aec v1.1.0 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.1.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/spf13/afero v1.9.4 // indirect
	github.com/vishvananda/netns v0.0.5 // indirect
	go.opentelemetry.io/auto/sdk v1.2.1 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.68.0 // indirect
	go.opentelemetry.io/otel v1.43.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.43.0 // indirect
	go.opentelemetry.io/otel/metric v1.43.0 // indirect
	go.opentelemetry.io/otel/trace v1.43.0 // indirect
	golang.org/x/sys v0.42.0 // indirect
	golang.org/x/text v0.35.0 // indirect
	golang.org/x/time v0.15.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20260401024825-9d38bb4040a9 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260401024825-9d38bb4040a9 // indirect
	gotest.tools/v3 v3.5.2 // indirect
)

exclude github.com/docker/docker/api v1.54.2

exclude github.com/docker/docker/client v0.4.1

replace github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.9.3
