package middleware

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/muhrizqiardi/spendtracker/internal/service"
)

type AuthMiddleware interface {
	Authenticate(next echo.HandlerFunc) echo.HandlerFunc
}

type authMiddleware struct {
	us     service.UserService
	secret string
}

func NewAuthMiddleware(us service.UserService, secret string) *authMiddleware {
	return &authMiddleware{us, secret}
}

func (am *authMiddleware) Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		authorizationHeader := strings.Split(ctx.Request().Header.Get("Authorization"), " ")
		if len(authorizationHeader) < 2 {
			ctx.Logger().Error("invalid Authorization header")
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}

		token, err := jwt.Parse(authorizationHeader[1], func(t *jwt.Token) (interface{}, error) {
			return []byte(am.secret), nil
		})
		if err != nil {
			ctx.Logger().Error(err.Error())
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok && !token.Valid {
			ctx.Logger().Error(err.Error())
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}

		var claimsSub string
		if _, ok := claims["sub"]; ok {
			claimsSub, ok = claims["sub"].(string)
			if !ok {
				ctx.Logger().Error(err.Error())
				return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
			}
		}
		userIDInt, err := strconv.Atoi(claimsSub)
		if err != nil {
			ctx.Logger().Error(err.Error())
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}

		user, err := am.us.GetOneByID(userIDInt)
		if err != nil {
			ctx.Logger().Error(err.Error())
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}

		ctx.Set("user", user)
		return next(ctx)
	}
}
