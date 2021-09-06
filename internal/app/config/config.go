package config

import (
	"fmt"

	"github.com/hashicorp/hcl"
)

type Config struct {
	Database Database `hcl:"database,block"`
	App App `hcl:"app,block"`
}

type App struct {
	Domain string `hcl:"app"`
	HTTPAddr string `hcl:"httpAddr"`
	GRPCAddr string `hcl:"grpcAddr"`
}

type Database struct {
	Host     string `hcl:"host"`
	Port     int    `hcl:"port"`
	User     string `hcl:"user"`
	Password string `hcl:"password"`
	Database string `hcl:"database"`
	SSLMode  string `hcl:"sslmode"`
}

func Parse(bs []byte) (Config, error) {
	result := &Config{}

	if err := hcl.Unmarshal(bs, result); err != nil {
		return *result, fmt.Errorf("error parsing config: %v", err)
	}

	return *result, nil
}
