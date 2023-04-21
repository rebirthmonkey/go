package region

import (
	v20170312cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

type Interface interface {
	ListRegions() (regions []*v20170312cvm.RegionInfo, err error)
}
