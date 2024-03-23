package products

import (
	"context"
	"encoding/json"
	"github.com/rasyidmm/superindo-test/src/adapter/repository/database/entity"
	"github.com/rasyidmm/superindo-test/src/shared/constanta"
	"github.com/rasyidmm/superindo-test/src/usecase/payload"
	"sort"
	"strings"
	"time"
)

func (uc *ProductsUsecase) GetProducts(ctx context.Context, input payload.GetProductsRequest) (*[]entity.Product, error) {
	var products *[]entity.Product

	useCache := true
	c, err := uc.rpr.GetCache(ctx, constanta.ProductList)
	if err != nil {
		useCache = false
	}
	if err := json.Unmarshal([]byte(c), &products); err != nil {
		useCache = false
	}
	if !useCache {
		products, err = uc.rpd.GetProducts(ctx)
		if err != nil {
			return nil, err
		}
		j, _ := json.Marshal(products)
		uc.rpr.SetCache(ctx, constanta.ProductList, string(j), 12*time.Hour)
	}
	if input.Find != "" {
		productsFind := &[]entity.Product{}
		for _, v := range *products {
			if strings.Contains(v.Name, input.Find) {
				*productsFind = append(*productsFind, v)
			}
		}
		products = productsFind
	}

	if input.Filter != "" {
		productsFilter := &[]entity.Product{}
		for _, v := range *products {
			if v.Types == input.Filter {
				*productsFilter = append(*productsFilter, v)
			}
		}
		products = productsFilter
	}
	if input.Sort != "" {
		products = SortProduct(input.Sort, products)
	}
	return products, nil

}

func SortProduct(sortType string, products *[]entity.Product) *[]entity.Product {

	sortTypeItem := strings.Split(sortType, ".")
	if len(sortTypeItem) < 2 {
		return products
	}

	if sortTypeItem[0] == "name" {
		if sortTypeItem[1] == "asc" {
			sort.Slice(*products, func(i, j int) bool {
				return (*products)[i].Name < (*products)[j].Name
			})
		} else if sortTypeItem[1] == "desc" {
			sort.Slice(*products, func(i, j int) bool {
				return (*products)[i].Name > (*products)[j].Name
			})
		}

	} else if sortTypeItem[0] == "price" {
		if sortTypeItem[1] == "asc" {
			sort.Slice(*products, func(i, j int) bool {
				return (*products)[i].Price < (*products)[j].Price
			})
		} else if sortTypeItem[1] == "desc" {
			sort.Slice(*products, func(i, j int) bool {
				return (*products)[i].Price > (*products)[j].Price
			})
		}

	} else if sortTypeItem[0] == "created_at" {
		if sortTypeItem[1] == "asc" {
			sort.Slice(*products, func(i, j int) bool {
				return (*products)[i].CreatedAt.Before((*products)[j].CreatedAt)
			})
		} else if sortTypeItem[1] == "desc" {
			sort.Slice(*products, func(i, j int) bool {
				return (*products)[i].CreatedAt.After((*products)[j].CreatedAt)
			})
		}

	}

	return products
}
