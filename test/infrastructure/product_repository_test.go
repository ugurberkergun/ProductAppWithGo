package infrastructure

import (
	"ProductAppWithGo/persistence"
	"github.com/jackc/pgx/v4/pgxpool"
	"testing"
)

//integration test: bu testler integration testtir.Çünk bunların hepsinin bir entegrasyon noktası var bunlar
// db miz ile entegre.Db ye yazdığımız logic ve sorgularımızı doğru entegre ettik mi diye test ediyoruz.

import (
	"ProductAppWithGo/common/postgresql"
	"ProductAppWithGo/domain"
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
)

var productRepository persistence.IProductRepository

var dbPool *pgxpool.Pool
var ctx context.Context

func TestMain(m *testing.M) {
	ctx = context.Background()

	dbPool = postgresql.GetConnectionPool(ctx, postgresql.Config{
		Host:                   "localhost",
		Port:                   "6432",
		DbName:                 "productapp",
		UserName:               "postgres",
		Password:               "postgres",
		MaxConnections:         "10",
		MaxConnectionsIdleTime: "30s",
	})
	productRepository = persistence.NewProductRepository(dbPool)
	fmt.Println("Before all tests")
	exitCode := m.Run()
	fmt.Println("After all tests")
	os.Exit(exitCode)
}

func setup(ctx context.Context, dbPool *pgxpool.Pool) {
	TestDataInitialize(ctx, dbPool)
}
func clear(ctx context.Context, dbPool *pgxpool.Pool) {
	TruncateTestData(ctx, dbPool)
}

func TestGetAllProducts(t *testing.T) {
	setup(ctx, dbPool)

	expected := []domain.Product{
		domain.Product{
			Id:       1,
			Name:     "AirFryer",
			Price:    3000.0,
			Discount: 22.0,
			Store:    "ABC TECH",
		}, domain.Product{
			Id:       2,
			Name:     "Ütü",
			Price:    1500.0,
			Discount: 10.0,
			Store:    "ABC TECH",
		}, domain.Product{
			Id:       3,
			Name:     "Çamaşır Makinesi",
			Price:    10000.0,
			Discount: 15.0,
			Store:    "ABC TECH",
		}, domain.Product{
			Id:       4,
			Name:     "Lambader",
			Price:    2000.0,
			Discount: 0.0,
			Store:    "Dekorasyon Sarayı",
		},
	}

	t.Run("GetAllProducts", func(t *testing.T) {
		actualProducts := productRepository.GetAllProducts()
		assert.Equal(t, 4, len(actualProducts))
		assert.Equal(t, expected, actualProducts)
	})
	clear(ctx, dbPool)
}

func TestAddProduct(t *testing.T) {

	expected := []domain.Product{
		domain.Product{
			Id:       1,
			Name:     "Kupa",
			Price:    100.0,
			Discount: 0.0,
			Store:    "Kırtasiye Merkezi",
		},
	}
	newProduct := domain.Product{
		Name:     "Kupa",
		Price:    100.0,
		Discount: 0.0,
		Store:    "Kırtasiye Merkezi",
	}
	t.Run("AddProduct", func(t *testing.T) {
		productRepository.AddProduct(newProduct)
		actualProducts := productRepository.GetAllProducts()
		assert.Equal(t, 1, len(actualProducts))
		assert.Equal(t, expected, actualProducts)
	})
	clear(ctx, dbPool)
}

func TestGetProductById(t *testing.T) {
	setup(ctx, dbPool)

	t.Run("GetProductById", func(t *testing.T) {
		product, _ := productRepository.GetById(1)
		_, err := productRepository.GetById(5)
		assert.Equal(t, domain.Product{
			Id:       1,
			Name:     "AirFryer",
			Price:    3000.0,
			Discount: 22.0,
			Store:    "ABC TECH",
		}, product)
		assert.Equal(t, "error while getting product with id 5", err.Error())
	})
	clear(ctx, dbPool)
}
