package main

import (
    "strings"
    "net/http"

    "github.com/gin-gonic/gin"

    "./models"
)

func main() {
    r := gin.Default()

    r.GET("/health", func(c *gin.Context) {
        c.Header("Content-Type", "application/json")
        c.JSON(http.StatusOK, gin.H {
            "status": "success",
        })
    })

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
            "status": "success",
        })
    })

    r.GET("/products/query", func(c *gin.Context) {
        c.Header("Content-Type", "application/json")

        title :=    c.DefaultQuery("title", "Coffee")
        comp :=     c.DefaultQuery("comp", "EQ")
        price :=    c.Query("price")

        products := product.Query(title, price, strings.ToUpper(comp))

        c.JSON(http.StatusOK, gin.H {
            "status": "success",
            "data": products,
        })
    })

    r.Run()
}
