package config

import (
	"github.com/spf13/viper"
)

type Config interface {
	ReadConfig() error
	GetValue(string) string
}

type ViperConfig struct {
	Path string
	Name string
}

func NewViperConfig(name, path string) ViperConfig {
	viper.SetConfigName(name)
	viper.AddConfigPath(path)
	c := ViperConfig{
		Name: name,
		Path: path,
	}
	return c
}

func (v ViperConfig) ReadConfig() error {
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		return err
	}
	return nil
}

func (v ViperConfig) GetValue(key string) string {
	result := viper.Get(key)
	return result.(string)
}
