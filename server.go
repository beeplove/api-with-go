package main

import "github.com/gin-gonic/gin"

func main() {
    r:= gin.Default()

    r.POST("/products", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "method": "POST",
            "path": "/products",
        })
    })

    r.GET("/products", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "method": "GET",
            "path": "/products",
        })
    })

    r.GET("/products/search", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "method": "GET",
            "path": "/products/search",
        })
    })

    r.Run()
}
