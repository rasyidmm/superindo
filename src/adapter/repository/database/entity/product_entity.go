package entity

import "time"

type Product struct {
	ID        int64     `json:"id" gorm:"primaryKey"`
	SKU       string    `json:"sku"`
	Name      string    `json:"name"`
	Types     string    `json:"types"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}
