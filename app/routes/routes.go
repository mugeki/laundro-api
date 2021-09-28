package routes

import (
	middlewareApp "laundro-api-ca/app/middleware"
	"laundro-api-ca/business"
	controller "laundro-api-ca/controllers"
	"laundro-api-ca/controllers/laundromats"
	"laundro-api-ca/controllers/users"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware        middleware.JWTConfig
	UserController       users.UserController
	LaundromatController laundromats.LaundromatController
}

func (ctrlList *ControllerList) RouteRegister(e *echo.Echo) {
	users := e.Group("users")
	users.POST("", ctrlList.UserController.Register)
	users.POST("/login", ctrlList.UserController.Login)

	laundro := e.Group("laundro", middleware.JWTWithConfig(ctrlList.JWTMiddleware))
	laundro.GET("/find-ip", ctrlList.LaundromatController.GetByIP)
	laundro.GET("/find-name/:name", ctrlList.LaundromatController.GetByName)

	laundroAdmin := laundro
	laundroAdmin.Use(RoleValidation(2,ctrlList.UserController))
	laundroAdmin.POST("", ctrlList.LaundromatController.Insert)
	laundroAdmin.PUT("/edit/:id", ctrlList.LaundromatController.Update, OwnerValidation(ctrlList.LaundromatController))
	laundroAdmin.DELETE("/:id", ctrlList.LaundromatController.Delete, OwnerValidation(ctrlList.LaundromatController))
}

func RoleValidation(roleID int, userController users.UserController) echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := middlewareApp.GetUser(c)

			userRole := userController.GetRoleByID(claims.ID)

			if userRole == roleID {
				return hf(c)
			} else {
				return controller.NewErrorResponse(c, http.StatusForbidden, business.ErrUnauthorized)
			}
		}
	}
}

func OwnerValidation(laundroController laundromats.LaundromatController) echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := middlewareApp.GetUser(c)

			userID := claims.ID
			param := c.Param("id")
			laundroID, _ := strconv.Atoi(param)

			ownerID := int(laundroController.GetByID(laundroID).OwnerID)
			

			if userID == ownerID {
				return hf(c)
			} else {
				return controller.NewErrorResponse(c, http.StatusForbidden, business.ErrUnauthorized)
			}
		}
	}
}