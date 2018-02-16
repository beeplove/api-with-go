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
        c.Header("Content-Type", "application/json")

        p := product.Product{}

        if err := c.BindJSON(&p); err != nil {
            c.JSON(http.StatusBadRequest, gin.H {
                "error":  "json decoding : " + err.Error(),
                "status": http.StatusBadRequest,
            })

            return
        }


        // TODO: what if product is not created for whatever reason
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

    r.GET("/products/query", func(c *gin.Context) {
        c.JSON(200, gin.H {
            "method": "GET",
            "path": "/products/query",
        })
    })

    r.Run()
}
