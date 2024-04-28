package persistence

import (
	"ProductAppWithGo/domain"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

type IProductRepository interface {
	GetAllProducts() []domain.Product
}

type ProductRepository struct {
	dbPool *pgxpool.Pool
}

func NewProductRepository(dbPool *pgxpool.Pool) IProductRepository {
	return &ProductRepository{dbPool: dbPool}
}

func (productRepository *ProductRepository) GetAllProducts() []domain.Product {
	ctx := context.Background()
	productRows, err := productRepository.dbPool.Query(ctx, "Select * from products")
	if err != nil {
		log.Error("Error while getting all products: %v", err)
		return []domain.Product{}
	}

	var products = []domain.Product{}
	var id int64
	var name string
	var price float32
	var discount float32
	var store string

	//okunacak satır var mı
	for productRows.Next() {
		err := productRows.Scan(&id, &name, &price, &discount, &store)
		if err != nil {

		}
		products = append(products, domain.Product{id, name, price, discount, store})
	}
	return products
}
