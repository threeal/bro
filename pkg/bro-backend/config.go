package brobackend

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/threeal/bro/pkg/utils"
)

type Config struct {
	ListenAddr *string `json:"listen_addr"`
}

func (c *Config) Read() error {
	return utils.ReadConfigFromFile(c, "backend_config.json")
}

func (c *Config) Write() error {
	return utils.WriteConfigToFile(c, "backend_config.json")
}

func (c *Config) Init() error {
	if c.ListenAddr == nil {
		reader := bufio.NewReader(os.Stdin)
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
