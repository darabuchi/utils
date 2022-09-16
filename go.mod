module github.com/darabuchi/utils

go 1.16

require (
	github.com/AndreKR/multiface v0.0.0-20211114051930-f51f19dee2dc
	github.com/GoogleCloudPlatform/cloudsql-proxy v1.31.2
	github.com/VividCortex/ewma v1.2.0
	github.com/auyer/steganography v1.0.1
	github.com/crystal-construct/analytics v0.0.0-20160309172048-27445d731098
	github.com/darabuchi/log v0.0.0-20220323032131-710e5e1eab84
	github.com/elliotchance/pie v1.39.0
	github.com/garyburd/redigo v1.6.3
	github.com/glebarez/go-sqlite v1.18.1 // indirect
	github.com/glebarez/sqlite v1.4.6
	github.com/go-playground/validator/v10 v10.11.0
	github.com/go-xman/go.emoji v0.1.2
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0
	github.com/gookit/color v1.5.1
	github.com/hackebrot/turtle v0.2.0
	github.com/jchavannes/go-pgp v0.0.0-20200131171414-e5978e6d02b4
	github.com/moznion/go-unicode-east-asian-width v0.0.0-20140622124307-0231aeb79f9b
	github.com/nsqio/nsq v1.2.1
	github.com/onsi/gomega v1.20.0 // indirect
	github.com/petermattis/goid v0.0.0-20220302125637-5f11c28912df // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/shomali11/xredis v0.0.0-20190608143638-0b54a6bbf40b
	github.com/wcharczuk/go-chart/v2 v2.1.0
	github.com/zachomedia/go-bdf v0.0.0-20210522061406-1a147053be95 // indirect
	go.uber.org/atomic v1.9.0
	go.uber.org/multierr v1.8.0 // indirect
	golang.org/x/crypto v0.0.0-20220411220226-7b82a4e95df4
	golang.org/x/image v0.0.0-20220413100746-70e8d0d3baa9
	gopkg.in/yaml.v3 v3.0.1
	gorm.io/driver/clickhouse v0.4.2
	gorm.io/driver/mysql v1.3.6
	gorm.io/driver/postgres v1.3.9
	gorm.io/driver/sqlite v1.3.6
	gorm.io/driver/sqlserver v1.3.2
	gorm.io/gorm v1.23.8
)

replace github.com/nsqio/nsq => ../nsq
