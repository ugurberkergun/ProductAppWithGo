package service

import (
	"ProductAppWithGo/domain"
	"ProductAppWithGo/persistence"
)

type FakeProductRepository struct {
	products []domain.Product
}

func NewFakeProductRepository(initialProducts []domain.Product) persistence.IProductRepository {
	return &FakeProductRepository{
		products: initialProducts,
	}
}

func (fakeRepository *FakeProductRepository) GetAllProducts() []domain.Product {
	return fakeRepository.products
}

func (fakeRepository *FakeProductRepository) GetAllProductsByStore(storeName string) []domain.Product {
	// todo:ü
	return []domain.Product{}
}

func (fakeRepository *FakeProductRepository) AddProduct(product domain.Product) error {
	fakeRepository.products = append(fakeRepository.products, domain.Product{
		Id:       int64(len(fakeRepository.products)) + 1,
		Name:     product.Name,
		Price:    product.Price,
		Store:    product.Store,
		Discount: product.Discount,
	})
	return nil
}

func (fakeRepository *FakeProductRepository) GetById(productId int64) (domain.Product, error) {
	// todo
	return domain.Product{}, nil
}

func (fakeRepository *FakeProductRepository) DeleteById(productId int64) error {
	// todo
	return nil
}

func (fakeRepository *FakeProductRepository) UpdatePrice(newPrice float32, productId int64) error {
	// todo
	return nil
}
