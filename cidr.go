package utils

import (
	"encoding/binary"
	"fmt"
	"net"
	"sort"
)

func MergeCIDRs(cidrs []string) ([]string, error) {
	if cidrs == nil {
		return nil, nil
	}
	if len(cidrs) == 0 {
		return make([]string, 0), nil
	}

	var networks []*net.IPNet
	for _, cidr := range cidrs {
		_, network, err := net.ParseCIDR(cidr)
		if err != nil {
			return nil, err
		}
		networks = append(networks, network)
	}
	mergedNets, err := MergeIPNets(networks)
	if err != nil {
		return nil, err
	}

	return ipNets(mergedNets).toCIDRs(), nil
}

type ipNets []*net.IPNet

func (nets ipNets) toCIDRs() []string {
	var cidrs []string
	for _, net := range nets {
		cidrs = append(cidrs, net.String())
	}

	return cidrs
}

type cidrBlock4 struct {
	first uint32
	last  uint32
}

type cidrBlock6 struct {
	first uint64
	last  uint64
}

func newBlock4(ip net.IP, mask net.IPMask) *cidrBlock4 {
	var block cidrBlock4

	block.first = ipv4ToUInt32(ip)
	prefix, _ := mask.Size()
	block.last = broadcast4(block.first, uint(prefix))

	return &block
}

func newBlock6(ip net.IP, mask net.IPMask) *cidrBlock6 {
	var block cidrBlock6

	block.first = ipv6ToUInt64(ip)
	prefix, _ := mask.Size()
	block.last = broadcast6(block.first, uint(prefix))

	return &block
}

func broadcast4(addr uint32, prefix uint) uint32 {
	return addr | ^netmask4(prefix)
}

func broadcast6(addr uint64, prefix uint) uint64 {
	return addr | ^netmask6(prefix)
}

func netmask4(prefix uint) uint32 {
	if prefix == 0 {
		return 0
	}
	return ^uint32((1 << (32 - prefix)) - 1)
}

func netmask6(prefix uint) uint64 {
	if prefix == 0 {
		return 0
	}
	return ^uint64((1 << (64 - prefix)) - 1)
}

func ipv4ToUInt32(ip net.IP) uint32 {
	return binary.BigEndian.Uint32(ip)
}

func ipv6ToUInt64(ip net.IP) uint64 {
	return binary.BigEndian.Uint64(ip)
}

func uint32ToIPV4(addr uint32) net.IP {
	ip := make([]byte, net.IPv4len)
	binary.BigEndian.PutUint32(ip, addr)
	return ip
}

func uint64ToIPV6(addr uint64) net.IP {
	ip := make([]byte, net.IPv6len)
	binary.BigEndian.PutUint64(ip, addr)
	return ip
}

type cidrBlock4s []*cidrBlock4

func (c cidrBlock4s) Len() int {
	return len(c)
}

func (c cidrBlock4s) Less(i, j int) bool {
	lhs := c[i]
	rhs := c[j]

	// By last IP in the range.
	if lhs.last < rhs.last {
		return true
	} else if lhs.last > rhs.last {
		return false
	}

	// Then by first IP in the range.
	if lhs.first < rhs.first {
		return true
	} else if lhs.first > rhs.first {
		return false
	}

	return false
}

func (c cidrBlock4s) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

type cidrBlock6s []*cidrBlock6

func (c cidrBlock6s) Len() int {
	return len(c)
}

func (c cidrBlock6s) Less(i, j int) bool {
	lhs := c[i]
	rhs := c[j]

	// By last IP in the range.
	if lhs.last < rhs.last {
		return true
	} else if lhs.last > rhs.last {
		return false
	}

	// Then by first IP in the range.
	if lhs.first < rhs.first {
		return true
	} else if lhs.first > rhs.first {
		return false
	}

	return false
}

func (c cidrBlock6s) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func MergeIPNets(nets []*net.IPNet) ([]*net.IPNet, error) {
	if nets == nil {
		return nil, nil
	}
	if len(nets) == 0 {
		return make([]*net.IPNet, 0), nil
	}

	// Split into IPv4 and IPv6 lists.
	// Merge the list separately and then combine.
	var block4s cidrBlock4s
	var block6s cidrBlock6s
	for _, net := range nets {
		if net.IP.To4() != nil {
			block4s = append(block4s, newBlock4(net.IP.To4(), net.Mask))
		} else if net.IP.To16() != nil {
			block6s = append(block6s, newBlock6(net.IP.To16(), net.Mask))
		}
	}

	var merged []*net.IPNet

	merged4, err := merge4(block4s)
	if err != nil {
		return nil, err
	}

	merged = append(merged, merged4...)

	merged6, err := merge6(block6s)
	if err != nil {
		return nil, err
	}

	merged = append(merged, merged6...)

	return merged, nil
}

