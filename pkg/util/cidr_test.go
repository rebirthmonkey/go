package util

import (
	"testing"

	"github.com/onsi/gomega"
)

func TestCIDRContains(t *testing.T) {
	gomega.RegisterTestingT(t)
	gomega.Expect(CIDRContains("10.0.0.0/8", "10.0.0.0/23")).To(gomega.BeTrue())
	gomega.Expect(CIDRContains("10.0.0.0/16", "10.0.0.0/8")).To(gomega.BeFalse())
	gomega.Expect(CIDRContains("172.21.0.0/16", "10.0.0.0/8")).To(gomega.BeFalse())
	gomega.Expect(CIDRContains("10.0.0.0/8", "172.21.0.0/16")).To(gomega.BeFalse())
}
