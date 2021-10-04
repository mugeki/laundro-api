package products

import (
	"laundro-api-ca/business/products"
	controller "laundro-api-ca/controllers"
	"laundro-api-ca/controllers/products/request"
	"laundro-api-ca/controllers/products/response"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
)

type ProductController struct {
	productService products.Service
}

func NewProductController(service products.Service) *ProductController {
	return &ProductController{
		productService: service,
	}
}

func (ctrl *ProductController) Insert(c echo.Context) error {
	req := request.Products{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	if _, err := govalidator.ValidateStruct(req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	param := c.Param("id")
	laundroID, _ := strconv.Atoi(param)

	data, err := ctrl.productService.Insert(uint(laundroID), req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusNotFound, err)
	}
	res := response.FromDomain(data)
	return controller.NewSuccessResponse(c, res)
}

func (ctrl *ProductController) GetAllByLaundromat(c echo.Context) error {
	param := c.Param("id")
	laundroID, _ := strconv.Atoi(param)

	data, err := ctrl.productService.GetAllByLaundromat(uint(laundroID))
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusNotFound, err)
	}
	res := response.FromDomainArray(data)
	return controller.NewSuccessResponse(c, res)
}

func (ctrl *ProductController) Update(c echo.Context) error {
	req := request.Products{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	if _, err := govalidator.ValidateStruct(req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	
	param := c.Param("productId")
	productID, _ := strconv.Atoi(param)
	
	data, err := ctrl.productService.Update((uint(productID)), req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusNotFound, err)
	}
	res := response.FromDomain(data)
	return controller.NewSuccessResponse(c, res)
}

func (ctrl *ProductController) Delete(c echo.Context) error {
	param := c.Param("productId")
	productID, _ := strconv.Atoi(param)
	data, err := ctrl.productService.Delete(uint(productID))
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusNotFound, err)
	}

	return controller.NewSuccessResponse(c, data)
}

func (ctrl *ProductController) GetLaundromatID(id uint) uint {
	data := ctrl.productService.GetLaundromatID(id)
	return data
}

func (ctrl *ProductController) GetLaundromatByCategory(c echo.Context) error {
	param := c.Param("categoryId")
	categoryID, _ := strconv.Atoi(param)
	data, err := ctrl.productService.GetLaundromatByCategory(categoryID)
	if len(data) == 0 {
		return controller.NewErrorResponse(c, http.StatusNotFound, err)
	}
	return controller.NewSuccessResponse(c, data)
}