package product

import (
    "github.com/jinzhu/gorm"
)

type Product struct {
    gorm.Model
    Name        string  `json:"name"`
    Description string  `json:"description"`
    Price       float64 `json:"price"`
}
