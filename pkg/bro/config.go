package bro

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/fatih/color"
	"github.com/threeal/bro/pkg/utils"
)

var CONFIG_FILENAME = "config.json"

type Config struct {
	BackendAddr *string `json:"backend_addr"`
}

func (c *Config) Read() error {
	return utils.ReadConfigFromFile(c, CONFIG_FILENAME)
}

func (c *Config) Write() error {
	return utils.WriteConfigToFile(c, CONFIG_FILENAME)
}

func (c *Config) Init(rd io.Reader) error {
	if c.BackendAddr == nil {
		reader := bufio.NewReader(rd)
		fmt.Print(color.HiBlackString("question"), " backend address ", color.HiGreenString("(localhost:320)"), ": ")
		text, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		if text == "\n" {
			text = "localhost:320"
		}
		text = strings.TrimSpace(text)
		c.BackendAddr = &text
	}
	return nil
}
