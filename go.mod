module github.com/darabuchi/utils

go 1.16

require (
	cloud.google.com/go/compute v1.10.0 // indirect
	github.com/AndreKR/multiface v0.0.0-20211114051930-f51f19dee2dc
	github.com/GoogleCloudPlatform/cloudsql-proxy v1.32.0
	github.com/VividCortex/ewma v1.2.0
	github.com/auyer/steganography v1.0.1
	github.com/bytedance/sonic v1.5.0
	github.com/chenzhuoyu/base64x v0.0.0-20220526154910-8bf9453eb81a // indirect
	github.com/crystal-construct/analytics v0.0.0-20160309172048-27445d731098
	github.com/darabuchi/log v0.0.0-20220923020052-f1a3dc411901
	github.com/elliotchance/pie v1.39.0
	github.com/garyburd/redigo v1.6.4
	github.com/glebarez/sqlite v1.4.7
	github.com/go-playground/validator/v10 v10.11.1
	github.com/go-xman/go.emoji v0.1.2
	github.com/golang-sql/civil v0.0.0-20220223132316-b832511892a9 // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0
	github.com/gookit/color v1.5.2
	github.com/jackc/pgx/v4 v4.17.2 // indirect
	github.com/jchavannes/go-pgp v0.0.0-20200131171414-e5978e6d02b4
	github.com/klauspost/cpuid/v2 v2.1.1 // indirect
	github.com/mcuadros/go-defaults v1.2.0
	github.com/moznion/go-unicode-east-asian-width v0.0.0-20140622124307-0231aeb79f9b
	github.com/nsqio/nsq v1.2.1
	github.com/onsi/gomega v1.20.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/shomali11/xredis v0.0.0-20190608143638-0b54a6bbf40b
	github.com/shopspring/decimal v1.3.1 // indirect
	github.com/wcharczuk/go-chart/v2 v2.1.0
	github.com/xo/terminfo v0.0.0-20220910002029-abceb7e1c41e // indirect
	github.com/zachomedia/go-bdf v0.0.0-20210522061406-1a147053be95 // indirect
	go.etcd.io/etcd/api/v3 v3.5.5
	go.etcd.io/etcd/client/v3 v3.5.5
	go.uber.org/atomic v1.10.0
	go.uber.org/zap v1.23.0
	golang.org/x/arch v0.0.0-20220927172834-6a65923eb742 // indirect
	golang.org/x/crypto v0.0.0-20220926161630-eccd6366d1be
	golang.org/x/image v0.0.0-20220902085622-e7cb96979f69
	golang.org/x/sys v0.0.0-20220928140112-f11e5e49a4ec // indirect
	gopkg.in/yaml.v3 v3.0.1
	gorm.io/driver/mysql v1.3.6
	gorm.io/driver/postgres v1.3.9
	gorm.io/driver/sqlite v1.3.6
	gorm.io/driver/sqlserver v1.3.2
	gorm.io/gorm v1.23.9
	modernc.org/libc v1.19.0 // indirect
)

replace github.com/nsqio/nsq => github.com/darabuchi/nsq v1.2.2-0.20220625061629-f6fd2f159a76
