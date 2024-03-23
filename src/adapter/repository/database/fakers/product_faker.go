package fakers

import (
	"github.com/go-faker/faker/v4"
	"github.com/gosimple/slug"
	"github.com/rasyidmm/superindo-test/src/adapter/repository/database/entity"
	"gorm.io/gorm"
	"math/rand/v2"
	"time"
)

func ProductFaker() *entity.Product {
	numType := []string{"SAYUR", "PROTEIN", "BUAH", "SNAK"}
	name := faker.Name()
	rund := time.Duration(rand.IntN(100))
	return &entity.Product{
		SKU:       slug.Make(name),
		Name:      name,
		Types:     numType[rand.IntN(len(numType))],
		Price:     rand.ExpFloat64()*1000 + 1000,
		CreatedAt: time.Now().Add(rund * time.Minute),
	}

}

func ProductSeeds(db *gorm.DB) {
	if db.Find(&entity.Product{}).RowsAffected == 0 {
		tx := db.Begin()
		for i := 0; i < 10; i++ {
			if err := tx.Create(ProductFaker()).Error; err != nil {
				tx.Rollback()
			}
		}
		tx.Commit()
	}
}
