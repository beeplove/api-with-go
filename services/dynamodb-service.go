package dynamodbService

// TODO:
//  - move static config to somewhere else

import (
    "fmt"
    "os"
    "strings"

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

func conditionFromPrice(price string, comp string) string {
    switch comp {
    case "EQ":
        return "price  = :price"
    case "LE":
        return "price <= :price"
    case "LT":
        return "price <  :price"
    case "GE":
        return "price >= :price"
    case "GT":
        return "price >  :price"
    }

    return "price  = :price"
}

func conditionFromPriceRange(price string) string {
    return "price BETWEEN :price1 AND :price2"
}

func expressionAttributesForPriceRance(title string, price string) map[string]*dynamodb.AttributeValue {
    prices := strings.Split(price, "-")

    return map[string]*dynamodb.AttributeValue {
        ":title": {
            S: aws.String(title),
        },
        ":price1": {
            N: aws.String(prices[0]),
        },
        ":price2": {
            N: aws.String(prices[1]),
        },
    }
}

// Query Condition: EQ | LE | LT | GE | GT | BEGINS_WITH | BETWEEN
// TODO: Cover scenario when price is absent
func Query(tableName string, title string, price string, comp string) {
    condition := "title = :title AND "
    expressionAttributes := map[string]*dynamodb.AttributeValue {
        ":title": {
            S: aws.String(title),
        },
        ":price": {
            N: aws.String(price),
        },
    }

    switch comp {
    case "EQ", "LE", "LT", "GE", "GT":
        condition += conditionFromPrice(price, comp)
    case "BETWEEN":
        condition += conditionFromPriceRange(price)
        expressionAttributes = expressionAttributesForPriceRance(title, price)
    }

    // TODO: Get table metadata from model
    input := &dynamodb.QueryInput {
        ExpressionAttributeValues:  expressionAttributes,
        KeyConditionExpression:     aws.String(condition),
        TableName:                  aws.String(tableName),
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
