package functional_options

import (
	"time"
)

// https://github.com/uptrace/uptrace-go/blob/master/uptrace/config.go

// type Option func(c configurator)

type configurator interface {
	setTimeout(timeout time.Duration)
}

type Client struct {
	timeout time.Duration
}

func (c Client) setTimeout(timeout time.Duration) {
	c.timeout = timeout
}

func NewClient(options ...Option) *Client {
	c := new(Client)
	for _, option := range options {
		option.apply(c)
	}
	return c
}

// clients

type ClusterClient struct {
	timeout time.Duration
}

func (c ClusterClient) setTimeout(timeout time.Duration) {
	c.timeout = timeout
}

func NewClusterClient(options ...Option) *ClusterClient {
	c := new(ClusterClient)
	for _, option := range options {
		option.apply(c)
	}
	return c
}

func (c *Client) setAddr(addr string) {
	c.addr = addr
}

func (c *ClusterClient) setAddr(addr string) {
	c.addrs = append(c.addrs, addr)
}

type config struct{}

type Option interface {
	apply(conf *config)
}

type option func(conf *config)

func (fn option) apply(conf *config) {
	fn(conf)
}

// WithDSN configures a data source name that is used to connect to Uptrace, for example,
// `https://<token>@uptrace.dev/<project_id>`.
//
// The default is to use UPTRACE_DSN environment variable.
func WithDSN(dsn string) Option {
	return option(func(conf *config) {
		conf.dsn = dsn
	})
}
