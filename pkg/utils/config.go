package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	color "github.com/threeal/bro/pkg/utils/color"
)

type Config interface {
	Read() error
	Write() error
	init()
}

var HOME_DIR = getHomeDirectory()
var BRO_CONFIG_DIR = filepath.Join(*HOME_DIR, ".bro")

type BackendConfig struct {
	ListenAddr *string `json:"listen_addr"`
}

type ClientConfig struct {
	BackendAddr *string `json:"backend_addr"`
}

func getHomeDirectory() (*string) {
	homeDir, err := os.UserHomeDir()
	if (err != nil) {
		log.Fatalf("failed to determine home directory: %v", err)
	}
	return &homeDir
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

func (c *BackendConfig) init() {
	if c.ListenAddr == nil {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(color.Gray + "question" + color.Reset + " listen address (:320): ")
		text, _ := reader.ReadString('\n')
		if text == "\n" {
			text = ":320"
		}
		c.ListenAddr = &text
	}
}

func (c *ClientConfig) init() {
	if c.BackendAddr == nil {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(color.Gray + "question" + color.Reset + " backend address (localhost:320): ")
		text, _ := reader.ReadString('\n')
		if text == "\n" {
			text = "localhost:320"
		}
		c.BackendAddr = &text
	}
}

func NewBackendConfig() (*BackendConfig) {
	return &BackendConfig{ListenAddr: nil}
}

func NewClientConfig() (*ClientConfig) {
	return &ClientConfig{BackendAddr: nil}
}

func InitializeBackendConfig() (*BackendConfig) {
	backendConfig := NewBackendConfig()
	initializeConfig(backendConfig)
	return backendConfig
}

func InitializeClientConfig() (*ClientConfig) {
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
	file, _ := json.MarshalIndent(c, "", "  ")
	configPath := filepath.Join(BRO_CONFIG_DIR, configName)
	if err := createFolder(BRO_CONFIG_DIR); err != nil {
		log.Fatalf("failed to create folder %s: %v", BRO_CONFIG_DIR, err)
	}
	return ioutil.WriteFile(configPath, file, 0644)
}

func readConfigFromFile(c Config, configName string) error {
	configPath := filepath.Join(BRO_CONFIG_DIR, configName)
	file, e := ioutil.ReadFile(configPath)
	if e != nil {
		return e
	}
	return json.Unmarshal([]byte(file), c)
}

func createFolder(path string) error {
	if _, err := os.Stat(BRO_CONFIG_DIR); os.IsNotExist(err) { 
		return os.MkdirAll(BRO_CONFIG_DIR, 0700)
	}
	return nil
}
