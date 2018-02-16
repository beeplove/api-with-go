package main

import (
    // "fmt"
    "net/http"

    "github.com/gin-gonic/gin"

    "./models"
)

func main() {
    r := gin.Default()

    r.POST("/products", func(c *gin.Context) {
        c.Header("Content-Type", "application/json; charset=utf-8")

        p := product.Product{}

        if err := c.BindJSON(&p); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error":  "json decoding : " + err.Error(),
                "status": http.StatusBadRequest,
            })

            return
        }

        product.Create(p)

        c.JSON(http.StatusOK, gin.H {
            "Title": p.Title,
            "Price": p.Price,
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
