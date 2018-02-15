package dynamodbService

// TODO:
//  - make a simple command line client to add item in table
//  - convert to a service
//      - keep everything in this layer to talk to DynamoDB
//      - move applicatioin model out of this layer
//      - move static config to somewhere else

import (
    "fmt"
    "os"

    // "time"

    // "github.com/google/uuid"
    
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    // "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)



// func Add(product Product, tableName string, svc *dynamodb.DynamoDB) {
//     av, err := dynamodbattribute.MarshalMap(product)

//     input := &dynamodb.PutItemInput {
//         Item: av,
//         TableName: aws.String(tableName),
//     }

//     _, err = svc.PutItem(input)

//     if err != nil {
//         fmt.Println("Got error calling PutItem:")
//         fmt.Println(err.Error())
//         os.Exit(1)
//     }

// }

func New() *dynamodb.DynamoDB {
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String("us-west-1")},
    )

    if err != nil {
        fmt.Println("Got error when creating new AWS session")
        fmt.Println(err.Error())
        os.Exit(1)
    }

    return dynamodb.New(sess)
}

func Create() {
    fmt.Println("Service Create")
}

// func main() {
//     sess, err := session.NewSession(&aws.Config{
//         Region: aws.String("us-west-1")},
//     )

//     if err != nil {
//         fmt.Println("Got error when creating new AWS session")
//         fmt.Println(err.Error())
//         os.Exit(1)
//     }

//     // Create DynamoDB client
//     svc := dynamodb.New(sess)

//     u := uuid.New()
//     table := "shipt.test"

//     product := Product {
//         Id: u.String(),
//         Title: "Hazelnut coffee",
//         Price: 8.19,
//         CreatedAt: time.Now().Format(time.RFC3339),
//     }

//     Add(product, table, svc)

//     fmt.Println("Item added, successfully")
// }
