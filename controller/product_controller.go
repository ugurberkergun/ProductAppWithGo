package controller

import (
	"ProductAppWithGo/controller/request"
	"ProductAppWithGo/controller/response"
	"ProductAppWithGo/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type ProductController struct {
	productService service.IProductService
}

func NewProductController(productService service.IProductService) *ProductController {
	return &ProductController{
		productService: productService,
	}
}

func (productController *ProductController) RegisterRoutes(e *echo.Echo) {
	e.GET("/api/v1/products/:id", productController.GetProductById)
	e.GET("/api/v1/products", productController.GetAllProducts)
	e.POST("/api/v1/products", productController.AddProduct)
	e.PUT("/api/v1/products/:id", productController.UpdateProducts)
	e.DELETE("/api/v1/products/:id", productController.DeleteProduct)
}

func (productController *ProductController) GetAllProducts(c echo.Context) error {
	store := c.QueryParam("store")
	if len(store) == 0 {
		allProducts := productController.productService.GetAllProducts()
		return c.JSON(http.StatusOK, response.ToResponseList(allProducts))
	}
	productsWithGivenStore := productController.productService.GetAllProductsByStore(store)
	return c.JSON(http.StatusOK, response.ToResponseList(productsWithGivenStore))
}

func (productController *ProductController) GetProductById(c echo.Context) error {
	param := c.Param("id")
	productId, _ := strconv.Atoi(param)

	product, err := productController.productService.GetById(int64(productId))
	if err != nil {
		return c.JSON(http.StatusNotFound, response.ErrorResponse{
			ErrorDescription: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.ToResponse(product))

}

func (productController *ProductController) AddProduct(c echo.Context) error {
	//productController.productService.Add()
	var addProductRequest request.AddProductRequest
	err := c.Bind(&addProductRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: err.Error(),
		})
	}
	addErr := productController.productService.Add(addProductRequest.ToModel())
	if addErr != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.ErrorResponse{
			ErrorDescription: addErr.Error(),
		})
	}
	return c.NoContent(http.StatusCreated)
}

func (productController *ProductController) UpdateProducts(c echo.Context) error {
	param := c.Param("id")
	productId, _ := strconv.Atoi(param)

	newPrice := c.QueryParam("newPrice")

	if len(newPrice) == 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: "Parameter newPrice is required",
		})
	}
	newPriceFloat, err := strconv.ParseFloat(newPrice, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: err.Error(),
		})
	}
	updateErr := productController.productService.UpdatePrice(int64(productId), float32(newPriceFloat))
	if updateErr != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.ErrorResponse{
			ErrorDescription: updateErr.Error(),
		})
	}
	return c.NoContent(http.StatusOK)
}

func (productController *ProductController) DeleteProduct(c echo.Context) error {
	param := c.Param("id")
	productId, _ := strconv.Atoi(param)

	err := productController.productService.DeleteById(int64(productId))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.ErrorResponse{
			ErrorDescription: err.Error(),
		})
	}
	return c.NoContent(http.StatusOK)
}
