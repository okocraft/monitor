package config

import (
	"time"

	"github.com/okocraft/monitor/lib/errlib"
)

type CleanupConfig struct {
	DBConfig                   DBConfig
	AccessTokenExpireDuration  time.Duration
	RefreshTokenExpireDuration time.Duration
}

func NewCleanupConfigFromEnv() (CleanupConfig, error) {
	dbConfig, err := NewDBConfigFromEnv()
	if err != nil {
		return CleanupConfig{}, errlib.AsIs(err)
	}

	accessTokenExpire, err := getDurationFromEnv("MONITOR_ACCESS_TOKEN_EXPIRE", 15*time.Minute)
	if err != nil {
		return CleanupConfig{}, errlib.AsIs(err)
	}

	refreshTokenExpire, err := getDurationFromEnv("MONITOR_REFRESH_TOKEN_EXPIRE", 7*24*time.Hour)
	if err != nil {
		return CleanupConfig{}, errlib.AsIs(err)
	}

	return CleanupConfig{
		DBConfig:                   dbConfig,
		AccessTokenExpireDuration:  accessTokenExpire,
		RefreshTokenExpireDuration: refreshTokenExpire,
	}, nil
}
