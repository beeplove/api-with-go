package dynamodbService

import (
    "fmt"
    "os"
    "strings"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/aws/awserr"
)

var svc = new()

func new() *dynamodb.DynamoDB {
    sess, err := session.NewSession(&aws.Config {
        Region: aws.String("us-west-1")},
    )

    if err != nil {
        fmt.Println("Got error when creating new AWS session")
        fmt.Println(err.Error())
        os.Exit(1)
    }

    return dynamodb.New(sess)
}

func AddRecord(item map[string]*dynamodb.AttributeValue, tableName string) (*dynamodb.PutItemOutput, error){
    input := &dynamodb.PutItemInput {
        Item: item,
        TableName: aws.String(tableName),
    }

    output, err := svc.PutItem(input)

    return output, err
}

// TODO:
//  - Need to have data type of partitionKey and sortKey
//  - consider creating a struct or look into dynamodb referece for a struct that can be used
//      to describe an attribute and it's value, this may allow to deal with too many params
func generateQueryInput(
        tableName string,
        partitionKeyName string,
        partitionKeyValue string,
        sortKeyName string,
        sortKeyValue string,
        comparisonOperator string) *dynamodb.QueryInput {

    condition := fmt.Sprintf("%s = :%s", partitionKeyName, partitionKeyName)

    expressionAttributes := map[string]*dynamodb.AttributeValue {
        fmt.Sprintf(":%s", partitionKeyName): {
            S: aws.String(partitionKeyValue),
        },
    }

    if sortKeyValue != "" {
        condition += " AND "

        expressionAttributes[fmt.Sprintf(":%s", sortKeyName)] = &dynamodb.AttributeValue {
            N: aws.String(sortKeyValue),
        }

        // Query Condition: EQ | LE | LT | GE | GT | BEGINS_WITH | BETWEEN
        switch comparisonOperator {
        case "EQ", "LE", "LT", "GE", "GT":
            condition += conditionForSortKey(sortKeyName, comparisonOperator)
        case "BETWEEN":
            condition += conditionForSortKeyRange(sortKeyName)
            expressionAttributes = expressionAttributesForSortKeyRange(partitionKeyName, partitionKeyValue, sortKeyName, sortKeyValue)
        }
    }

    input := &dynamodb.QueryInput {
        ExpressionAttributeValues:  expressionAttributes,
        KeyConditionExpression:     aws.String(condition),
        TableName:                  aws.String(tableName),
    }

    return input
}


func conditionForSortKey(sortKeyName string, comp string) string {
    switch comp {
    case "EQ":
        return fmt.Sprintf("%s  = :%s", sortKeyName, sortKeyName)
    case "LE":
        return fmt.Sprintf("%s <= :%s", sortKeyName, sortKeyName)
    case "LT":
        return fmt.Sprintf("%s <  :%s", sortKeyName, sortKeyName)
    case "GE":
        return fmt.Sprintf("%s >= :%s", sortKeyName, sortKeyName)
    case "GT":
        return fmt.Sprintf("%s >  :%s", sortKeyName, sortKeyName)
    }

    return fmt.Sprintf("%s = :%s", sortKeyName, sortKeyName)
}

func conditionForSortKeyRange(sortKeyName string) string {
    return fmt.Sprintf("%s BETWEEN :%s1 AND :%s2", sortKeyName, sortKeyName, sortKeyName)
}

func expressionAttributesForSortKeyRange(
        partitionKeyName string,
        partitionKeyValue string,
        sortKeyName string,
        sortKeyValue string) (map[string]*dynamodb.AttributeValue) {

    ranges := strings.Split(sortKeyValue, "-")

    return map[string]*dynamodb.AttributeValue {
        fmt.Sprintf(":%s", partitionKeyName): {
            S: aws.String(partitionKeyValue),
        },
        fmt.Sprintf(":%s1", sortKeyName): {
            N: aws.String(ranges[0]),
        },
        fmt.Sprintf(":%s2", sortKeyName): {
            N: aws.String(ranges[1]),
        },
    }
}


func Query(
        tableName string,
        partitionKeyName string,
        partitionKeyValue string,
        sortKeyName string,
        sortKeyValue string,
        comparisonOperator string) *dynamodb.QueryOutput {

    input := generateQueryInput(tableName, partitionKeyName, partitionKeyValue, sortKeyName, sortKeyValue, comparisonOperator)

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
        return &dynamodb.QueryOutput{}
    }

    return result
}
