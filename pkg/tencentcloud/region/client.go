package region

import (
	"github.com/go-logr/logr"
	"github.com/rebirthmonkey/go/pkg/tencentcloud"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	v20170312cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

type regionClient struct {
	options *tencentcloud.Options
	client  *v20170312cvm.Client
}

// NewClient generate a client instance
func NewClient(opts ...tencentcloud.Option) (client Interface, err error) {
	var options *tencentcloud.Options
	options, err = tencentcloud.NewOptions(opts...)
	if err != nil {
		return
	}
	var cli *v20170312cvm.Client
	cred := common.NewCredential(options.SecretId, options.SecretKey)
	if options.Token != "" {
		cred.Token = options.Token
	}
	cpf := options.ClientProfile()
	cli, err = v20170312cvm.NewClient(cred, options.Region, cpf)
	if err != nil {
		return
	}
	client = &regionClient{
		options: options,
		client:  cli,
	}
	return
}

// Client return the cvm instance
func (c *regionClient) Client() *v20170312cvm.Client {
	return c.client
}

// Logger return the log instance
func (c *regionClient) Logger() logr.Logger {
	return c.options.Logger
}

func (c *regionClient) ListRegions() (regions []*v20170312cvm.RegionInfo, err error) {
	req := v20170312cvm.NewDescribeRegionsRequest()
	var resp *v20170312cvm.DescribeRegionsResponse
	if resp, err = c.Client().DescribeRegions(req); err != nil {
		return
	}
	regions = resp.Response.RegionSet
	return
}
