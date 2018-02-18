package product

//
// TODO:
//  - Add more fields such as, categories
//  - create shipt.test from app
//

import (
    "fmt"
    "os"
    "time"

    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

    "../services"
)

/**
 * Partition Key: Title
 *  Title doesn't seems to be agod Partition Key, perhaps primaryCategory,
 *  For now keeping Title as Partition key just for demonstration purpose
 * Sort Key: Price
**/
type Product struct {
    Title string`json:"title" binding:"required"`
    Price int64`json:"price" binding:"required"`
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
func Query(title string, price string, comp string) []Product {
    resp := dynamodbService.Query(tableName, title, price, comp)

    products := []Product{}
    err := dynamodbattribute.UnmarshalListOfMaps(resp.Items,  &products)
    if err != nil {
        fmt.Errorf("failed to unmarshal Query result items, %v", err)
    }

    return products
}
