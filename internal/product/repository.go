package product

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite" // SQLite driver
)

var DB *gorm.DB

func init() {
    var err error
    DB, err = gorm.Open("sqlite3", "./ecommerce.db")
    if err != nil {
        panic("failed to connect to database")
    }

    DB.AutoMigrate(&Product{})
}

func GetProducts() ([]Product, error) {
    var products []Product
    if err := DB.Find(&products).Error; err != nil {
        return nil, err
    }
    return products, nil
}

func AddProduct(product Product) error {
    return DB.Create(&product).Error
}
