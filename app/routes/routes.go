package routes

import (
	"errors"
	middlewareApp "laundro-api-ca/app/middleware"
	controller "laundro-api-ca/controllers"
	"laundro-api-ca/controllers/users"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware      middleware.JWTConfig
	UserController     users.UserController
}

func (ctrlList *ControllerList) RouteRegister(e *echo.Echo) {
	users := e.Group("users")
	users.POST("", ctrlList.UserController.Register)
	users.GET("/login", ctrlList.UserController.Login)
}

func RoleValidation(role int, userControler users.UserController) echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := middlewareApp.GetUser(c)

			userRole := userControler.GetRoleByID(claims.ID)

			if userRole == role {
				return hf(c)
			} else {
				return controller.NewErrorResponse(c, http.StatusForbidden, errors.New("forbidden roles"))
			}
		}
	}
}