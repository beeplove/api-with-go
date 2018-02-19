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
    "strings"

    "github.com/aws/aws-sdk-go/aws"
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

var tableName = "shipt.test"

func Create(product Product) {
    product.CreatedAt = time.Now().Format(time.RFC3339)

    item, err := dynamodbattribute.MarshalMap(product)
    if err != nil {
        fmt.Println("Got error calling MarshalMap:")
        fmt.Println(err.Error())
        os.Exit(1)                                  // TODO: Can we raise exception instead of exit
    }
    fmt.Println(item)
    dynamodbService.AddRecord(item, tableName)
}

func Query(title string, price string, comp string) []Product {
    input       := queryInput(title, price, comp)
    resp        := dynamodbService.Query(input)
    products    := []Product{}

    err := dynamodbattribute.UnmarshalListOfMaps(resp.Items,  &products)
    if err != nil {
        fmt.Errorf("failed to unmarshal Query result items, %v", err)
    }

    return products
}

func queryInput(title string, price string, comp string) *dynamodb.QueryInput {
    condition := "title = :title"

    expressionAttributes := map[string]*dynamodb.AttributeValue {
        ":title": {
            S: aws.String(title),
        },
    }

    if price != "" {
        condition += " AND "

        expressionAttributes = map[string]*dynamodb.AttributeValue {
            ":title": {
                S: aws.String(title),
            },
            ":price": {
                N: aws.String(price),
            },
        }

        // Query Condition: EQ | LE | LT | GE | GT | BEGINS_WITH | BETWEEN
        switch comp {
        case "EQ", "LE", "LT", "GE", "GT":
            condition += conditionFromPrice(price, comp)
        case "BETWEEN":
            condition += conditionFromPriceRange(price)
            expressionAttributes = expressionAttributesForPriceRance(title, price)
        }
    }

    input := &dynamodb.QueryInput {
        ExpressionAttributeValues:  expressionAttributes,
        KeyConditionExpression:     aws.String(condition),
        TableName:                  aws.String(tableName),
    }

    return input
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

    return "price = :price"
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
