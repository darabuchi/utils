package unit

import (
	"testing"
)

func TestSize(t *testing.T) {
	t.Log(Size2B(1024 * Byte))
	t.Log(Size2B(1024 * bits))

	t.Log(Size2B(10 * GB))
	t.Log(Size2b(10 * Gb))
}

func TestSpeed(t *testing.T) {
	t.Log(Speed2Bs(10 * GB))
	t.Log(Speed2bps(10 * Gb))
}
