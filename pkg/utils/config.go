package utils

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/fatih/color"
)

type Config interface {
	Read() error
	Write() error
	Init(rd io.Reader) error
}

func getConfigDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Printf("%v failed to determine home directory: %v", color.RedString("ERROR:"), err)
		return "", err
	}
	configDir := filepath.Join(homeDir, ".bro")
	return configDir, err
}

func InitializeConfig(c Config) error {
	if err := c.Read(); err != nil {
		if err := c.Write(); err != nil {
			log.Printf("%v failed to initialize config: %v", color.RedString("ERROR:"), err)
			return err
		}
	}
	if err := c.Init(os.Stdin); err != nil {
		return err
	}
	if err := c.Write(); err != nil {
		return err
	}
	return nil
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
		log.Printf("%v failed to create folder %s: %v", color.RedString("ERROR:"), configDir, err)
	}
	return os.WriteFile(configPath, file, 0644)
}

func ReadConfigFromFile(c Config, configName string) error {
	configDir, err := getConfigDir()
	if err != nil {
		return err
	}
	configPath := filepath.Join(configDir, configName)
	file, e := os.ReadFile(configPath)
	if e != nil {
		return e
	}
	return json.Unmarshal([]byte(file), c)
}

func DeleteConfig(configName string) error {
	configDir, err := getConfigDir()
	if err != nil {
		return err
	}
	configPath := filepath.Join(configDir, configName)
	return os.Remove(configPath)
}

func createFolder(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, 0700)
	}
	return nil
}
