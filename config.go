package pokeapi

import "time"

const (
	DefaultTimeout       = 10 * time.Second
	DefaultRetries       = 3
	DefaultRetryInterval = 500 * time.Millisecond
)

type Config struct {
	timeout       time.Duration
	retries       int
	RetryInterval time.Duration
}

func NewConfig() *Config {
	return &Config{
		timeout:       DefaultTimeout,
		retries:       DefaultRetries,
		RetryInterval: DefaultRetryInterval,
	}
}

func (c *Config) WithTimeout(timeout time.Duration) *Config {
	c.timeout = timeout
	return c
}

func (c *Config) WithRetries(retries int) *Config {
	c.retries = retries
	return c
}

func (c *Config) WithRetryInterval(retryInterval time.Duration) *Config {
	c.RetryInterval = retryInterval
	return c
}
