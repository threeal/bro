package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/fatih/color"
	"github.com/threeal/bro/pkg/utils"
)

var CONFIG_FILENAME = "backend_config.json"

type Config struct {
	ListenAddr *string `json:"listen_addr"`
}

func (c *Config) Read() error {
	return utils.ReadConfigFromFile(c, CONFIG_FILENAME)
}

func (c *Config) Write() error {
	return utils.WriteConfigToFile(c, CONFIG_FILENAME)
}

func (c *Config) Init(rd io.Reader) error {
	if c.ListenAddr == nil {
		reader := bufio.NewReader(rd)
		fmt.Print(color.HiBlackString("question"), " listen address ", color.HiGreenString("(:320)"), ": ")
		text, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		if text == "\n" {
			text = ":320"
		}
		text = strings.TrimSpace(text)
		c.ListenAddr = &text
	}
	return nil
}
