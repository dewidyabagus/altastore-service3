package middleware

import (
	"AltaStore/config"
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	middleware.ErrJWTMissing = echo.NewHTTPError(
		http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"message": "missing or malformed jwt",
		})
	middleware.ErrJWTInvalid = echo.NewHTTPError(
		http.StatusUnauthorized, map[string]interface{}{
			"code":    401,
			"massage": "invalid or expired jwt",
		})
}

func JWTMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte(config.GetConfig().JwtSecretKey),
	})
}

func ExtractTokenUser(ctx echo.Context) (string, error) {
	user := ctx.Get("user").(*jwt.Token)
	if user.Valid {
		claim := user.Claims.(jwt.MapClaims)
		userId := claim["userId"].(string)
		if userId == "" {
			return "", errors.New("Unauthorize")
		}
		return userId, nil
	}
	return "", errors.New("Unauthorize")
}

func ExtractTokenRule(ctx echo.Context) (bool, error) {
	user := ctx.Get("user").(*jwt.Token)
	if user.Valid {
		claim := user.Claims.(jwt.MapClaims)
		isAdmin := claim["isAdmin"].(bool)
		return isAdmin, nil
	}
	return false, errors.New("Unauthorize")
}
