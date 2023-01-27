package bro

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/threeal/bro/pkg/utils"
)

type Config struct {
	BackendAddr *string `json:"backend_addr"`
}

func (c *Config) Read() error {
	return utils.ReadConfigFromFile(c, "config.json")
}

func (c *Config) Write() error {
	return utils.WriteConfigToFile(c, "config.json")
}

func (c *Config) Init() error {
	if c.BackendAddr == nil {
		reader := bufio.NewReader(os.Stdin)
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
