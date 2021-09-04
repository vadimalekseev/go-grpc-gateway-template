package config

import (
	"fmt"

	"github.com/hashicorp/hcl"
)

type Config struct {
	Database Database `hcl:"database,block"`
}

type Database struct {
	Host     string `hcl:"host"`
	Port     int    `hcl:"port"`
	User     string `hcl:"user"`
	Password string `hcl:"password"`
	DBName   string `hcl:"dbname"`
}

func Parse(bs []byte) (*Config, error) {
	result := &Config{}

	if err := hcl.Unmarshal(bs, result); err != nil {
		return nil, fmt.Errorf("error parsing config: %v", err)
	}

	return result, nil
}
