package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/fatih/color"
)

type Config interface {
	Read() error
	Write() error
	init() error
}

type BackendConfig struct {
	ListenAddr *string `json:"listen_addr"`
}

type ClientConfig struct {
	BackendAddr *string `json:"backend_addr"`
}

func getHomeDir() (*string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Printf("%v failed to determine home directory: %v", color.RedString("ERROR:"), err)
	}
	return &homeDir, err
}

func getConfigDir() (string, error) {
	homeDir, err := getHomeDir()
	if err != nil {
		return "", err
	}
	configDir := filepath.Join(*homeDir, ".bro")
	return configDir, err
}

func (c *BackendConfig) Read() error {
	return readConfigFromFile(c, "backend_config.json")
}

func (c *ClientConfig) Read() error {
	return readConfigFromFile(c, "config.json")
}

func (c *BackendConfig) Write() error {
	return writeConfigToFile(c, "backend_config.json")
}

func (c *ClientConfig) Write() error {
	return writeConfigToFile(c, "config.json")
}

func (c *BackendConfig) init() error {
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
		c.ListenAddr = &text
	}
	return nil
}

func (c *ClientConfig) init() error {
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
		c.BackendAddr = &text
	}
	return nil
}

func NewBackendConfig() *BackendConfig {
	return &BackendConfig{ListenAddr: nil}
}

func NewClientConfig() *ClientConfig {
	return &ClientConfig{BackendAddr: nil}
}

func InitializeBackendConfig() *BackendConfig {
	backendConfig := NewBackendConfig()
	initializeConfig(backendConfig)
	return backendConfig
}

func InitializeClientConfig() *ClientConfig {
	clientConfig := NewClientConfig()
	initializeConfig(clientConfig)
	return clientConfig
}

func initializeConfig(c Config) {
	if err := c.Read(); err != nil {
		if err = c.Write(); err != nil {
			log.Fatalf("failed to initialize config: %v", err)
		}
	}
	c.init()
	c.Write()
}

func writeConfigToFile(c Config, configName string) error {
	file, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		log.Printf("%v failed to marshal json: %v", color.RedString("ERROR:"), err)
	}
	configDir, err := getConfigDir()
	if err != nil {
		return err
	}
	configPath := filepath.Join(configDir, configName)
	if err := createFolder(configDir); err != nil {
		log.Fatalf("failed to create folder %s: %v", configDir, err)
	}
	return ioutil.WriteFile(configPath, file, 0644)
}

func readConfigFromFile(c Config, configName string) error {
	configDir, err := getConfigDir()
	if err != nil {
		return err
	}
	configPath := filepath.Join(configDir, configName)
	file, e := ioutil.ReadFile(configPath)
	if e != nil {
		return e
	}
	return json.Unmarshal([]byte(file), c)
}

func createFolder(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, 0700)
	}
	return nil
}
