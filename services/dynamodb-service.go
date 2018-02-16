package dynamodbService

// TODO:
//  - move static config to somewhere else

import (
    "fmt"
    "os"
    
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
)

var svc = New()


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

func AddRecord(item map[string]*dynamodb.AttributeValue, tableName string) {
    input := &dynamodb.PutItemInput {
        Item: item,
        TableName: aws.String(tableName),
    }

    _, err := svc.PutItem(input)

    if err != nil {
        fmt.Println("Got error calling PutItem:")
        fmt.Println(err.Error())
        os.Exit(1)
    }
}
