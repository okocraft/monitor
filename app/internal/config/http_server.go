package config

import (
	"github.com/okocraft/monitor/lib/errlib"
)

type HTTPServerConfig struct {
	Port     string
	DBConfig DBConfig
}

func NewHTTPServerConfigFromEnv() (HTTPServerConfig, error) {
	port, err := getRequiredString("MONITOR_HTTP_PORT")
	if err != nil {
		return HTTPServerConfig{}, errlib.AsIs(err)
	}

	dbConfig, err := NewDBConfigFromEnv()
	if err != nil {
		return HTTPServerConfig{}, errlib.AsIs(err)
	}

	return HTTPServerConfig{
		Port:     port,
		DBConfig: dbConfig,
	}, nil
}
