package middleware

import (
	"database/sql"
	"errors"
	"github.com/labstack/echo/v4"
	"goBase/app/repositories"
	"goBase/app/services"
	"net/http"
)

func GetUserByKey(c echo.Context) (int, error) {
	db := c.Get("db").(*sql.DB)
	authHeader := c.Request().Header.Get("Authorization")

	if authHeader == "" {
		return 0, errors.New("authorization header is empty")
	}
	userTokenRepo := repositories.NewUserTokenRepository(db)
	userTokenService := services.NewUserTokenService(userTokenRepo)
	userID, err := userTokenService.GetUserByToken(authHeader)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func AuthenticationMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userID, err := GetUserByKey(c)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized: Missing or invalid token")
			}
			c.Set("userID", userID)
			return next(c)
		}
	}
}
