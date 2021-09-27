package users

import (
	"laundro-api-ca/business/users"
	controller "laundro-api-ca/controllers"
	"laundro-api-ca/controllers/users/request"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userService users.Service
}

func NewUserController(service users.Service) *UserController{
	return &UserController{
		userService: service,
	}
}

func (ctrl *UserController) Register(c echo.Context) error{
	req := request.Users{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	data, err := ctrl.userService.Register(req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, data)
}

func (ctrl *UserController) Login(c echo.Context) error{
	req := request.UsersLogin{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	token, err := ctrl.userService.Login(req.Username, req.Password)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	
	res := struct {
		Token string `json:"token"`
	}{Token: token}

	return controller.NewSuccessResponse(c,res)
}

func (ctrl *UserController) GetRoleByID(id int) int{
	user, err := ctrl.userService.GetByID(uint(id))
	if err != nil {
		return -1
	}
	return int(user.RoleID)
}