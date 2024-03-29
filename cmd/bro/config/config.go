package config

import (
	"io"

	"github.com/threeal/bro/pkg/utils"
)

var CONFIG_FILENAME = "config.json"

type Config struct {
	BackendAddr string `json:"backend_addr"`
}

func (c *Config) Read() error {
	return utils.ReadConfigFromFile(c, CONFIG_FILENAME)
}

func (c *Config) Write() error {
	return utils.WriteConfigToFile(c, CONFIG_FILENAME)
}

func (c *Config) Init(rd io.Reader) error {
	if c.BackendAddr == "" {
		text, err := utils.Prompt("backend address", "localhost:13120", rd)
		if err != nil {
			return err
		}
		if text == "" {
			text = "localhost:13120"
		}
		c.BackendAddr = text
	}
	return nil
}
