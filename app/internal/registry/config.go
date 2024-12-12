package registry

import (
	"github.com/google/wire"
	"github.com/okocraft/monitor/internal/config"
)

var configSet = wire.NewSet(
	getAuthConfigFromHTTPConfig,
	getGoogleAuthConfigFromHTTPConfig,
)

func getAuthConfigFromHTTPConfig(cfg config.HTTPServerConfig) config.AuthConfig {
	return cfg.AuthConfig
}

func getGoogleAuthConfigFromHTTPConfig(cfg config.HTTPServerConfig) config.GoogleAuthConfig {
	return cfg.GoogleAuthConfig
}
