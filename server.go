package main

import (
    // "fmt"

    "github.com/gin-gonic/gin"

    "./models"
)

func main() {
    r := gin.Default()

    r.POST("/products", func(c *gin.Context) {

        product.Create()

        c.JSON(200, gin.H {
            "method": "POST",
            "path": "/products",
        })
    })

    r.GET("/products", func(c *gin.Context) {
        c.JSON(200, gin.H {
            "method": "GET",
            "path": "/products",
        })
    })

    r.GET("/products/search", func(c *gin.Context) {
        c.JSON(200, gin.H {
            "method": "GET",
            "path": "/products/search",
        })
    })

    r.Run()
}
