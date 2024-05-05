package persistence

import (
	"ProductAppWithGo/domain"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

type IProductRepository interface {
	GetAllProducts() []domain.Product
	GetAllProductsByStore(storeName string) []domain.Product
	AddProduct(product domain.Product) error
	GetById(productId int64) (domain.Product, error)
	DeleteById(productId int64) error
	UpdatePrice(newPrice float32, productId int64) error
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
	return extractProductsFromRows(productRows)
}

func (productRepository *ProductRepository) GetAllProductsByStore(storeName string) []domain.Product {
	ctx := context.Background()
	getProductsByStoreNameSql := `
	Select * From products WHERE store = $1;`
	productRows, err := productRepository.dbPool.Query(ctx, getProductsByStoreNameSql, storeName)
	if err != nil {
		log.Error("Error while getting all products: %v", err)
		return []domain.Product{}
	}
	return extractProductsFromRows(productRows)
}

func (productRepository *ProductRepository) AddProduct(product domain.Product) error {
	ctx := context.Background()

	insert_sql := `INSERT INTO products(name,price,discount,store) VALUES($1,$2,$3,$4)`

	addNewProduct, err := productRepository.dbPool.Exec(ctx, insert_sql, product.Name, product.Price, product.Discount, product.Store)
	if err != nil {
		log.Error("Error while adding product: %v", err)
		return err
	}
	log.Info(fmt.Sprintf("Added new product: %v", addNewProduct))
	return nil
}

func extractProductsFromRows(productRows pgx.Rows) []domain.Product {
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

func (productRepository *ProductRepository) GetById(productId int64) (domain.Product, error) {
	ctx := context.Background()

	getByIdSql := `Select * From products WHERE id = $1;`
	productRow := productRepository.dbPool.QueryRow(ctx, getByIdSql, productId)
	var id int64
	var name string
	var price float32
	var discount float32
	var store string
	scanErr := productRow.Scan(&id, &name, &price, &discount, &store)
	if scanErr != nil {
		log.Error("Error while getting product: %v", scanErr)
		return domain.Product{}, errors.New(fmt.Sprintf("error while getting product with id %d", productId))
	}
	return domain.Product{
		Id:       id,
		Name:     name,
		Price:    price,
		Discount: discount,
		Store:    store,
	}, nil
}

func (productRepository *ProductRepository) DeleteById(productId int64) error {
	ctx := context.Background()
	_, getErr := productRepository.GetById(productId)
	if getErr != nil {
		return errors.New("error while getting product")
	}
	deleteByIdSql := `DELETE FROM products WHERE id = $1;`
	_, err := productRepository.dbPool.Exec(ctx, deleteByIdSql, productId)
	if err != nil {
		return errors.New("error while deleting product")
	}
	return nil
}

func (productRepository *ProductRepository) UpdatePrice(newPrice float32, productId int64) error {
	ctx := context.Background()

	updateSql := `UPDATE products SET price = $1 WHERE id = $2;`
	_, err := productRepository.dbPool.Exec(ctx, updateSql, newPrice, productId)
	if err != nil {
		return errors.New("error while updating product")
	}
	return nil
}
