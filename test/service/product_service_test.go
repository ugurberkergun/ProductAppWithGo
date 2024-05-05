package service

import (
	"ProductAppWithGo/domain"
	"ProductAppWithGo/service"
	"ProductAppWithGo/service/model"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var productService service.IProductService

func TestMain(m *testing.M) {

	//Db ye veya herhangi bir entegrasyon noktasına bağımlı olmaaması için fake data oluşturuyoruz
	initialProducts := []domain.Product{
		domain.Product{
			Id:    1,
			Name:  "Airfryer",
			Price: 1000.0,
			Store: "Abc Tech",
		},
		domain.Product{
			Id:    2,
			Name:  "Ütü",
			Price: 4000.0,
			Store: "Abc Tech",
		},
	}

	fakeProductRepository := NewFakeProductRepository(initialProducts)
	productService = service.NewProductService(fakeProductRepository)
	exitCode := m.Run()
	os.Exit(exitCode)
}

func Test_ShouldGetAllProducts(t *testing.T) {
	t.Run("ShouldGetAllProducts", func(t *testing.T) {
		products := productService.GetAllProducts()
		assert.Equal(t, 2, len(products))
	})
}

func Test_WhenNoValidationErrorOccured_ShouldAddProduct(t *testing.T) {
	t.Run("Test_WhenNoValidationErrorOccured_ShouldAddProduct", func(t *testing.T) {
		productService.Add(model.ProductCreate{
			Name:     "Ütü",
			Price:    2000.0,
			Discount: 50,
			Store:    "Abc Tech",
		})
		actualProducts := productService.GetAllProducts()
		assert.Equal(t, 3, len(actualProducts))
		assert.Equal(t, domain.Product{
			Id:       3,
			Name:     "Ütü",
			Price:    2000.0,
			Discount: 50,
			Store:    "Abc Tech",
		}, actualProducts[len(actualProducts)-1])
	})
}
func Test_WhenDiscountIsHigherThan70_ShouldNotAddProduct(t *testing.T) {
	t.Run("Test_WhenNoValidationErrorOccured_ShouldAddProduct", func(t *testing.T) {
		err := productService.Add(model.ProductCreate{
			Name:     "Ütü",
			Price:    2000.0,
			Discount: 80,
			Store:    "Abc Tech",
		})
		actualProducts := productService.GetAllProducts()
		assert.Equal(t, 2, len(actualProducts))
		assert.Equal(t, "Product Discount can not be greater than 70", err.Error())
	})
}
