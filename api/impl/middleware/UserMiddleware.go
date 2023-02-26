package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func ValidateToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, "Missing authorization header")
		}
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, "Missing token in authorization header")
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(viper.GetString("jwt_secretkey")), nil
		})
		if err != nil {
			return c.JSON(http.StatusUnauthorized, "Invalid token")
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			sub, ok := claims["sub"].(string)
			if !ok {
				return c.JSON(http.StatusUnauthorized, "Invalid token")
			}
			c.Set("user", sub)
			return next(c)
		} else {
			return c.JSON(http.StatusUnauthorized, "Invalid token")
		}
	}
}
