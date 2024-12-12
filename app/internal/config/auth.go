package config

import (
	"time"

	"github.com/okocraft/monitor/lib/errlib"
)

type AuthConfig struct {
	HMACSecret                 []byte
	LoginExpireDuration        time.Duration
	AccessTokenExpireDuration  time.Duration
	RefreshTokenExpireDuration time.Duration
}

func NewAuthConfigFromEnv() (AuthConfig, error) {
	hmacSecret, err := getRequiredString("MONITOR_HMAC_SECRET")
	if err != nil {
		return AuthConfig{}, errlib.AsIs(err)
	}

	loginExpire, err := getDurationFromEnv("MONITOR_LOGIN_EXPIRE", 15*time.Minute)
	if err != nil {
		return AuthConfig{}, errlib.AsIs(err)
	}

	accessTokenExpire, err := getDurationFromEnv("MONITOR_ACCESS_TOKEN_EXPIRE", 15*time.Minute)
	if err != nil {
		return AuthConfig{}, errlib.AsIs(err)
	}

	refreshTokenExpire, err := getDurationFromEnv("MONITOR_REFRESH_TOKEN_EXPIRE", 7*24*time.Hour)
	if err != nil {
		return AuthConfig{}, errlib.AsIs(err)
	}

	return AuthConfig{
		HMACSecret:                 []byte(hmacSecret),
		LoginExpireDuration:        loginExpire,
		AccessTokenExpireDuration:  accessTokenExpire,
		RefreshTokenExpireDuration: refreshTokenExpire,
	}, nil
}
