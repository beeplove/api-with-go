package main

// TODO:
//  - make a simple command line client to add item in table
//  - convert to a service
//      - keep everything in this layer to talk to DynamoDB
//      - move applicatioin model out of this layer
//      - move static config to somewhere else

import (
    "fmt"
    "os"

    // "reflect"
    "time"

    "github.com/google/uuid"
    
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
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

func add(product Product, tableName string, svc *dynamodb.DynamoDB) {
    av, err := dynamodbattribute.MarshalMap(product)

    input := &dynamodb.PutItemInput {
        Item: av,
        TableName: aws.String(tableName),
    }

    _, err = svc.PutItem(input)

    if err != nil {
        fmt.Println("Got error calling PutItem:")
        fmt.Println(err.Error())
        os.Exit(1)
    }

}

func main() {
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String("us-west-1")},
    )

    if err != nil {
        fmt.Println("Got error when creating new AWS session")
        fmt.Println(err.Error())
        os.Exit(1)
    }

    // Create DynamoDB client
    svc := dynamodb.New(sess)

    // t := time.Unix()
    // fmt.Println(time.Now().Format(time.RFC3339))


    u := uuid.New()
    table := "shipt.test"

    product := Product {
        Id: u.String(),
        Title: "Hazelnut coffee",
        Price: 8.19,
        CreatedAt: time.Now().Format(time.RFC3339),
    }


    add(product, table, svc)

    fmt.Println("Item added, successfully")
}
