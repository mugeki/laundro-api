package orders

import (
	"laundro-api-ca/app/middleware"
	"laundro-api-ca/business/orders"
	controller "laundro-api-ca/controllers"
	"laundro-api-ca/controllers/orders/request"
	"laundro-api-ca/controllers/orders/response"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
)

type OrderController struct {
	orderService orders.Service
}

func NewOrderController(service orders.Service) *OrderController {
	return &OrderController{
		orderService: service,
	}
}

func (ctrl *OrderController) Create(c echo.Context) error {
	req := request.Orders{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	if _, err := govalidator.ValidateStruct(req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	
	userID := middleware.GetUser(c).ID

	data, err := ctrl.orderService.Create(uint(userID), req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controller.NewSuccessResponse(c, data)
}

func (ctrl *OrderController) GetByUserID(c echo.Context) error {
	userID := middleware.GetUser(c).ID
	data, err := ctrl.orderService.GetByUserID(uint(userID))
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusNotFound, err)
	}
	res := response.FromDomainArray(data)
	return controller.NewSuccessResponse(c, res)
}