module github.com/shohi/gocode

go 1.12

require (
	cloud.google.com/go v0.39.0
	github.com/BurntSushi/toml v0.3.1
	github.com/Microsoft/go-winio v0.4.12 // indirect
	github.com/adlio/trello v1.0.0
	github.com/armon/go-metrics v0.0.0-20190430140413-ec5e00d3c878 // indirect
	github.com/atrox/homedir v1.0.0
	github.com/axgle/mahonia v0.0.0-20180208002826-3358181d7394
	github.com/bramvdbogaerde/go-scp v0.0.0-20190409174733-583e65a51240
	github.com/djimenez/iconv-go v0.0.0-20160305225143-8960e66bd3da
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v1.13.1
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/go-sql-driver/mysql v1.4.1
	github.com/gogo/protobuf v1.2.1 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.3.2
	github.com/goph/emperror v0.17.1
	github.com/gorilla/mux v1.7.1
	github.com/hashicorp/go-msgpack v0.5.5 // indirect
	github.com/hashicorp/raft v1.0.1 // indirect
	github.com/jinzhu/copier v0.0.0-20180308034124-7e38e58719c3
	github.com/jinzhu/now v1.0.0
	github.com/klauspost/cpuid v1.2.1
	github.com/lib/pq v1.1.1 // indirect
	github.com/mitchellh/go-ps v0.0.0-20190716172923-621e5597135b
	github.com/nats-io/gnatsd v1.4.1 // indirect
	github.com/nats-io/go-nats v1.7.2
	github.com/nats-io/go-nats-streaming v0.4.4
	github.com/nats-io/nats-server v1.4.1 // indirect
	github.com/nats-io/nats-streaming-server v0.14.1 // indirect
	github.com/nats-io/nkeys v0.0.2
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/opencontainers/go-digest v1.0.0-rc1 // indirect
	github.com/otiai10/gosseract v2.2.1+incompatible
	github.com/otiai10/mint v1.2.4 // indirect
	github.com/pelletier/go-toml v1.4.0
	github.com/pkg/errors v0.8.1
	github.com/prometheus/procfs v0.0.0-20190507164030-5867b95ac084 // indirect
	github.com/sirupsen/logrus v1.4.1
	github.com/spf13/viper v1.3.2
	github.com/stretchr/testify v1.3.0
	github.com/tealeg/xlsx v1.0.3
	github.com/valyala/fasthttp v1.2.0
	go.etcd.io/bbolt v1.3.2 // indirect
	go.uber.org/atomic v1.4.0 // indirect
	go.uber.org/goleak v0.10.0
	go.uber.org/multierr v1.1.0 // indirect
	go.uber.org/zap v1.10.0
	golang.org/x/crypto v0.0.0-20190513172903-22d7a77e9e5f
	golang.org/x/net v0.0.0-20190514140710-3ec191127204
	golang.org/x/text v0.3.2
	google.golang.org/grpc v1.23.1
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.23.1
