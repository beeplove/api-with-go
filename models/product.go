package product

//
// TODO:
//  - Add more fields such as, categories
//  - create shipt.test from app
//  - need to check to conver price from float to decimal
//

import (
    "fmt"
    "os"
    "time"

    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

    "../services"
)

// Partition Key: Title
// Sort Key: Price
type Product struct {
    Title string`json:"title" binding:"required"`
    Price float64`json:"price" binding:"required"`
    CreatedAt string`json:"createdAt"`
}

var tableName = "shipt.test"

func Create(product Product) {
    product.CreatedAt = time.Now().Format(time.RFC3339)

    item, err := dynamodbattribute.MarshalMap(product)
    if err != nil {
        fmt.Println("Got error calling MarshalMap:")
        fmt.Println(err.Error())
        os.Exit(1)
    }

    dynamodbService.AddRecord(item, tableName)
}

// TODO: Make changes to accomodate price include range and other comparable operators
func Query(title string, price string, comp string) {
    dynamodbService.Query(tableName, title, price, comp)
}
