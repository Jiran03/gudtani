package routes

import (
	"errors"

	authMiddleware "github.com/Jiran03/gudtani/auth/middleware"
	errConv "github.com/Jiran03/gudtani/helper/error"
	userAPI "github.com/Jiran03/gudtani/user/handler/api"
	"github.com/labstack/echo/v4"
)

func RoleValidation(role string, userController userAPI.UserHandler) echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			claims := authMiddleware.GetUser(ctx)
			userRole, err := userController.UserRole(claims.ID)

			if err != nil {
				return errors.New(errConv.ErrDBNotFound)
			}

			if userRole == role {
				return hf(ctx)
			} else {
				return errors.New(errConv.ErrUserFailure)
			}
		}
	}
}
