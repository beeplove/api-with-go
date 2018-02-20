## To Start Server

```
% export AWS_ACCESS_KEY_ID=<AWS_ACCESS_KEY_ID>
$ export AWS_SECRET_ACCESS_KEY=<AWS_SECRET_ACCESS_KEY>
% go run server.go
```
server should be running on port 8080 and can be checked by visiting http://localhost:8080/health, which should return the following:

```
{
    "status": "success"
}
```

## Available Endpoints

#### POST /products

To create new product, which take json in the request body, example:

```
{
    "title": "Product Title",
    "price": 1275
}
```

price take integer value which represents price of the product in cents

#### GET /products/query

query endpoint take folling three params
- title - title of the product
- price - price of the product in cents
- comp - EQ | LE | LT | GE | GT | BETWEEN

Examples:

- `/products/query?title=Coffee`

to return all item which has title Coffee

- `/products/query?title=Cofdee&price=1275`

to return item with title Coffee and price 1275

- `/products/query?title=Coffee&price=1275&comp=GE`

to return items with title Coffee and price greater than or equal to 1275

- `/proudcts/query?title=Coffee&price=500-2200&comp=BETWEEN`

to return items with title Coffee and price is between 500 and 2200
