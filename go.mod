module github.com/hobord/infra2

go 1.12

// exclude github.com/Sirupsen/logrus v1.4.2
replace github.com/Sirupsen/logrus v1.4.2 => github.com/sirupsen/logrus v1.4.2

require (
	github.com/Sirupsen/logrus v1.4.2
	github.com/golang/protobuf v1.3.2
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/google/uuid v1.1.1
	github.com/hobord/infra v0.0.0-20190829102330-cb4041da876a
	github.com/spf13/viper v1.4.0
	google.golang.org/grpc v1.23.0
)
