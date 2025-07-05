package pokeapi

import "time"

const (
	DefaultTimeout       = 10 * time.Second
	DefaultRetries       = 3
	DefaultRetryInterval = 500 * time.Millisecond
)

// Config contains the settings for the SDK HTTP client.
type Config struct {
	timeout       time.Duration
	retries       int
	RetryInterval time.Duration
}

// NewConfig returns a new Config with default values.
func NewConfig() *Config {
	return &Config{
		timeout:       DefaultTimeout,
		retries:       DefaultRetries,
		RetryInterval: DefaultRetryInterval,
	}
}

// WithTimeout sets the timeout duration for HTTP requests.
func (c *Config) WithTimeout(timeout time.Duration) *Config {
	c.timeout = timeout
	return c
}

// WithRetries sets the number of retry attempts on failure.
func (c *Config) WithRetries(retries int) *Config {
	c.retries = retries
	return c
}

// WithRetryDelay sets the delay between retries.
func (c *Config) WithRetryInterval(retryInterval time.Duration) *Config {
	c.RetryInterval = retryInterval
	return c
}
