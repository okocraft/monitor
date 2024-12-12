package usecases

import (
	"context"
	"fmt"
	"time"

	"github.com/Siroshun09/serrors"
	"github.com/gofrs/uuid/v5"
	"github.com/golang-jwt/jwt/v5"
	"github.com/okocraft/monitor/internal/config"
	"github.com/okocraft/monitor/internal/domain/auth"
	"github.com/okocraft/monitor/internal/domain/user"
	"github.com/okocraft/monitor/internal/repositories"
	"github.com/okocraft/monitor/lib/errlib"
)

type AuthUsecase interface {
	CreateStateJWT(ctx context.Context, currentPageURL string) (uuid.UUID, string, error)
	CreateStateJWTWithLoginKey(ctx context.Context, loginKey string) (uuid.UUID, string, error)
	VerifyStateJWT(ctx context.Context, tokenString string) (jwt.MapClaims, error)
	CreateRefreshTokens(ctx context.Context, userID user.ID) (string, time.Time, error)
	VerifyRefreshToken(ctx context.Context, tokenString string) (user.ID, int64, time.Time, error)
	RefreshAccessToken(ctx context.Context, userID user.ID, refreshTokenID int64, maxExpiresAt time.Time) (string, error)
	InvalidateTokens(ctx context.Context, refreshTokenID int64) error
	VerifyAccessToken(ctx context.Context, tokenString string) (user.ID, error)
}

func NewAuthUsecase(conf config.AuthConfig, repo repositories.AuthRepository) AuthUsecase {
	return authUsecase{
		conf: conf,
		repo: repo,
	}
}

type authUsecase struct {
	conf config.AuthConfig
	repo repositories.AuthRepository
}

func (u authUsecase) CreateStateJWT(_ context.Context, currentPageURL string) (uuid.UUID, string, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return uuid.Nil, "", serrors.WithStackTrace(err)
	}

	stateJWT := auth.NewStateJWT(id, time.Now().Add(u.conf.LoginExpireDuration), currentPageURL)

	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS512, stateJWT).SignedString(u.conf.HMACSecret)
	if err != nil {
		return uuid.Nil, "", serrors.WithStackTrace(err)
	}

	return id, tokenString, nil
}

func (u authUsecase) CreateStateJWTWithLoginKey(_ context.Context, loginKey string) (uuid.UUID, string, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return uuid.Nil, "", serrors.WithStackTrace(err)
	}

	stateJWT := auth.NewStateJWTWithLoginKey(id, loginKey, time.Now().Add(u.conf.LoginExpireDuration))

	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS512, stateJWT).SignedString(u.conf.HMACSecret)
	if err != nil {
		return uuid.Nil, "", serrors.WithStackTrace(err)
	}

	return id, tokenString, nil
}

func (u authUsecase) VerifyStateJWT(_ context.Context, tokenString string) (jwt.MapClaims, error) {
	claims, err := u.verifyJWT(tokenString)
	if err != nil {
		return nil, errlib.AsIs(err)
	}

	expiresAt, err := auth.ReadExpiresAtFromJWT(claims)
	if err != nil {
		return nil, errlib.AsIs(err)
	}

	if time.Now().After(expiresAt) {
		return nil, serrors.New("outdated token")
	}

	return claims, nil
}

func (u authUsecase) CreateRefreshTokens(ctx context.Context, userID user.ID) (string, time.Time, error) {
	createdAt := time.Now()
	expiresAt := createdAt.Add(u.conf.RefreshTokenExpireDuration)

	refreshToken, err := u.createRefreshToken(ctx, userID, createdAt, expiresAt)
	if err != nil {
		return "", time.Time{}, errlib.AsIs(err)
	}

	return refreshToken, expiresAt, nil
}

func (u authUsecase) createRefreshToken(ctx context.Context, userID user.ID, createdAt time.Time, expiresAt time.Time) (string, error) {
	jti, err := uuid.NewV4()
	if err != nil {
		return "", serrors.WithStackTrace(err)
	}

	err = u.repo.SaveRefreshToken(ctx, userID, jti, createdAt)
	if err != nil {
		return "", errlib.AsIs(err)
	}

	claims := auth.NewJWT(jti, expiresAt)
	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS512, claims).SignedString(u.conf.HMACSecret)
	if err != nil {
		return "", serrors.WithStackTrace(err)
	}

	return tokenString, nil
}

