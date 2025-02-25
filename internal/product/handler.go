package product

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "ecommerce-website/internal/product"
)

func GetAllProducts(c *gin.Context) {
    products, err := product.GetProducts()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get products"})
        return
    }
    c.JSON(http.StatusOK, products)
}

func AddProduct(c *gin.Context) {
    var newProduct product.Product
    if err := c.ShouldBindJSON(&newProduct); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    product.AddProduct(newProduct)
    c.JSON(http.StatusOK, gin.H{"message": "Product added successfully"})
}
