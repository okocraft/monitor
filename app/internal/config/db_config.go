package config

import "github.com/okocraft/monitor/lib/errlib"

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func NewDBConfigFromEnv() (DBConfig, error) {
	host, err := getRequiredString("MONITOR_DB_HOST")
	if err != nil {
		return DBConfig{}, errlib.AsIs(err)
	}

	port, err := getRequiredString("MONITOR_DB_PORT")
	if err != nil {
		return DBConfig{}, errlib.AsIs(err)
	}

	user, err := getRequiredString("MONITOR_DB_USER")
	if err != nil {
		return DBConfig{}, errlib.AsIs(err)
	}

	password, err := getRequiredString("MONITOR_DB_PASSWORD")
	if err != nil {
		return DBConfig{}, errlib.AsIs(err)
	}

	dbName, err := getRequiredString("MONITOR_DB_NAME")
	if err != nil {
		return DBConfig{}, errlib.AsIs(err)
	}

	return DBConfig{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		DBName:   dbName,
	}, nil
}
