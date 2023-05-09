package configs

import (
	"fmt"
	"os"
	"time"

	"github.com/pkg/errors"

	"gopkg.in/yaml.v3"
)

type Probes struct {
	Readiness string `yaml:"readiness"`
	Liveness  string `yaml:"liveness"`
}

type API struct {
	Timeout   time.Duration `yaml:"timeout"`
	ServeAddr string        `yaml:"serve_addr"`
	Probes    Probes        `yaml:"probes"`
	DB        Database      `yaml:"database"`
}

type Database struct {
	Kind           string `yaml:"kind"`
	Host           string `yaml:"host"`
	Port           string `yaml:"port"`
	UserEnvKey     string `yaml:"user_env_key"`
	PassEnvKey     string `yaml:"pass_env_key"`
	DBName         string `yaml:"dbname"`
	MaxConnections int    `yaml:"max_connections"`
	Timeout        int    `yaml:"timeout"`
}

func (d *Database) String() string {
	passPlaceholder := ""
	for range os.Getenv(d.PassEnvKey) {
		passPlaceholder += "*"
	}

	return fmt.Sprintf(
		"DB: (Kind: %s, Host: %s, Port: %s, User: %s, Pass: %s, DBName: %s, MaxConn: %d, Timeout: %d)",
		d.Kind, d.Host, d.Port, os.Getenv(d.UserEnvKey), passPlaceholder, d.DBName, d.MaxConnections, d.Timeout)
}

func Read(path string, cfg interface{}) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return errors.Errorf("cant read config file: %s", err.Error())
	}

	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		return errors.Errorf("cant parse config: %s", err.Error())
	}

	return nil
}
