package configs

import (
	"fmt"
	"os"
	"path/filepath"
)

type ConfigFilePath string

type Configuration interface {
	GetInt(key string, defaultVal int) (int, error)
	GetFloat(key string, defaultVal float64) (float64, error)
	GetString(key, defaultVal string) string
	GetBool(key string, defaultVal bool) (bool, error)
}

func LoadConfig(filePath ConfigFilePath) (Configuration, error) {
	fullPath, err := filepath.Abs(string(filePath))
	if err != nil {
		return nil, fmt.Errorf("invalid config file: %s", filePath)
	}

	fileBytes, err := os.ReadFile(fullPath)
	if err != nil {
		return nil, err
	}

	switch filepath.Ext(fullPath) {
	case ".ini":
		return NewIniConfig(fileBytes)
	default:
		return nil, fmt.Errorf("not supported config file: %s", filePath)
	}
}
