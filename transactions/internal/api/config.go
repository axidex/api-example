package api

import "fmt"

type Config struct {
	Port int `mapstructure:"API_PORT"`
}

func (c Config) Info() string {
	return fmt.Sprintf("API: Port - %d", c.Port)
}
