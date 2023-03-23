package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type (
	Auth interface {
		Auth(next echo.HandlerFunc) echo.HandlerFunc
	}

	authImpl struct {
	}
)

func NewAuth() Auth {
	return &authImpl{}
}

func (m *authImpl) Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		jwtKey := c.Request().Header.Get("X-jwtKey")
		jwtToken, err := extractBearerToken(c)
		if err != nil {
			return err
		}

		c.Set("jwtToken", jwtToken)
		c.Set("jwtKey", jwtKey)

		return next(c)
	}

}

func extractBearerToken(c echo.Context) (string, error) {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return "", echo.NewHTTPError(http.StatusUnauthorized, "missing authorization header")
	}

	authParts := strings.Split(authHeader, " ")
	if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
		return "", echo.NewHTTPError(http.StatusUnauthorized, "invalid authorization header")
	}

	return authParts[1], nil
}
