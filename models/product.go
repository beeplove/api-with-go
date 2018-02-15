package product


import (
    "fmt"
    "os"
    "time"
    // "reflect"

    "github.com/google/uuid"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

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
    tableName := "shipt.test"
    u := uuid.New()

    product := Product {
        Id: u.String(),
        Title: "Hazelnut coffee",
        Price: 8.19,
        CreatedAt: time.Now().Format(time.RFC3339),
    }

    item, err := dynamodbattribute.MarshalMap(product)
    if err != nil {
        fmt.Println("Got error calling MarshalMap:")
        fmt.Println(err.Error())
        os.Exit(1)
    }

    dynamodbService.AddRecord(item, tableName, service)
}
