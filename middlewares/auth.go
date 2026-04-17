package middlewares

import (
	"compass-wealth/enums"
	"compass-wealth/utils"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

const authUserContextKey = "auth_user"

type AuthUser struct {
	ID    uint
	Email string
	Role  enums.Roles
}

func JWTUserMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, ok := c.Get("user").(*jwt.Token)
		if !ok || token == nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "missing token")
		}

		claims, ok := token.Claims.(*utils.Claims)
		if !ok || claims == nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid token claims")
		}

		c.Set(authUserContextKey, &AuthUser{
			ID:    claims.ID,
			Email: claims.Email,
			Role:  claims.Role,
		})

		return next(c)
	}
}

func GetAuthUser(c echo.Context) (*AuthUser, error) {
	authUser, ok := c.Get(authUserContextKey).(*AuthUser)
	if !ok || authUser == nil {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "user not found in request context")
	}
	return authUser, nil
}

func AdminOnly(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authUser, err := GetAuthUser(c)
		if err != nil {
			return err
		}

		if authUser.Role != enums.Admin {
			return echo.NewHTTPError(http.StatusForbidden, "admin access required")
		}

		return next(c)
	}
}
