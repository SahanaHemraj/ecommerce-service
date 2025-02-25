package main

import (
    "github.com/gin-gonic/gin"
    "ecommerce-website/internal/product"
    "ecommerce-website/internal/order"
    "ecommerce-website/internal/user"
	"log"
	"ecommerce-website/internal/database"

)

func main() {

	database.InitDB()
	database.Migrate()

    r := gin.Default()

    // Routes for products, orders, and users
    r.GET("/products", product.GetAllProducts)
    r.POST("/products", product.AddProduct)

    r.GET("/orders", order.GetAllOrders)
    r.POST("/orders", order.PlaceOrder)

    r.POST("/signup", user.SignUp)
    r.POST("/login", user.Login)

    // Static files for frontend (e.g., CSS, JS)
    r.Static("/static", "./static")

    r.Run(":8080") // Listen on port 8080
}
