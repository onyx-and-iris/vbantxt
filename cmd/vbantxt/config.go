package main

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"
)

type config struct {
	Connection connection `toml:"connection"`
}

func (c config) String() string {
	return fmt.Sprintf(
		"host: %s port: %d streamname: %s",
		c.Connection.Host, c.Connection.Port, c.Connection.Streamname)
}

type connection struct {
	Host       string `toml:"host"`
	Port       int    `toml:"port"`
	Streamname string `toml:"streamname"`
}

func loadConfig(configPath string) (*connection, error) {
	_, err := os.Stat(configPath)
	if err != nil {
		return nil, err
	}

	var config config

	_, err = toml.DecodeFile(configPath, &config)
	if err != nil {
		return nil, err
	}
	log.Debug(config)

	return &config.Connection, nil
}
