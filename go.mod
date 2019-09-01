module github.com/hobord/infra2

go 1.12

// exclude github.com/Sirupsen/logrus v1.4.2
replace github.com/Sirupsen/logrus v1.4.2 => github.com/sirupsen/logrus v1.4.2

require (
	github.com/Sirupsen/logrus v1.4.2
	github.com/golang/protobuf v1.3.2
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/google/uuid v1.1.1
	github.com/spf13/viper v1.4.0
	golang.org/x/net v0.0.0-20190522155817-f3200d17e092
	google.golang.org/grpc v1.23.0
)
