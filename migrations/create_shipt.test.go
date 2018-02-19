package main

import (
    "fmt"
    "os"
    "reflect"

    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/aws/awserr"
)

func tableNames(sess *session.Session) *dynamodb.ListTablesOutput {
    svc := dynamodb.New(sess)
    input := &dynamodb.ListTablesInput{}

    result, err := svc.ListTables(input)
    if err != nil {
        if aerr, ok := err.(awserr.Error); ok {
            switch aerr.Code() {
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
        return &dynamodb.ListTablesOutput{}
    }

    return result
}

func create(sess *session.Session) {
    svc := dynamodb.New(sess)
    input := &dynamodb.CreateTableInput{
        AttributeDefinitions: []*dynamodb.AttributeDefinition{
            {
                AttributeName: aws.String("title"),
                AttributeType: aws.String("S"),
            },
            {
                AttributeName: aws.String("price"),
                AttributeType: aws.String("N"),
            },
        },
        KeySchema: []*dynamodb.KeySchemaElement{
            {
                AttributeName: aws.String("title"),
                KeyType:       aws.String("HASH"),
            },
            {
                AttributeName: aws.String("price"),
                KeyType:       aws.String("RANGE"),
            },
        },
        ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
            ReadCapacityUnits:  aws.Int64(5),
            WriteCapacityUnits: aws.Int64(5),
        },
        TableName: aws.String("shipt.test"),
    }

    result, err := svc.CreateTable(input)
    if err != nil {
        if aerr, ok := err.(awserr.Error); ok {
            switch aerr.Code() {
            case dynamodb.ErrCodeResourceInUseException:
                fmt.Println(dynamodb.ErrCodeResourceInUseException, aerr.Error())
            case dynamodb.ErrCodeLimitExceededException:
                fmt.Println(dynamodb.ErrCodeLimitExceededException, aerr.Error())
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

func main() {
    // Unused variables
    _ = sess
    _ = reflect.TypeOf

    sess, err := session.NewSession(&aws.Config {
        Region: aws.String("us-west-1")},
    )

    if err != nil {
        fmt.Println("Got error when creating new AWS session")
        fmt.Println(err.Error())
        os.Exit(1)
    }

    create(sess)
}
