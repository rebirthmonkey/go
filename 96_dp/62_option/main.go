package main

import (
	"fmt"
	"time"
)

type Connection struct {
	addr    string
	timeout time.Duration
	cache   bool
}

const (
	defaultTimeout = 10
	defaultCaching = false
)

type options struct {
	timeout time.Duration
	caching bool
}

// Option overrides behavior of Connect.
type Option interface {
	apply(*options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) {
	f(o)
}

func WithTimeout(t time.Duration) Option {
	return optionFunc(func(o *options) {
		o.timeout = t
	})
}

func WithCaching(cache bool) Option {
	return optionFunc(func(o *options) {
		o.caching = cache
	})
}

// NewConnect creates a connection.
func NewConnect(addr string, opts ...Option) (*Connection, error) {
	options := options{
		timeout: defaultTimeout,
		caching: defaultCaching,
	}

	for _, o := range opts {
		o.apply(&options)
	}

	return &Connection{
		addr:    addr,
		cache:   options.caching,
		timeout: options.timeout,
	}, nil
}

func main() {
	// 不使用任何参数
	fmt.Println(NewConnect("127.0.0.1"))
	// 选择性启用某些选项
	fmt.Println(NewConnect("127.0.0.1", WithTimeout(8)))
	fmt.Println(NewConnect("127.0.0.1", WithTimeout(8), WithCaching(true)))
}
