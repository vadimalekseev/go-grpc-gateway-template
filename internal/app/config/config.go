package config

import (
	"fmt"
	"io/ioutil"

	"github.com/hashicorp/hcl"
)

// Config config section.
type Config struct {
	Database Database `hcl:"database,block"`
	App      App      `hcl:"app,block"`
}

// App config section.
type App struct {
	HTTPAddr       string `hcl:"http_addr"`
	GRPCAddr       string `hcl:"grpc_addr"`
	MetricsAddr    string `hcl:"metrics_addr"`
	UseGRPCReflect bool   `hcl:"use_grpc_reflect"`
}

// Database config section.
type Database struct {
	Host     string `hcl:"host"`
	Port     int    `hcl:"port"`
	User     string `hcl:"user"`
	Password string `hcl:"password"`
	Database string `hcl:"database"`
	SSLMode  string `hcl:"sslmode"`
}

// Parse config from byte slice.
func Parse(bs []byte) (Config, error) {
	result := &Config{}

	if err := hcl.Unmarshal(bs, result); err != nil {
		return *result, fmt.Errorf("error parsing config: %w", err)
	}

	return *result, nil
}

// FromFile parse config from config path.
func FromFile(cfgPath string) (cfg Config, err error) {
	cfgBytes, err := ioutil.ReadFile(cfgPath) //nolint:gosec
	if err != nil {
		return cfg, fmt.Errorf("error reading file: %s", err)
	}

	cfg, err = Parse(cfgBytes)
	if err != nil {
		return cfg, fmt.Errorf("error parsing file: %s", err)
	}

	return cfg, nil
}
