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

1. POST /products

which take json in the request body, example:

```
{
    "title": "Product Title",
    "price": 1275
}
```

price take integer value which represents price of the product in cents

2. GET /products/query

query endpoint take folling three params
- title
- price
- comp
