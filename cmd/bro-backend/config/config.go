package config

import (
	"io"

	"github.com/threeal/bro/pkg/utils"
)

var CONFIG_FILENAME = "backend_config.json"

type Config struct {
	ListenAddr string `json:"listen_addr"`
}

func (c *Config) Read() error {
	return utils.ReadConfigFromFile(c, CONFIG_FILENAME)
}

func (c *Config) Write() error {
	return utils.WriteConfigToFile(c, CONFIG_FILENAME)
}

func (c *Config) Init(rd io.Reader) error {
	if c.ListenAddr == "" {
		text, err := utils.Prompt("listen address", ":320", rd)
		if err != nil {
			return err
		}
		if text == "" {
			text = ":320"
		}
		c.ListenAddr = text
	}
	return nil
}
