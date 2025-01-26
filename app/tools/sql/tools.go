//go:build tools

package sql

//go:generate go run github.com/sqlc-dev/sqlc/cmd/sqlc@v1.28.0 generate
