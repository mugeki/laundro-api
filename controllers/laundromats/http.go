package laundromats

import (
	"laundro-api-ca/app/middleware"
	"laundro-api-ca/business/laundromats"
	controller "laundro-api-ca/controllers"
	"laundro-api-ca/controllers/laundromats/request"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
)

type LaundromatController struct {
	laundromatService laundromats.Service
}

func NewLaundromatController(service laundromats.Service) *LaundromatController {
	return &LaundromatController{
		laundromatService: service,
	}
}

func (ctrl *LaundromatController) Insert(c echo.Context) error {
	req := request.Laundromats{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	if _, err := govalidator.ValidateStruct(req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	userID := uint(middleware.GetUser(c).ID)
	laundroData, addrData := req.ToDomain()
	data, err := ctrl.laundromatService.Insert(userID, laundroData, addrData)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, data)
}

func (ctrl *LaundromatController) GetByIP(c echo.Context) error {
	data, err :=  ctrl.laundromatService.GetByIP()
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusNotFound, err)
	}
	return controller.NewSuccessResponse(c,data)
}

func (ctrl *LaundromatController) GetByName(c echo.Context) error {
	name := c.Param("name")
	data, err := ctrl.laundromatService.GetByName(name)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusNotFound, err)
	}
	return controller.NewSuccessResponse(c, data)
}

func (ctrl *LaundromatController) GetByID(id int) laundromats.Domain{
	data, err := ctrl.laundromatService.GetByID(uint(id))
	if err != nil {
		return laundromats.Domain{}
	}
	return data
}

func (ctrl *LaundromatController) Update(c echo.Context) error {
	req := request.Laundromats{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	if _, err := govalidator.ValidateStruct(req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	param := c.Param("id")
	laundroID, _ := strconv.Atoi(param)
	laundroData, addrData := req.ToDomain()
	data, err := ctrl.laundromatService.Update(uint(laundroID), laundroData, addrData)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, data)
}

func (ctrl *LaundromatController) Delete(c echo.Context) error {
	param := c.Param("id")
	laundroID, _ := strconv.Atoi(param)
	data, err := ctrl.laundromatService.Delete(uint(laundroID))
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, data)
}