package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
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

func logError(s string, err error) {
	log.Printf("%v %s: %v", color.RedString("ERROR:"), s, err)
}

func Prompt(q string, def string, rd io.Reader) (string, error) {
	reader := bufio.NewReader(rd)
	question := fmt.Sprintf(" %s ", q)
	defaultOption := fmt.Sprintf("(%s)", def)
	fmt.Print(color.HiBlackString("question"), question, color.HiGreenString(defaultOption), ": ")
	return reader.ReadString('\n')
}

func getConfigDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		logError("failed to determine home directory", err)
		return "", err
	}
	configDir := filepath.Join(homeDir, ".bro")
	return configDir, err
}

func InitializeConfig(c Config) error {
	if err := c.Read(); err != nil {
		if err := c.Write(); err != nil {
			logError("failed to initialize config", err)
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
		logError("failed to marshal json", err)
	}
	configDir, err := getConfigDir()
	if err != nil {
		return err
	}
	configPath := filepath.Join(configDir, configName)
	if err := createFolder(configDir); err != nil {
		logError(fmt.Sprintf("failed to create folder %s", configDir), err)
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
