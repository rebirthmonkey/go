package tencentcloud

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"time"

	"github.com/go-logr/logr"
)

const (
	SleepDurationOnRetryable = time.Second
	DefaultRegion            = "ap-guangzhou"
	DefaultTimeout           = 30
)

type Option func(*Options)

type Options struct {
	SecretId    string
	SecretKey   string
	Token       string
	Proxy       string
	RootDomain  string
	Endpoint    string
	COSEndpoint string
	Scheme      string
	Region      string
	Timeout     int
	Logger      logr.Logger
	Lang        string
}

func NewOptions(opts ...Option) (options *Options, err error) {
	if err != nil {
		return
	}
	options = &Options{
		Region:  DefaultRegion,
		Timeout: DefaultTimeout,
	}
	for _, opt := range opts {
		opt(options)
	}

	return
}

func (opt Options) Credential() *common.Credential {
	return common.NewCredential(opt.SecretId, opt.SecretKey)
}

func (opt Options) ClientProfile() *profile.ClientProfile {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Proxy = opt.Proxy
	if opt.Endpoint != "" {
		cpf.HttpProfile.Endpoint = opt.Endpoint
	}
	if opt.Scheme != "" {
		cpf.HttpProfile.Scheme = opt.Scheme
	}
	if opt.Lang != "" {
		cpf.Language = opt.Lang
	}
	if opt.Timeout > 0 {
		cpf.HttpProfile.ReqTimeout = opt.Timeout
	}
	if opt.RootDomain != "" {
		cpf.HttpProfile.RootDomain = opt.RootDomain
	}
	return cpf
}

func SecretId(secretId string) Option {
	return func(options *Options) {
		options.SecretId = secretId
	}
}

func SecretKey(secretKey string) Option {
	return func(options *Options) {
		options.SecretKey = secretKey
	}
}