func (u authUsecase) createAccessToken(ctx context.Context, userID user.ID, refreshTokenID int64, createdAt time.Time, expiresAt time.Time) (string, error) {
	jti, err := uuid.NewV4()
	if err != nil {
		return "", serrors.WithStackTrace(err)
	}

	err = u.repo.SaveAccessToken(ctx, userID, refreshTokenID, jti, createdAt)
	if err != nil {
		return "", errlib.AsIs(err)
	}

	claims := auth.NewJWT(jti, expiresAt)
	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS512, claims).SignedString(u.conf.HMACSecret)
	if err != nil {
		return "", serrors.WithStackTrace(err)
	}

	return tokenString, nil
}

func (u authUsecase) VerifyRefreshToken(ctx context.Context, tokenString string) (user.ID, int64, time.Time, error) {
	claims, err := u.verifyJWT(tokenString)
	if err != nil {
		return 0, 0, time.Time{}, errlib.AsIs(err)
	}

	expiresAt, err := auth.ReadExpiresAtFromJWT(claims)
	if err != nil {
		return 0, 0, time.Time{}, errlib.AsIs(err)
	}

	if time.Now().After(expiresAt) {
		return 0, 0, time.Time{}, auth.NewUnauthorizedError(serrors.New("outdated token"))
	}

	jti, err := auth.ReadJTIFromJWT(claims)
	if err != nil {
		return 0, 0, time.Time{}, errlib.AsIs(err)
	}

	userID, refreshTokenID, err := u.repo.GetUserIDAndRefreshTokenIDFromJTI(ctx, jti)
	if err != nil {
		return 0, 0, time.Time{}, errlib.AsIs(err)
	}

	return userID, refreshTokenID, expiresAt, nil
}

func (u authUsecase) RefreshAccessToken(ctx context.Context, userID user.ID, refreshTokenID int64, maxExpiresAt time.Time) (string, error) {
	createdAt := time.Now()

	expiresAt := createdAt.Add(u.conf.AccessTokenExpireDuration)
	if expiresAt.After(maxExpiresAt) {
		expiresAt = maxExpiresAt
	}

	tokenString, err := u.createAccessToken(ctx, userID, refreshTokenID, time.Now(), expiresAt)
	if err != nil {
		return "", errlib.AsIs(err)
	}

	return tokenString, nil
}

func (u authUsecase) verifyJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return u.conf.HMACSecret, nil
	})

	if err != nil {
		return nil, auth.NewUnauthorizedError(serrors.WithStackTrace(err))
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return claims, nil
	}

	return nil, auth.NewUnauthorizedError(serrors.New("invalid token"))
}

func (u authUsecase) InvalidateTokens(ctx context.Context, refreshTokenID int64) error {
	if err := u.repo.InvalidateTokensByRefreshTokenID(ctx, refreshTokenID); err != nil {
		return errlib.AsIs(err)
	}
	return nil
}

func (u authUsecase) VerifyAccessToken(ctx context.Context, tokenString string) (user.ID, error) {
	claims, err := u.verifyJWT(tokenString)
	if err != nil {
		return user.ID(0), errlib.AsIs(err)
	}

	expiresAt, err := auth.ReadExpiresAtFromJWT(claims)
	if err != nil {
		return user.ID(0), errlib.AsIs(err)
	}

	if time.Now().After(expiresAt) {
		return user.ID(0), auth.NewUnauthorizedError(serrors.New("outdated token"))
	}

	jti, err := auth.ReadJTIFromJWT(claims)
	if err != nil {
		return user.ID(0), errlib.AsIs(err)
	}

	userID, err := u.repo.GetUserIDFromAccessTokenJTI(ctx, jti)
	if err != nil {
		return user.ID(0), errlib.AsIs(err)
	}

	return userID, nil
}
