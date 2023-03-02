package utils_test

import (
	"github.com/darabuchi/utils"
	"testing"

	"github.com/darabuchi/log"
)

func TestIsLocalIp(t *testing.T) {
	log.Info(utils.IsLocalIp("127.0.0.1"))
}

func TestIpInt(t *testing.T) {
	ipStr := "2400:8901::f03c:93ff:fe78:cac7"
	t.Log(utils.Int2Ip(utils.Ip2Int(ipStr)).String() == ipStr)

	t.Log(utils.Ip2Int("60.73.56.108"))
	t.Log(utils.Int2Ip(1011431532))
}
