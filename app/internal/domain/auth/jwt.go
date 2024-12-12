package auth

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Siroshun09/serrors"
	"github.com/gofrs/uuid/v5"
	"github.com/golang-jwt/jwt/v5"
)

func NewStateJWT(id uuid.UUID, expiresAt time.Time, currentPageURL string) jwt.Claims {
	return jwt.MapClaims{
		"state":        id.String(),
		"exp":          jwt.NewNumericDate(expiresAt),
		"current_page": currentPageURL,
	}
}

func NewStateJWTWithLoginKey(id uuid.UUID, loginKey string, expiresAt time.Time) jwt.Claims {
	return jwt.MapClaims{
		"state":     id.String(),
		"login_key": loginKey,
		"exp":       jwt.NewNumericDate(expiresAt),
	}
}

func GetLoginKeyFromJWT(claims jwt.MapClaims) (int64, bool) {
	rawLoginKey, ok := claims["login_key"]
	if !ok {
		return 0, false
	}

	loginKey, err := strconv.ParseInt(fmt.Sprintf("%v", rawLoginKey), 10, 64)
	if err != nil {
		return 0, false
	}

	return loginKey, true
}

func GetIDFromJWT(claims jwt.MapClaims) (uuid.UUID, bool) {
	rawID, ok := claims["state"]
	if !ok {
		return uuid.Nil, false
	}

	id, err := uuid.FromString(fmt.Sprintf("%v", rawID))
	if err != nil {
		return uuid.Nil, false
	}

	return id, true
}

func GetCurrentPageURLFromJWT(claims jwt.MapClaims) string {
	rawURL, ok := claims["current_page"]
	if !ok {
		return ""
	}
	return fmt.Sprintf("%v", rawURL)
}

func NewJWT(jti uuid.UUID, expiresAt time.Time) jwt.Claims {
	return jwt.RegisteredClaims{
		ID:        jti.String(),
		ExpiresAt: jwt.NewNumericDate(expiresAt),
	}
}

func ReadJTIFromJWT(claims jwt.MapClaims) (uuid.UUID, error) {
	rawID, ok := claims["jti"]
	if !ok {
		return uuid.Nil, NewUnauthorizedError(serrors.New("missing jti claim"))
	}

	strID, ok := rawID.(string)
	if !ok {
		return uuid.Nil, NewUnauthorizedError(serrors.New("invalid jti claim"))
	}

	jti, err := uuid.FromString(strID)
	if err != nil {
		return uuid.Nil, NewUnauthorizedError(serrors.WithStackTrace(err))
	}

	return jti, nil
}

func ReadExpiresAtFromJWT(claims jwt.MapClaims) (time.Time, error) {
	expiresAt, err := claims.GetExpirationTime()
	if err != nil {
		return time.Time{}, NewUnauthorizedError(serrors.WithStackTrace(err))
	}

	return expiresAt.Time, nil
}
