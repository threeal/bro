package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/fatih/color"
)

type Config interface {
	Read() error
	Write() error
	Init() error
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

func InitializeConfig(c Config) Config {
	initializeConfig(c)
	return c
}

func initializeConfig(c Config) {
	if err := c.Read(); err != nil {
		if err = c.Write(); err != nil {
			log.Fatalf("failed to initialize config: %v", err)
		}
	}
	c.Init()
	c.Write()
}

func WriteConfigToFile(c Config, configName string) error {
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

func ReadConfigFromFile(c Config, configName string) error {
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
