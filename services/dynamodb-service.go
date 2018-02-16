package dynamodbService

// TODO:
//  - move static config to somewhere else

import (
    "fmt"
    "os"
    
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/aws/awserr"
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

func Query(tableName string, title string) {
    fmt.Println("Service -> Query")


    // TODO: Get table metadata from model
    input := &dynamodb.QueryInput {
        ExpressionAttributeValues: map[string]*dynamodb.AttributeValue {
            ":v1": {
                S: aws.String(title),
            },
        },
        KeyConditionExpression: aws.String("title = :v1"),
        TableName:              aws.String(tableName),
    }

    result, err := svc.Query(input)
    if err != nil {
        if aerr, ok := err.(awserr.Error); ok {
            switch aerr.Code() {
            case dynamodb.ErrCodeProvisionedThroughputExceededException:
                fmt.Println(dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
            case dynamodb.ErrCodeResourceNotFoundException:
                fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
            case dynamodb.ErrCodeInternalServerError:
                fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
            default:
                fmt.Println(aerr.Error())
            }
        } else {
            // Print the error, cast err to awserr.Error to get the Code and
            // Message from an error.
            fmt.Println(err.Error())
        }
        return
    }

    fmt.Println(result)
}
