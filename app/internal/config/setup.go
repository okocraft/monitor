package config

import "github.com/okocraft/monitor/lib/errlib"

type SetupConfig struct {
	DBConfig   DBConfig
	ForceSetup bool
}

func NewSetupConfigFromEnv() (SetupConfig, error) {
	dbConfig, err := NewDBConfigFromEnv()
	if err != nil {
		return SetupConfig{}, errlib.AsIs(err)
	}

	forceSetup, err := getBoolFromEnv("MONITOR_FORCE_SETUP", false)
	if err != nil {
		return SetupConfig{}, errlib.AsIs(err)
	}

	return SetupConfig{
		DBConfig:   dbConfig,
		ForceSetup: forceSetup,
	}, nil
}
