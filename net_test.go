package utils

import (
	"testing"

	"github.com/darabuchi/log"
)

func TestIsLocalIp(t *testing.T) {
	log.Info(IsLocalIp("127.0.0.1"))
}

func TestIpInt(t *testing.T) {
	ipStr := "2400:8901::f03c:93ff:fe78:cac7"
	t.Log(Int2Ip(Ip2Int(ipStr)).String() == ipStr)

	t.Log(Ip2Int("60.73.56.108"))
	t.Log(Int2Ip(1011431532))
}
