package config

import (
	"os"
	"strings"

	"github.com/okocraft/monitor/lib/errlib"
)

type HTTPServerConfig struct {
	Port             string
	AllowedOrigins   map[string]struct{}
	DBConfig         DBConfig
	AuthConfig       AuthConfig
	GoogleAuthConfig GoogleAuthConfig
}

func NewHTTPServerConfigFromEnv() (HTTPServerConfig, error) {
	port, err := getRequiredString("MONITOR_HTTP_PORT")
	if err != nil {
		return HTTPServerConfig{}, errlib.AsIs(err)
	}

	origins := createOriginSet(os.Getenv("MONITOR_ALLOWED_ORIGINS"))

	dbConfig, err := NewDBConfigFromEnv()
	if err != nil {
		return HTTPServerConfig{}, errlib.AsIs(err)
	}

	authConfig, err := NewAuthConfigFromEnv()
	if err != nil {
		return HTTPServerConfig{}, errlib.AsIs(err)
	}

	googleAuthConfig, err := NewGoogleAuthConfigFromEnv()

	return HTTPServerConfig{
		Port:             port,
		AllowedOrigins:   origins,
		DBConfig:         dbConfig,
		AuthConfig:       authConfig,
		GoogleAuthConfig: googleAuthConfig,
	}, nil
}

func createOriginSet(value string) map[string]struct{} {
	if value == "" {
		return map[string]struct{}{}
	}

	origins := strings.Split(value, ",")
	if len(origins) == 0 {
		return map[string]struct{}{}
	}

	set := make(map[string]struct{}, len(origins))
	for _, origin := range origins {
		origin = strings.TrimSpace(origin)
		set[origin] = struct{}{}
	}
	return set
}
