package adapter

import (
	"net/http"
	"strings"
	"user-service/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type MiddlewareAdapterInterface interface {
	CheckToken() echo.MiddlewareFunc
}

type middlewareAdapter struct {
	cfg *config.Config
}

// CheckToken implements MiddlewareAdapterInterface.
func (m *middlewareAdapter) CheckToken() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			redisConn := config.NewRedisClient()
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				log.Errorf("[Middleware-1] CheckToken: %v", "missing or invalid token")
				return echo.NewHTTPError(http.StatusUnauthorized, "missing or invalid token")
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")

			getSession, err := redisConn.HGetAll(c.Request().Context(), tokenString).Result()
			if err != nil || len(getSession) == 0 {
				log.Errorf("[Middleware-2] CheckToken: %v", err.Error())
				return echo.NewHTTPError(http.StatusUnauthorized, "missing or invalid token")
			}
			c.Set("user", getSession)
			return next(c)
		}

	}
}

func NewMiddlewareAdapter(cfg *config.Config) MiddlewareAdapterInterface {
	return &middlewareAdapter{cfg: cfg}
}
