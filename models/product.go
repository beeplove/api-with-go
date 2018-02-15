package product


import (
    "fmt"

    "../services"
)

// 
// TODO:
//  - Add more fields such as, categories
// 
type Product struct {
    Id string`json:"id"`
    Title string`json:"title"`
    Price float64`json:"price"`
    CreatedAt string`json:"createdAt"`
}

var service = dynamodbService.New()

func Create() {
    fmt.Println("Product Create")
    fmt.Println(service)
}
