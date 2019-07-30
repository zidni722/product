package response

import (
	"github.com/jinzhu/gorm"
	"github.com/zidni722/pawoon-product/app/dto/response"
	"github.com/zidni722/pawoon-product/app/models"
)

type ProductResponse struct {
	Db *gorm.DB
}

func NewProductResponse(db *gorm.DB) ProductResponse {
	return ProductResponse{Db: db}
}

func (r *ProductResponse) New(product models.Product) response.Product {
	response := response.Product{
		ID:   product.ID,
		Name: product.Name,
	}

	return response
}

func (r *ProductResponse) Collection(products []models.Product) []response.Product {
	var responses []response.Product

	for _, product := range products {
		responses = append(responses, r.New(product))
	}

	return responses
}
