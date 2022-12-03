package util

import (
	"fmt"
	"net"
)

func CIDRContains(parent, sub string) (bool, error) {
	pip, pnet, err := net.ParseCIDR(sub)
	if err != nil {
		return false, err
	}

	_, snet, err := net.ParseCIDR(parent)
	if err != nil {
		return false, err
	}

	pmsksize, _ := pnet.Mask.Size()
	smsksize, _ := snet.Mask.Size()

	return snet.Contains(pip) && pmsksize >= smsksize, nil
}

// CIDRContainsIp check cidr contains indicated ips
func CIDRContainsIp(cidr string, ips ...string) error {
	_, snet, err := net.ParseCIDR(cidr)
	if err != nil {
		return err
	}
	for _, v := range ips {
		ip := net.ParseIP(v)
		ip.DefaultMask()
		if !snet.Contains(ip) {
			return fmt.Errorf("cidr %s not contains current Ip %s", cidr, v)
		}
	}
	return nil
}

func CIDRSContainsIp(cidrs []string, ip string) bool {
	var contains bool
	for _, cidr := range cidrs {
		if CIDRContainsIp(cidr, ip) == nil || cidr == ip {
			contains = true
			break
		}
	}
	return contains
}