func merge4(blocks cidrBlock4s) ([]*net.IPNet, error) {
	sort.Sort(blocks)

	// Coalesce overlapping blocks.
	for i := len(blocks) - 1; i > 0; i-- {
		if blocks[i].first <= blocks[i-1].last+1 {
			blocks[i-1].last = blocks[i].last
			if blocks[i].first < blocks[i-1].first {
				blocks[i-1].first = blocks[i].first
			}
			blocks[i] = nil
		}
	}

	var merged []*net.IPNet
	for _, block := range blocks {
		if block == nil {
			continue
		}

		if err := splitRange4(0, 0, block.first, block.last, &merged); err != nil {
			return nil, err
		}
	}

	return merged, nil
}

func merge6(blocks cidrBlock6s) ([]*net.IPNet, error) {
	sort.Sort(blocks)

	// Coalesce overlapping blocks.
	for i := len(blocks) - 1; i > 0; i-- {
		if blocks[i].first <= blocks[i-1].last+1 {
			blocks[i-1].last = blocks[i].last
			if blocks[i].first < blocks[i-1].first {
				blocks[i-1].first = blocks[i].first
			}
			blocks[i] = nil
		}
	}

	var merged []*net.IPNet
	for _, block := range blocks {
		if block == nil {
			continue
		}

		if err := splitRange6(0, 0, block.first, block.last, &merged); err != nil {
			return nil, err
		}
	}

	return merged, nil
}

// splitRange4 recursively computes the CIDR blocks to cover the range lo to hi.
func splitRange4(addr uint32, prefix uint, lo, hi uint32, cidrs *[]*net.IPNet) error {
	if prefix > 32 {
		return fmt.Errorf("Invalid mask size: %d", prefix)
	}

	bc := broadcast4(addr, prefix)
	if (lo < addr) || (hi > bc) {
		return fmt.Errorf("%d, %d out of range for network %d/%d, broadcast %d", lo, hi, addr, prefix, bc)
	}

	if (lo == addr) && (hi == bc) {
		cidr := net.IPNet{IP: uint32ToIPV4(addr), Mask: net.CIDRMask(int(prefix), 8*net.IPv4len)}
		*cidrs = append(*cidrs, &cidr)
		return nil
	}

	prefix++
	lowerHalf := addr
	upperHalf := setBit4(addr, prefix, 1)
	if hi < upperHalf {
		return splitRange4(lowerHalf, prefix, lo, hi, cidrs)
	} else if lo >= upperHalf {
		return splitRange4(upperHalf, prefix, lo, hi, cidrs)
	} else {
		err := splitRange4(lowerHalf, prefix, lo, broadcast4(lowerHalf, prefix), cidrs)
		if err != nil {
			return err
		}
		return splitRange4(upperHalf, prefix, upperHalf, hi, cidrs)
	}
}

// splitRange4 recursively computes the CIDR blocks to cover the range lo to hi.
func splitRange6(addr uint64, prefix uint, lo, hi uint64, cidrs *[]*net.IPNet) error {
	if prefix > 32 {
		return fmt.Errorf("Invalid mask size: %d", prefix)
	}

	bc := broadcast6(addr, prefix)
	if (lo < addr) || (hi > bc) {
		return fmt.Errorf("%d, %d out of range for network %d/%d, broadcast %d", lo, hi, addr, prefix, bc)
	}

	if (lo == addr) && (hi == bc) {
		cidr := net.IPNet{IP: uint64ToIPV6(addr), Mask: net.CIDRMask(int(prefix), 8*net.IPv4len)}
		*cidrs = append(*cidrs, &cidr)
		return nil
	}

	prefix++
	lowerHalf := addr
	upperHalf := setBit6(addr, prefix, 1)
	if hi < upperHalf {
		return splitRange6(lowerHalf, prefix, lo, hi, cidrs)
	} else if lo >= upperHalf {
		return splitRange6(upperHalf, prefix, lo, hi, cidrs)
	} else {
		err := splitRange6(lowerHalf, prefix, lo, broadcast6(lowerHalf, prefix), cidrs)
		if err != nil {
			return err
		}
		return splitRange6(upperHalf, prefix, upperHalf, hi, cidrs)
	}
}

// setBit sets the specified bit in an address to 0 or 1.
func setBit4(addr uint32, bit uint, val uint) uint32 {
	if bit < 0 {
		panic("negative bit index")
	}

	if val == 0 {
		return addr & ^(1 << (32 - bit))
	} else if val == 1 {
		return addr | (1 << (32 - bit))
	} else {
		panic("set bit is not 0 or 1")
	}
}

// setBit sets the specified bit in an address to 0 or 1.
func setBit6(addr uint64, bit uint, val uint) uint64 {
	if bit < 0 {
		panic("negative bit index")
	}

	if val == 0 {
		return addr & ^(1 << (32 - bit))
	} else if val == 1 {
		return addr | (1 << (32 - bit))
	} else {
		panic("set bit is not 0 or 1")
	}
}
