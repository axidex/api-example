package config_provider

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
)

type ConfigType int
type ConfigParseFunc func(filename string) (*viper.Viper, error)

const (
	EnvConfig ConfigType = iota + 1
	FileConfig
)

var loadFunctions = map[ConfigType]ConfigParseFunc{
	EnvConfig:  loadEnvConfig,
	FileConfig: loadFileConfig,
}

// LoadConfig Load config file from given path
func loadFileConfig(filename string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName(filename)
	v.SetConfigType("yml")

	v.AddConfigPath("./config")
	v.AddConfigPath(".")

	if err := v.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return v, nil
}

func loadEnvConfig(filename string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigName(filename)
	v.SetConfigType("env")
	v.AddConfigPath(".")

	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFoundError) {
			return nil, err
		}
	}

	return v, nil
}

// ParseConfig Parse config file
func ParseConfig[T any](filename string, configType ConfigType) (*T, error) {
	loadFunction, ok := loadFunctions[configType]
	if !ok {
		return nil, fmt.Errorf("config type %d not supported", configType)
	}

	v, err := loadFunction(filename)
	if err != nil {
		return nil, err
	}

	var c T
	if err = v.Unmarshal(&c); err != nil {
		return nil, err
	}

	return &c, nil
}
