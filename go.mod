module github.com/tddey01/luffy

go 1.14

require (
	github.com/Shopify/sarama v1.27.0
	github.com/StackExchange/wmi v0.0.0-20190523213315-cbe66965904d // indirect
	github.com/axgle/mahonia v0.0.0-20180208002826-3358181d7394
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.60.2
	github.com/go-ole/go-ole v1.2.4 // indirect
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.2
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/websocket v1.4.2
	github.com/hpcloud/tail v1.0.0
	github.com/influxdata/influxdb1-client v0.0.0-20200827194710-b269163b24ab
	github.com/jmoiron/sqlx v1.2.0
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.18.0
	github.com/nsqio/go-nsq v1.0.8
	github.com/olivere/elastic/v7 v7.0.19
	github.com/satori/go.uuid v1.2.0
	github.com/shirou/gopsutil v2.20.7+incompatible
	github.com/sirupsen/logrus v1.6.0
	github.com/sony/sonyflake v1.0.0
	go.etcd.io/etcd v3.3.25+incompatible
	go.uber.org/zap v1.15.0
	golang.org/x/net v0.0.0-20200822124328-c89045814202
	google.golang.org/grpc v1.27.0
	google.golang.org/protobuf v1.25.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
