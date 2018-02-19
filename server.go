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

    /**
     *  POST /products endpoint take the following as an example in the request body
     *      {
     *          "title": "Coffee",
     *          "price": 1275       // price in cents
     *      }
    **/
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

    /**
     *  query endpoint takes three params,
     *      - title : product title
     *      - price : price of the product in cents, example: 575
     *      - comp  : comparable operator, acceptable values are: EQ | LE | LT | GE | GT | BETWEEN
    **/
    r.GET("/products/query", func(c *gin.Context) {
        c.Header("Content-Type", "application/json")

        title       := c.DefaultQuery("title", "Coffee")
        price       := c.Query("price")
        comp        := c.DefaultQuery("comp", "EQ")

        products := product.Query(title, price, strings.ToUpper(comp))

        c.JSON(http.StatusOK, gin.H {
            "status": "success",
            "data": products,
        })
    })

    r.Run()
}
