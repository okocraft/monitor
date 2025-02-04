// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package registry

import (
	"github.com/okocraft/monitor/internal/config"
	"github.com/okocraft/monitor/internal/handler"
	"github.com/okocraft/monitor/internal/handler/auditlog"
	"github.com/okocraft/monitor/internal/handler/oapi/auth"
	"github.com/okocraft/monitor/internal/handler/oapi/me"
	"github.com/okocraft/monitor/internal/repositories"
	"github.com/okocraft/monitor/internal/repositories/database"
	"github.com/okocraft/monitor/internal/usecases"
)

// Injectors from wire.go:

func NewHTTPHandler(cfg config.HTTPServerConfig, db database.DB) (handler.HTTPHandler, error) {
	authConfig := getAuthConfigFromHTTPConfig(cfg)
	authRepository := repositories.NewAuthRepository(db)
	authUsecase := usecases.NewAuthUsecase(authConfig, authRepository)
	userRepository := repositories.NewUserRepository(db)
	transaction := database.NewTransaction(db)
	userUsecase := usecases.NewUserUsecase(userRepository, transaction)
	permissionRepository := repositories.NewPermissionRepository(db)
	permissionUsecase := usecases.NewPermissionUsecase(permissionRepository)
	authHandler := auth.NewAuthHandler(authUsecase, userUsecase, permissionUsecase)
	googleAuthConfig := getGoogleAuthConfigFromHTTPConfig(cfg)
	googleAuthHandler := auth.NewGoogleAuthHandler(googleAuthConfig, authUsecase, userUsecase)
	meHandler := me.NewMeHandler(userUsecase)
	auditLogRepository := repositories.NewAuditLogRepository(db)
	auditLogUsecase := usecases.NewAuditLogUsecase(auditLogRepository)
	auditLogMiddleware := auditlog.NewAuditLogMiddleware(auditLogUsecase, userUsecase)
	httpHandler := handler.NewHTTPHandler(authHandler, googleAuthHandler, meHandler, auditLogMiddleware)
	return httpHandler, nil
}

func NewCleanupUsecase(db database.DB) usecases.CleanupUsecase {
	authRepository := repositories.NewAuthRepository(db)
	transaction := database.NewTransaction(db)
	cleanupUsecase := usecases.NewCleanupUsecase(authRepository, transaction)
	return cleanupUsecase
}
