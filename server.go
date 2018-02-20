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
            c.JSON(http.StatusUnprocessableEntity, gin.H {
                "status": "error",
                "code": 2001,       // application error code
                "message": err.Error(),
                "description": "for more information, visit http://localhost:8080/errors/2001",
            })

            return
        }

        _, err := product.Create(p)

        if err != nil {
            c.JSON(http.StatusUnprocessableEntity, gin.H {
                "status": "error",
                "code": 2101,       // application error code
                "message": err.Error(),
                "description": "for more information, visit http://localhost:8080/errors/2101",
            })

            return
        }

        c.JSON(http.StatusOK, gin.H {
            "status": "success",
        })
    })

    /**
     *  query endpoint takes three params,
     *      - title : product title
     *      - price : price of the product in cents, example: 575
     *      - comp  : comparable operator, acceptable values are: EQ | LE | LT | GE | GT | BETWEEN
     *  Example: ?title=Coffee&price=575&comp=GE
     *      shoule return all items which have title Coffee and price greater than or equal to 575
    **/
    r.GET("/products/query", func(c *gin.Context) {
        c.Header("Content-Type", "application/json")

        title       := c.DefaultQuery("title", "Coffee")
        price       := c.Query("price")
        comp        := c.DefaultQuery("comp", "EQ")

        products, err := product.Query(title, price, strings.ToUpper(comp))

        if err != nil {
            c.JSON(http.StatusUnprocessableEntity, gin.H {
                "status": "error",
                "code": 1001,       // application error code
                "message": "Something went wrong",
                "description": "for more information, visit http://localhost:8080/errors/1001",
            })

            return
        }

        c.JSON(http.StatusOK, gin.H {
            "status": "success",
            "data": products,
        })
    })

    r.Run()
}
