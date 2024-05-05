package response

import "ProductAppWithGo/domain"

type ErrorResponse struct {
	ErrorDescription string `json:"error_description"`
}

type ProductResponse struct {
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Discount float32 `json:"discount"`
	Store    string  `json:"store"`
}

func ToResponse(product domain.Product) ProductResponse {
	return ProductResponse{
		Name:     product.Name,
		Price:    product.Price,
		Discount: product.Discount,
		Store:    product.Store,
	}
}

func ToResponseList(product []domain.Product) []ProductResponse {
	var productResponseList = []ProductResponse{}
	for _, product := range product {
		productResponseList = append(productResponseList, ToResponse(product))
	}
	return productResponseList
}
