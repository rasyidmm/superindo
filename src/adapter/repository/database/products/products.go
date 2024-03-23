package products

import (
	"context"
	"github.com/rasyidmm/superindo-test/src/adapter/repository/database/entity"
	"gorm.io/gorm"
)

type ProductsDataHandler struct {
	db *gorm.DB
}

func NewProductsDataHandler(db *gorm.DB) *ProductsDataHandler {
	return &ProductsDataHandler{
		db: db,
	}
}

func (uc *ProductsDataHandler) CreateProduct(ctx context.Context, input *entity.Product) error {

	return uc.db.Create(input).Error
}
func (uc *ProductsDataHandler) GetProducts(ctx context.Context) (*[]entity.Product, error) {
	var products *[]entity.Product
	err := uc.db.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, err
}
