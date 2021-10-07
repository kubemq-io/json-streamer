package config

import (
	"fmt"
	"strconv"
	"strings"
)

type Config struct {
	Address  string `json:"address"`
	Queue    string `json:"queue"`
	Table    string `json:"table"`
	Interval int    `json:"interval"`
	host     string
	port     int
}

func (c *Config) Host() string {
	return c.host
}

func (c *Config) Port() int {
	return c.port
}

func (c *Config) SetAddress(address string) *Config {
	c.Address = address
	return c
}
func (c *Config) SetInterval(interval int) *Config {
	c.Interval = interval
	return c
}
func (c *Config) SetQueue(Queue string) *Config {
	c.Queue = Queue
	return c
}

func (c *Config) SetTable(Table string) *Config {
	c.Table = Table
	return c
}
func (c *Config) Validate() error {
	if c.Address == "" {
		return fmt.Errorf("missing kubemq address")
	}
	parts := strings.Split(c.Address, ":")
	if len(parts) != 2 {
		return fmt.Errorf("invalid kubemq address format")
	}
	c.host = parts[0]
	port, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return err
	}
	c.port = int(port)
	if c.Queue == "" {
		return fmt.Errorf("missing kubemq queue name")
	}
	if c.Queue == "" {
		return fmt.Errorf("missing sql server target table")
	}
	if c.Interval < 1 {
		return fmt.Errorf("interval must be >=1")
	}
	return nil
}
func NewConfig() *Config {
	return &Config{}
}
