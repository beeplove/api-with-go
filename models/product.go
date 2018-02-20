package product

/**
 * product package is to represent application model Product
 *
**/

import (
    "fmt"
    "time"

    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
    "github.com/aws/aws-sdk-go/service/dynamodb"

    "../services"
)

/**
 * Partition Key: Title
 *  Title doesn't seems to be agod Partition Key, perhaps primaryCategory,
 *  For now keeping Title as Partition key just for demonstration purpose
 * Sort Key: Price
**/
type Product struct {
    Title string        `json:"title" binding:"required"`   // Title of the product
    Price int64         `json:"price" binding:"required"`   // Price of product in Cents
    CreatedAt string    `json:"createdAt"`                  // to be generated automatically
}

type ProductError struct {
    Err error
}

var tableName = "shipt.test"


// TODO:
//  Not sure how to extract item created from PutItemOutput,
//  otherewise return Product instead of PutItemOutput,
//  which would allow model to not to be tightly coupled with service
func Create(product Product) (*dynamodb.PutItemOutput, error) {
    product.CreatedAt = time.Now().Format(time.RFC3339)

    item, err := dynamodbattribute.MarshalMap(product)
    if err != nil {
        return &dynamodb.PutItemOutput{}, err
    }

    output, err := dynamodbService.AddRecord(item, tableName)

    return output, err
}

func Query(title string, price string, comp string) (products []Product, err error) {
    resp        := dynamodbService.Query(tableName, "title", title, "price", price, comp)
    products    = []Product{}

    err = dynamodbattribute.UnmarshalListOfMaps(resp.Items, &products)
    if err != nil {
        fmt.Errorf("failed to unmarshal Query result items, %v", err)
    }

    return products, err
}
