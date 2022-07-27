module github.com/darabuchi/utils

go 1.16

require (
	github.com/AndreKR/multiface v0.0.0-20211114051930-f51f19dee2dc
	github.com/VividCortex/ewma v1.2.0
	github.com/auyer/steganography v1.0.1
	github.com/crystal-construct/analytics v0.0.0-20160309172048-27445d731098
	github.com/darabuchi/log v0.0.0-20220323032131-710e5e1eab84
	github.com/elliotchance/pie v1.39.0
	github.com/garyburd/redigo v1.6.3
	github.com/go-playground/validator/v10 v10.11.0
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0
	github.com/jchavannes/go-pgp v0.0.0-20200131171414-e5978e6d02b4
	github.com/nsqio/nsq v1.2.1
	github.com/petermattis/goid v0.0.0-20220302125637-5f11c28912df // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/shomali11/xredis v0.0.0-20190608143638-0b54a6bbf40b
	github.com/wcharczuk/go-chart/v2 v2.1.0
	github.com/zachomedia/go-bdf v0.0.0-20210522061406-1a147053be95 // indirect
	go.uber.org/atomic v1.9.0
	go.uber.org/multierr v1.8.0 // indirect
	golang.org/x/crypto v0.0.0-20220321153916-2c7772ba3064
	golang.org/x/image v0.0.0-20220413100746-70e8d0d3baa9
	golang.org/x/sys v0.0.0-20220319134239-a9b59b0215f8 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)

replace github.com/nsqio/nsq => ../nsq
