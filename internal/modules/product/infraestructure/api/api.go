package api

import (
	"net/http"

	"golang-crud-exagonal-arquitecture/internal/modules/product/model/entity"
	product "golang-crud-exagonal-arquitecture/internal/modules/product/model/service"
	shared "golang-crud-exagonal-arquitecture/internal/modules/shared/model"

	"github.com/labstack/echo/v4"
)

type ProductApi struct {
	service product.ProductService
	api     *echo.Echo
}

func NewProductApi(
	service product.ProductService,
	api *echo.Echo,
) *ProductApi {
	rest_api := &ProductApi{
		service: service,
		api:     api,
	}

	return rest_api
}

func (l *ProductApi) Register() {
	logs := l.api.Group("/products")
	logs.GET("", l.GetAllProducts)
	logs.POST("", l.saveProduct)
}

func (p *ProductApi) GetAllProducts(c echo.Context) error {
	products, _ := p.service.GetAllProducts()
	return c.JSON(http.StatusOK, products)
}

func (p *ProductApi) saveProduct(c echo.Context) error {
	params := entity.Product{}
	c.Bind(&params)

	err := p.service.RegisterProduct(params)

	if err != nil {
		return c.JSON(http.StatusBadRequest, shared.ResponseMessage{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, nil)
}
