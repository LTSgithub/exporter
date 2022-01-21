package configs

import (
	"strings"

	"gopkg.in/ini.v1"
)

type IniConfig struct {
	f *ini.File
}

func NewIniConfig(fileContent []byte) (Configuration, error) {
	f, err := ini.Load(fileContent)
	if err != nil {
		return nil, err
	}

	return &IniConfig{f: f}, nil
}

func (c IniConfig) split(key string) (section string, subKey string) {
	parts := strings.Split(key, ".")
	if len(parts) > 1 {
		section = parts[0]
		subKey = key[len(section)+1:]
		return
	}

	return key, ""
}

func (c IniConfig) getWithDefault(key string) (*ini.Key, bool, error) {
	section, key := c.split(key)
	if !c.f.Section(section).HasKey(key) {
		return nil, false, nil
	}

	return c.f.Section(section).Key(key), true, nil
}

func (c IniConfig) GetInt(key string, defaultVal int) (int, error) {
	iniKey, exists, err := c.getWithDefault(key)
	if err != nil {
		return 0, err
	}

	if !exists {
		return defaultVal, nil
	}

	return iniKey.Int()
}

func (c IniConfig) GetFloat(key string, defaultVal float64) (float64, error) {
	iniKey, exists, err := c.getWithDefault(key)
	if err != nil {
		return 0, err
	}

	if !exists {
		return defaultVal, nil
	}

	return iniKey.Float64()
}

func (c IniConfig) GetString(key, defaultVal string) string {
	iniKey, exists, _ := c.getWithDefault(key)
	if !exists {
		return defaultVal
	}

	return iniKey.String()
}

func (c IniConfig) GetBool(key string, defaultVal bool) (bool, error) {
	iniKey, exists, err := c.getWithDefault(key)
	if err != nil {
		return false, err
	}

	if !exists {
		return defaultVal, nil
	}

	return iniKey.Bool()
}
